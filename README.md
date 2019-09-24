# node-gotraceroute
A node module for performing traceroutes. The majority of functionality is
provided by Golang via golang.org/x/net/icmp and golang.org/x/net/ipv4. The Go
module exports a single function visible to C++ which provides the glue
to provide interop with Node.

A trace is performed by calling pingHost() sequentially with increasing TTL
values. pingHost returns a JSON string that contains information about the
discovered hop (or error).

> As of the creation of this C++ (and vicariously, Node) can't maintain a handle
> to a Go object between function calls thanks to garbage collection. This isn't   
> a problem as far as we're concerned since objects need only live long enough to
> discover a single new server.

## Requirements
```
go get golang.org/x/net/icmp
golang.org/x/net/ipv4
```

## Build Instructions
```
go build -buildmode c-archive -o module.a module.go
node-gyp configure
node-gyp build
```

### Why?
When looking for a traceroute module I couldn't find any that allowed me to
cancel the trace midway through. With long timeouts that meant that I wouldn't
be able to build a usable UI.

I chose Go + C++ + Node because it provided an interesting challenge. I hold no
expectations that this is an efficient solution to running traces.
