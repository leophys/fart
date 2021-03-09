# Architecture

We like interconnected yet disjoint programs interacting via open standards. The
idea is to create a server application that runs with high privileges somewhere
in the path of the traffic-to-be-intercepted (i.e. on the local machine or on
the network gateway). This server application should be able to intercept the
traffic, both exposing itself as a proxy server and also acting as a transparent
proxy.


```
┌──────────┐
│          │
│          │
│  CLIENT  │
│          │
│          │
└───┬──────┘
    │  ▲
    │  │
    │  │
    │  │
    │  │                       ┌─────────────────┐
    │  │                       │                 │
    │  │                       │                 │                     ┌────────────────┐
    │  └───────────────────────┤      FART       │◄────────────────────┤                │
    │                          │      CORE       │                     │     SERVER     │
    └─────────────────────────►│   INTERCEPTOR   ├────────────────────►│                │
                               │                 │                     └────────────────┘
                               │                 │
                               └─────┬─────┬─────┘
                                     │     │
                             CONTROL │     │ DATA
                             SOCKET  │     │ SOCKET
                                     │     │
                                     │     │
                               ┌─────┴─────┴─────┐
                               │                 │
                               │                 │
                               │                 │
                               │      FART       │
                               │     CLIENT      │
                               │                 │
                               │                 │
                               │                 │
                               └─────────────────┘
```

To allow interoperability with any client willing to integrate with Fart, the
core should expose the intercepted data with a websocket (such that the client
receives in an unsolicited way the intercepted calls) and it should allow the
client to control it at runtime via an HTTP enpoint (that may be bound both on a
ordinary socket, to allow remote control and to a unix socket, to allow local
control only).

## The server

The server should be able to read the configuration from the command line, from
environment variables, from configuration file and, lastly, using some default.
The intercepted data could be stored as simple text files, and could be
indicized using an embedded sqlite database.
The very same database could be used to store the filter statements for the
resources to be intercepted, leveraging the SQL engine to perform queries.

#### Interceptor abilities

The core should be able to expose itself as an ordinary HTTP proxy, but to be
able to intercept also the TLS traffic it should be able to act as a transparent
proxy, eavesdropping the connection data. Obviously, this requires either to
accept a self-signed certificate on the client side, or to deploy an appropriate
certificate chain on the client machine.
The core should also be able to proxy websocket messages (that are almost-raw
tcp).

## The client

The client should both allow easy data inspection/manipulation and direct core
configuration for the user.
In this repo the client will be a command line repl-like interface, that should
expose some set of commands to interact with the core. To easily inspect/modify
the requests/responses, it could spawn a pager/editor (like less/vim).
