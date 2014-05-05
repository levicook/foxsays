foxsays
=======

Fork and clone git@github.com:levicook/foxsays.git


## First pass: make sure everything builds:

```bash
cd foxsays
source dev.env
build-client
build-server
```


## Ongoing frontend development: strart the client daemons in a fresh terminal.

```bash
cd foxsays
source dev.env fast
./client/daemons/start
```

Then browse http://localhost:3081

When your client changes are ready for integration, run `lint-js && test-js`.


## Ongoing full stack development: strart the client daemons (above) and the server daemons in a fresh terminal.
```bash
cd foxsays
source dev.env fast
./server/daemons/start
```

Then browse http://localhost:3080

When your client changes are ready for integration with the server, run `build-client && restart-server`

When your server changes are ready for integration, run `build-server && restart-server`
