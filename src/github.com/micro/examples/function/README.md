# Function

This is an example of creating a micro function. A function is a one time executing service.

## Contents

- main.go - is the main definition of the function
- proto - contains the protobuf definition of the API

## Prereqs

Micro services need a discovery system so they can find each other. Micro uses consul by default but 
its easily swapped out with etcd, kubernetes, or various other systems. We'll run consul for convenience.

Install consul
```shell
brew install consul
```

Run Consul

```shell
consul agent -dev
```

## Install micro

```
go get github.com/micro/micro
```

## Run the example

```shell
micro run -r github.com/micro/examples/function
```

## Call the function

```shell
micro query go.micro.fnc.greeter Greeter.Hello '{"name": "john"}'
```
