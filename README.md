# dsqlstate

dsqlstate is a discord state tracker that uses postgres. For those that do not want to keep the entire state in memory.

**In early development**, everything is subject to change, the api, db schema, EVERTHING. Expect stuff to be wrong if you dare to use this at this stage, and probably panic as well.

Currently only supports postgres, this may or may not change in the future, probably not unless someone other than me makes the appropriate changes.

You might thing that means you can have multiple bots use the same database but that is not yet supported, what happens when bot 1 leaves a guild but bot 2 is still in it? without the bots communicating to eachother or another join table or an array in the guilds table it would be marked as `left` from both bots, may be added in the future but now the focus is to get it working.

There are a lot of benefits to keeping the state in a proper database, being able to inspect the state and all helps when you're hunting bugs in your code.

The biggest downside is probably the initial load, it has to invalidate all state data between `ready`'s because stuff may have changed in the meantime and we did not receive events about it.
The initial load consists of:

 1. Invalidating all state data this is done relatively quickly, by just setting `left_at` and `deleted_at` in everything to now
 2. processing the ready event, updating/creating guilds that were available and marking them for todo if not
 3. handling all the guild create

The initial load of a guild consists of:
 
 1. updating/creating all roles
 2. updating/creating all channels
 3. updating/creating all members/users/presences included in the guild create. This is the most taxing one by far.

The initial load is done on a per shard level, meaning if one of your shards disconnects and are unable to resume, it does not have to invalidate all other shards state info.

The inistial load per shard (1700 guilds) is in the range of a couple minutes
you can play a gamble, in which you trust the previous state info even though its invalid until the initial load is complete.

## Features

 - [x] Everything is stored in a postgres database, not bound by memory restrictions
 - [x] All message edits are tracked
 - [ ] Changes to state entries are tracked, meaning you can get a list of people that joined a guild while the bot and state server was offline
     - This is currently being worked on, some entries are tracked, some are not atm.



