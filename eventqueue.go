package dsqlstate

import (
	"encoding/json"
	"github.com/beeker1121/goque"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"io/ioutil"
)

type EventQueue struct {
	queue *goque.Queue
}

func NewEventQueue() (*EventQueue, error) {
	dir, err := ioutil.TempDir("", "dsqlstate")
	if err != nil {
		return nil, errors.Wrap(err, "NewEventQueue")
	}
	queue, err := goque.OpenQueue(dir)
	if err != nil {
		return nil, errors.Wrap(err, "NewEventQueue")
	}

	return &EventQueue{
		queue: queue,
	}, nil
}

func EvtToCode(evt interface{}) (code byte, ok bool) {
	ok = true
	switch evt.(type) {
	case *discordgo.GuildUpdate:
		code = 1
	// Members
	case *discordgo.GuildMemberAdd:
		code = 2
	case *discordgo.GuildMemberUpdate:
		code = 3
	case *discordgo.GuildMemberRemove:
		code = 4
	case *discordgo.GuildRoleCreate:
		code = 5
	case *discordgo.GuildRoleUpdate:
		code = 6
	case *discordgo.GuildRoleDelete:
		code = 7

	// Channels
	case *discordgo.ChannelCreate:
		code = 8
	case *discordgo.ChannelUpdate:
		code = 9
	case *discordgo.ChannelDelete:
		code = 10
	// Messages
	case *discordgo.MessageCreate:
		code = 11
	case *discordgo.MessageUpdate:
		code = 12
	case *discordgo.MessageDelete:
		code = 13

	// Other
	case *discordgo.VoiceStateUpdate:
		code = 14
	case *discordgo.UserUpdate:
		code = 15
	case *discordgo.PresenceUpdate:
		code = 16
	case *discordgo.GuildMembersChunk:
		code = 17
	default:
		ok = false
	}
	return
}

func CodeToEvt(code byte) (evt interface{}, ok bool) {
	ok = true
	switch code {
	case 1:
		evt = &discordgo.GuildUpdate{}
	// Members
	case 2:
		evt = &discordgo.GuildMemberAdd{}
	case 3:
		evt = &discordgo.GuildMemberUpdate{}
	case 4:
		evt = &discordgo.GuildMemberRemove{}
	case 5:
		evt = &discordgo.GuildRoleCreate{}
	case 6:
		evt = &discordgo.GuildRoleUpdate{}
	case 7:
		evt = &discordgo.GuildRoleDelete{}

	// Channels
	case 8:
		evt = &discordgo.ChannelCreate{}
	case 9:
		evt = &discordgo.ChannelUpdate{}
	case 10:
		evt = &discordgo.ChannelDelete{}
	// Messages
	case 11:
		evt = &discordgo.MessageCreate{}
	case 12:
		evt = &discordgo.MessageUpdate{}
	case 13:
		evt = &discordgo.MessageDelete{}

	// Other
	case 14:
		evt = &discordgo.VoiceStateUpdate{}
	case 15:
		evt = &discordgo.UserUpdate{}
	case 16:
		evt = &discordgo.PresenceUpdate{}
	case 17:
		evt = &discordgo.GuildMembersChunk{}
	default:
		ok = false
	}
	return
}

func (q *EventQueue) QueueEvent(evt interface{}) error {
	evtCode, ok := EvtToCode(evt)
	if !ok {
		return nil
	}

	serialized, err := json.Marshal(evt)
	if err != nil {
		return errors.Wrap(err, "QueueEvent")
	}
	body := append([]byte{evtCode}, serialized...)
	_, err = q.queue.Enqueue(body)
	if err != nil {
		return errors.Wrap(err, "QueueEvent")
	}

	return nil
}

var (
	ErrEmpty = errors.New("Empty queue")
)

func (q *EventQueue) GetEvent() (interface{}, error) {

	item, err := q.queue.Dequeue()
	if err != nil {
		if err == goque.ErrEmpty {
			return nil, ErrEmpty
		}
		return nil, errors.Wrap(err, "GetEvent")
	}
	v := item.Value
	if len(v) < 1 {
		return nil, errors.New("Invalid queued event")
	}

	id := v[0]
	dest, ok := CodeToEvt(id)
	if !ok {
		return nil, errors.New("Invalid event code")
	}

	err = json.Unmarshal(v[1:], dest)
	if err != nil {
		return nil, errors.Wrap(err, "GetEvent")
	}

	return dest, nil
}
