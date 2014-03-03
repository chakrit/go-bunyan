
![build status](https://api.travis-ci.org/chakrit/go-bunyan.png)

# GO-BUNYAN

`go-bunyan` is a `node-bunyan` compatible logger implementation written in Go1.2.

In short, this is a simple structured JSON logger. Logs outputs are JSON records that are
easy to further parse/query and mine. The format is inspired by the [node-bunyan][0] tool.
However the implementation and configuration is purely from the ground-up as Go1.2 is
obviously not JavaScript.

This thing is a work-in-progress, so expect some bugs and some missing features. But do
feel free to try it out and file bugs and feature requests.

# HOWTO

Use the `NewStdLogger` method to create a new, standard logger:

```go
import "github.com/chakrit/go-bunyan"

var Log = bunyan.NewStdLogger("appname", bunyan.StdoutSink())
```

Log with one of `Tracef`, `Debugf`, `Infof`, `Warnf`, `Errorf` and `Fatalf`. Include
additional information to the logger with the `Reocrd` or `Recordf` method. And finally,
create child loggers with the `Child()` method. Here are some examples:

```go
address := ":8080"

// simple record with one extra value:
Log.Record("address", address).Infof("starting server...")

// record with fmt-style args
Log.Info("starting server at `%s`...", address)

// create child loggers with added info
child := Log.Record("address", address).Child()
child.Info("starting server...")
```

More documentation will be available once this library is more complete.

[0]: http://github.com/trentm/node-bunyan

