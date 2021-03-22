# Roadmap

Let's try to produce something in steps:

 - [x] Simple HTTP proxy that logs all requests/responses
 - [x] Simple CLI for the server, to specify some connection parameters
 - [ ] Expose the captured data via websocket
 - [ ] A simple client to read and pipe through a pager data from the server
 - [ ] Filter the requests/responses using some config-provided expression
 - [ ] Add the database and store the filter expression therein
 - [ ] Add the control socket on the server
 - [ ] Implement a repl for the client
 - [ ] Pipe the requests/responses on the client to an editor and send them back
   via the server
 - [ ] Add the ability to persist data at the server
 - [ ] Add the transparent proxy capability
 - [ ] Add the websocket proxy capability
 - [ ] Add a dummy repeater on the server (to be able to repeat the requests)
 - [ ] Add a template render engine to make the repeater a bit smarter (and
   allow basic fuzz)
 - [ ] Add a diff view on the client for responses coming from repeated requests
 - [ ] Add some way to integrate request/response processig with external
   processes (plugins)
 - [ ] Hack the world
