# Graceful shutdown in HTTP server

This is an example of graceful shutdown.

The server has a handler that server 1GB of repeated string 'A', this is to ensure we can have a long lived connection when shutting down.

There are two main components here:
- goroutine that starts the server
- main routine that listens for signals afterwards

Once the server is up and running, the main routine can receive signals, triggering the server shutdown

A context is used for some timeout in case the requests are not fullfilled without that window.

A channel `gracefulShutdownChan` is used to sync the output between the two routines.

# Sources
- https://dev.to/mokiat/proper-http-shutdown-in-go-3fji
- https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/

