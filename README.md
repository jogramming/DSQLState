# dsqlstate

dsqlstate is a discord state tracker that uses postgres as the storage, for those that do not want to keep the entire state in memory.

**In early development**, everythign is subject to change, the api, db schema, EVERTHING. Expect stuff to be wrong if you dare to use this at this stage, and probably panic aswell.

Currently only supports postgres, this may or may not change in the future, probably not unless someone other than me makes the appropriate changes.

There are a lot of benefits to keeping the state in a proper database, being able to inspect the state and all helps when you're hunting bugs in your code.