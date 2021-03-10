# Resources

### Architecture

The general architecture is heavily borrowed from [circolog][c]. The
client/server architecture and the data websocket/http control socket are
copy-pasted from it.

### Proxying

To avoid reinventing the wheel, this MARVELOUS library for proxying could be
used: [github.com/elazarl/goproxy][p].

It has a ton of example and allows to easily implement a HTTP(S) proxy, a
websocket proxy, a transparent proxy, together with a programmable logic (handy
for filtering the intercepted request based on some logic).

### Data and configuration

I know it might seem overkill, but an old good sqlite embedded database could be
directly leveraged when matching which requests/responses to intercept and which
to let pass.
The de facto standard is [github.com/mattn/go-sqlite3][s] that implements the
classic `database/sql` interface.

### CLI

I would stick to the golang top notch standard:

 - [github.com/spf13/cobra][c1] for the CLI
 - [github.com/spf13/viper][c2] for the config files

They play very well together and, most importantly, do generate shell completion
for free.

### TUI

For the client interface, I have very clearly in mind the UX from
[github.com/bettercap/bettercap][b]. It is a repl with a custom set of commands.

### TLS

TLS is hard; TLS done properly is crazy. Fortunately, there are a bunch of crazy
people out there that provided us with good tools to handle tough stuff. One of
such tools is [mkcert][k], to generate and install a certificate chain. This
should allow one to perform transparent proxying properly.



[c]: https://git.lattuga.net/boyska/circolog
[p]: https://github.com/elazarl/goproxy
[s]: https://github.com/mattn/go-sqlite3/
[c1]: https://github.com/spf13/cobra
[c2]: https://github.com/spf13/viper
[p]: https://github.com/bettercap/bettercap
[k]: https://github.com/FiloSottile/mkcert
