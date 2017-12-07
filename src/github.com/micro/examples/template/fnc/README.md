# Template Fnc

This is the Template function with fqdn go.micro.fnc.template.

## Getting Started

### Service Discovery

Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -dev
```
### Micro Toolkit

Install Micro

```
go get github.com/micro/micro
```

### Run Function

```
$ micro run -r github.com/micro/examples/template/fnc
```

### Building a container

If you would like to build the docker container do the following
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o template-fnc ./main.go
docker build -t template-fnc .

```