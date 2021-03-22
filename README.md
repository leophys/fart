# This is a work in progress, nothing of what is described here is ready yet

This is a draft of something wanted. If you are curious, take a look at
[docs/ROADMAP.md][m] and [docs/RESOURCES.md][r]

[m]: ./docs/ROADMAP.md
[r]: ./docs/RESOURCES.md

# Fart
#### A nicely named intercept proxy

Fart is the name for the project that collects a series of tools to accomplish
what the monolithic (and more importantly closed source) [Burp Suite][bs] from
Portswigger does.

This repository contains the core parts, that are written in golang. The
architecture is described in [docs/ARCHITECTURE.md][a]


[bs]: https://portswigger.net/burp
[a]: ./docs/ARCHITECTURE.md

## Features

With Fart, one should be able to spawn an intercept proxy that should be able to
intercept http/https and websocket messages, allowing the user to inspect and
modify both the outgoing request and the incoming response between a client
(i.e. a browser) and a server, in an interactive fashon.

To easen the burden of manual inspecting the request/response, one should be
able to configure Fart to selectively intercept and hold for
inspection/modification, based on some chosen rule such as the host, the
extension of the resource, it's `Content-Type`, some regular expression on the
URI and so on.

The core should also be able to persist the intercepted request in some
organized (and possibly querable) form for later inspection and to export the
requests in some format that allows reproduction (i.e. curl shell command).

The core should also expose the possibility to:

 - repeat a selected request (and easily show diffs between responses)
 - fuzz a parametrized request
 - (nice to have) being expandable by means of some form of scripting

## Why?

Burp Suite is great, but it's closed source and we love open source.

## Usage

The application in this repository is a cli, called `fart`, with the following
verbs:

  - [ ] `fart serve`: spin up a server instance
  - [ ] `fart tui`: start an interactive user interface

### `fart serve`

This subcommand starts the instance of a server and binds two sockets:

 - [ ] a proxy socket, where it behaves as an HTTP proxy
 - [ ] a websocket to expose the captured data to a client
 - [ ] a control socket, to control the internal state of the server from
   another program (likely the cli)

## Questions

You can open an issue here, or drop an email at `blallo -|AT|- autistici[.]org`
