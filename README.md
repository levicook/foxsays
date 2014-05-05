foxsays
=======

Fork and clone git@github.com:levicook/foxsays.git


### First pass: make sure everything builds:

```bash
cd foxsays
source dev.env
build-client
build-server
```


### Ongoing frontend development

Start the client daemons in a fresh terminal.

```bash
cd foxsays
source dev.env fast
start-client
```

Then browse http://localhost:3081

When your client changes are ready for integration, run `lint-js && test-js`.


### Ongoing full stack development

Start the client daemons (above) and the server daemons in a fresh terminal.

```bash
cd foxsays
source dev.env fast
./server/daemons/start # note, this will emit errors until build-client has a clean run
start-server
```

Then browse http://localhost:3080

When your client changes are ready for integration with the server, run `build-client && restart-server`

When your server changes are ready for integration, run `build-server && restart-server`

### Project Layout

```
client/src/
├── admin
│   ├── components
│   │   ├── footer
│   │   └── header
│   ├── helpers
│   │   └── global_helpers
│   └── pages
│       └── dashboard
├── demo
├── main
│   ├── components
│   │   ├── footer
│   │   └── header
│   ├── helpers
│   │   └── global_helpers
│   └── pages
│       └── dashboard
└── shared
    ├── assert
    ├── data_pool
    ├── layout
    ├── models
    │   └── first_person
    └── strings
server/src/foxsays/
├── config
├── httpd
│   ├── admin
│   │   ├── api
│   │   │   └── images
│   │   └── routes
│   ├── main
│   │   ├── api
│   │   │   └── users
│   │   ├── images
│   │   ├── pages
│   │   │   └── dashboard
│   │   └── routes
│   ├── page
│   ├── route
│   ├── router
│   └── utils
│       └── gzip
├── log
├── mime
├── models
└── repos

```
