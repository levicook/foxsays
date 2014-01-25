# foxsays

Simple Twitter clone


## Motivation

This project exists to demonstrate and explain a boatload of effective
development practices.


## Development

Prereqs (make sure this stuff is installed)

* daemontools
* git
* go
* node

###  Working on the project

Fork and clone this repository.

Always work in a hermetic project environment:

```bash
cd foxsays
source dev.env
```


### Using glitch

```bash
cd foxsays
./script/glitch
```


### Daemons management

Starting them: `./script/daemons/start`

Restarting them: `./script/daemons/restart`

Checking their status `./script/daemons/status`
