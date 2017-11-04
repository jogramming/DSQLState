package dsqlstate

import (
	"database/sql"
	"sync"
	"time"
)

type UserCheckState struct {
	done     bool
	lastTime time.Time
}

// UserCheckCache keeps track of what users are currently being checked in the db
// because all users are updated at the same time during initialisation and inside giant transactions
// we have to maintain this list to make sure we don't run into a deadlock
type UserCheckCache struct {
	sync.RWMutex

	checking map[string]*UserCheckState
	// checked  map[string]string
}

func NewUserCheckCache() *UserCheckCache {
	return &UserCheckCache{
		checking: make(map[string]*UserCheckState),
	}
}

func (uc *UserCheckCache) WaitForUser(id string) {
	for {
		uc.RLock()
		if v, ok := uc.checking[id]; ok {
			if v.done {
				uc.RUnlock()
				return
			}
			uc.RUnlock()
			time.Sleep(time.Millisecond * 10)
		} else {
			uc.RUnlock()
			return
		}
	}
}

func (uc *UserCheckCache) MarkAsChecking(id string) (didMark bool, shouldWaitFor bool) {
	var existing *UserCheckState
	if v, ok := uc.checking[id]; ok {
		if !v.done {
			return false, true
		}

		if time.Since(v.lastTime) < time.Minute*5 {
			return false, false
		}

		existing = v
	} else {
		existing = &UserCheckState{}
		uc.checking[id] = existing
	}

	existing.lastTime = time.Now()
	existing.done = false

	return true, false
}

func (uc *UserCheckCache) MarkAsDoneChecking(id string) {
	uc.checking[id].done = true
}

// NumGuildsPerShard returns the number of guilds per shard
func NumGuildsPerShard(db *sql.DB, numShard int) ([]int, error) {
	if numShard < 1 {
		numShard = 1
	}

	result := make([]int, numShard)

	transaction, err := db.Begin()
	if err != nil {
		return result, err
	}

	for i := 0; i < numShard; i++ {
		var count int
		row := transaction.QueryRow(`SELECT COUNT(*) FROM d_guilds WHERE (id >> 22) % $1 = $2`, numShard, i)

		err = row.Scan(&count)
		if err != nil {
			return result, err
		}

		result[i] = count
	}

	err = transaction.Commit()
	return result, err
}
