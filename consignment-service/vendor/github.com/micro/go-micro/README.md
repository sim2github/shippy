# Go Micro [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-micro?status.svg)](https://godoc.org/github.com/micro/go-micro) [![Travis CI](https://api.travis-ci.org/micro/go-micro.svg?branch=master)](https://travis-ci.org/micro/go-micro) [![Go Report Card](https://goreportcard.com/badge/micro/go-micro)](https://goreportcard.com/report/github.com/micro/go-micro)

Go Micro is a pluggable RPC framework for **microservices**. It is part of the [Micro](https://github.com/micro/micro) toolkit.

The **Micro** philosophy is sane defaults with a pluggable architecture. We provide defaults to get you started quickly but everything can be easily swapped out. It comes with built in support for {json,proto}-rpc encoding, consul or multicast dns for service discovery, http for communication and random hashed client side load balancing.

Everything in go-micro is **pluggable**. You can find and contribute to plugins at [github.com/micro/go-plugins](https://github.com/micro/go-plugins).

Follow us on [Twitter](https://twitter.com/microhq) or join the [Slack](http://slack.micro.mu/) community.

## Features

Go Micro abstracts away the details of distributed systems. Here are the main features.

- **Service Discovery** - Automatic registration and name resolution with service discovery
- **Load Balancing** - Smart client side load balancing of services built on discovery
- **Synchronous Comms** - RPC based communication with support for bidirectional streaming
- **Asynchronous Comms** - PubSub interface built in for event driven architectures
- **Message Encoding** - Dynamic encoding based on content-type with protobuf and json out of the box
- **Service Interface** - All features are packaged in a simple high level interface for developing microservices

Go Micro supports both the Service and Function programming models. Read on to learn more.

## Docs

For more detailed information on the architecture, installation and use of go-micro checkout the [docs](https://micro.mu/docs).

## Learn By Example

An example service can be found in [**examples/service**](https://github.com/micro/examples/tree/master/service) and function in [**examples/function**](https://github.com/micro/examples/tree/master/function). 

The [**examples**](https://github.com/micro/examples) directory contains examples for using things such as middleware/wrappers, selector filters, pub/sub, grpc, plugins and much more. For the complete greeter example look at [**examples/greeter**](https://github.com/micro/examples/tree/master/greeter). Other examples can be found throughout the GitHub repository.

Watch the [Golang UK Conf 2016](https://www.youtube.com/watch?v=xspaDovwk34) video for a high level overview.

## Getting started

- [Service Discovery](#service-discovery)
- [Writing a Service](#writing-a-service)
- [Writing a Function](#writing-a-function)
- [Plugins](#plugins)
- [Wrappers](#wrappers)

## Service Discovery

Service discovery is used to resolve service names to addresses. It's the only dependency of go-micro.

### Consul

[Consul](https://www.consul.io/) is used as the default service discovery system. 

Discovery is pluggable. Find plugins for etcd, kubernetes, zookeeper and more in the [micro/go-plugins](https://github.com/micro/go-plugins) repo.

[Install guide](https://www.consul.io/intro/getting-started/install.html)

### Multicast DNS

[Multicast DNS](https://en.wikipedia.org/wiki/Multicast_DNS) is a built in service discovery plugin for a zero dependency configuration. 

Pass `--registry=mdns` to any command or the enviroment variable MICRO_REGISTRY=mdns

```
go run main.go --registry=mdns
```

## Writing a service

This is a simple greeter RPC service example

Find this example at [examples/service](https://github.com/micro/examples/tree/master/service).

### Create service proto

One of the key requirements of microservices is strongly defined interfaces. Micro uses protobuf to achieve this.

Here we define the Greeter handler with the method Hello. It takes a HelloRequest and HelloResponse both with one string arguments.

```proto
syntax = "proto3";

service Greeter {
	rpc Hello(HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
	string name = 1;
}

message HelloResponse {
	string greeting = 2;
}
```

### Install protobuf

Install [protobuf](https://developers.google.com/protocol-buffers/)

Now install the micro fork of protoc-gen-go. The protobuf compiler for Go. 

```shell
go get github.com/micro/protobuf/{proto,protoc-gen-go}
```

### Generate the proto

After writing the proto definition we must compile it using protoc with the micro plugin.

```shell
protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src \
	$GOPATH/src/github.com/micro/examples/service/proto/greeter.proto
```

### Write the service

Below is the code for the greeter service. 

It does the following:

1. Implements the interface defined for the Greeter handler
2. Initialises a micro.Service
3. Registers the Greeter handler
4. Runs the service

```go
package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	proto "github.com/micro/examples/service/proto"
	"golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
```

### Run service
```
go run examples/service/main.go
```

Output
```
2016/03/14 10:59:14 Listening on [::]:50137
2016/03/14 10:59:14 Broker Listening on [::]:50138
2016/03/14 10:59:14 Registering node: greeter-ca62b017-e9d3-11e5-9bbb-68a86d0d36b6
```

### Define a client

Below is the client code to query the greeter service. 

The generated proto includes a greeter client to reduce boilerplate code.

```go
package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	proto "github.com/micro/examples/service/proto"
	"golang.org/x/net/context"
)


func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))

	// Create new greeter client
	greeter := proto.NewGreeterClient("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
```

### Run the client

```shell
go run client.go
```

Output
```
Hello John
```

## Writing a Function

Go Micro includes the Function programming model. 

A Function is a one time executing Service which exits after completing a request. 

### Defining a Function

```go
package main

import (
	proto "github.com/micro/examples/function/proto"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// create a new function
	fnc := micro.NewFunction(
		micro.Name("go.micro.fnc.greeter"),
	)

	// init the command line
	fnc.Init()

	// register a handler
	fnc.Handle(new(Greeter))

	// run the function
	fnc.Run()
}
```

It's that simple.

## Plugins

By default go-micro only provides a few implementation of each interface at the core but it's completely pluggable. There's already dozens of plugins which are available at [github.com/micro/go-plugins](https://github.com/micro/go-plugins). Contributions are welcome!

### Build with plugins

If you want to integrate plugins simply link them in a separate file and rebuild

Create a plugins.go file

```go
import (
        // etcd v3 registry
        _ "github.com/micro/go-plugins/registry/etcdv3"
        // nats transport
        _ "github.com/micro/go-plugins/transport/nats"
        // kafka broker
        _ "github.com/micro/go-plugins/broker/kafka"
)
```

Build binary

```shell
// For local use
go build -i -o service ./main.go ./plugins.go
```

Flag usage of plugins
```shell
service --registry=etcdv3 --transport=nats --broker=kafka
```

## Wrappers

Go-micro includes the notion of middleware as wrappers. The client or handlers can be wrapped using the decorator pattern.

### Handler

Here's an example service handler wrapper which logs the incoming request

```go
// implements the server.HandlerWrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[%v] server request: %s", time.Now(), req.Method())
		return fn(ctx, req, rsp)
	}
}
```

It can be initialised when creating the service

```go
service := micro.NewService(
	micro.Name("greeter"),
	// wrap the handler
	micro.WrapHandler(logWrapper),
)
```

### Client

Here's an example of a client wrapper which logs requests made

```go
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[wrapper] client request to service: %s method: %s\n", req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}
```

It can be initialised when creating the service

```go
service := micro.NewService(
	micro.Name("greeter"),
	// wrap the client
	micro.WrapClient(logWrap),
)
```

## Other Languages

Check out [ja-micro](https://github.com/Sixt/ja-micro) to write services in Java

## Sponsors

Open source development of Micro is sponsored by Sixt

<a href="https://micro.mu/blog/2016/04/25/announcing-sixt-sponsorship.html"><img src="https://micro.mu/sixt_logo.png" width=150px height="auto" /></a>