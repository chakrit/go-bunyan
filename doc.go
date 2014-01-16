// Package bunyan provides a structured JSON logger implementation. With default
// configuration, output from go-bunyan can be piped into node-bunyan
// (https://github.com/trentm/node-bunyan) for formatting and filtering.
//
// Hello World Example
//
// Use NewStdLogger to creates a basic logger give it StdoutSink so logs are written to
// the process's standard output:
//
//     package main
//
//     import "time"
//     import "github.com/chakrit/go-bunyan"
//
//     func main() {
//         logger := bunyan.NewStdLogger("test", bunyan.StdoutSink())
//
//         for t := range time.NewTicker(time.Second).C {
//             logger.Infof("Hello! It's %s!", t.Format(time.Kitchen)))
//             logger.Tracef("Random trace message.")
//         }
//     }
//
// Running the example program produce the following output (truncated for brevity):
//
//     $ ./example
//     {"hostname":"cx24.local","level":30,"msg":"Hello! It's 9:45AM!","name":"test","pid":25593,"time...
//     {"hostname":"cx24.local","level":10,"msg":"Random trace message.","name":"test","pid":25593,"ti...
//     {"hostname":"cx24.local","level":30,"msg":"Hello! It's 9:45AM!","name":"test","pid":25593,"time...
//     {"hostname":"cx24.local","level":10,"msg":"Random trace message.","name":"test","pid":25593,"ti...
//
// Example Formatted Output
//
// This package will ultimately provides its own CLI for formatting the log output but for
// now using the node-bunyan CLI tool which works with the same log format should suffice:
//
//     $ npm install -g bunyan
//
// Non-colored output looks like the following:
//
//     $ ./example | bunyan
//     [2014-01-16T09:46:28+07:00]  INFO: test/25762 on cx24.local: Hello! It's 9:46AM!
//     [2014-01-16T09:46:28+07:00] TRACE: test/25762 on cx24.local: Random trace message.
//     [2014-01-16T09:46:29+07:00]  INFO: test/25762 on cx24.local: Hello! It's 9:46AM!
//     [2014-01-16T09:46:29+07:00] TRACE: test/25762 on cx24.local: Random trace message.
//
// You can also do filtering by level via the CLI tool:
//
//     $ ./example | bunyan -l info
//     [2014-01-16T09:46:28+07:00]  INFO: test/25762 on cx24.local: Hello! It's 9:46AM!
//     [2014-01-16T09:46:29+07:00]  INFO: test/25762 on cx24.local: Hello! It's 9:46AM!
//
// Or pass in a JS condition string:
//
//     $ ./example | bunyan -c 'level == 10'
//     [2014-01-16T09:46:28+07:00] TRACE: test/25762 on cx24.local: Random trace message.
//     [2014-01-16T09:46:29+07:00] TRACE: test/25762 on cx24.local: Random trace message.
//
// The ability to filter logs via a JS condition becomes immensely useful later on once
// you start recording extra information into your log messages and having accumulated
// a lot of records.
package bunyan
