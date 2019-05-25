# node-gotraceroute
A node module for performing traceroutes. Parameters of the trace (timeout, maxhops, etc) are configurable. Feedback is provided in a JSON payload for eash parsing, display.

## Build Instructions
```
go build -buildmode c-archive -o module.a module.go
node-gyp configure
node-gyp build
```
