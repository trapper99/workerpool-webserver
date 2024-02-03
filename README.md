A Golang WorkerPool for distributed tasks
==========================================

## Overview

Go's implementation of web server already uses go routines to run
handlers. However, it is not possible to run multiple handlers
simultaneously. The `WorkerPool` is a simple implementation of a
worker pool that can be used to run multiple handlers in parallel.

This is a simple Project which provides a way for the web server
to provide instant response but process the heavy work in parallel in the background.

## Getting Started

1. Clone the repository

```
git clone https://github.com/trapper99/workerpool-webserver.git
```

```
cd workerpool-webserver
```

2. Run the server

```
go run ./server/main.go
```

3. Run the client to test it

```
go run ./client/main.go
```

The server will print out the status of the server and the client will print out the status of the client.
