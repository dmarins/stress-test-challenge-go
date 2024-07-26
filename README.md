# Stress Test - FullCycle Go Expert Challenge

https://goexpert.fullcycle.com.br/pos-goexpert/

[![Go](https://img.shields.io/badge/go-1.22.5-informational?logo=go)](https://go.dev)

## Clone the project

```
$ git clone https://github.com/dmarins/stress-test-challenge-go.git
$ cd stress-test-challenge-go
```

## Download dependencies

```
$ go mod tidy
```

## Run tests

```
$ make tests
```

## Build local docker image

```
$ docker build -t stress-test .
```

## Run local docker image

```
$ docker run stress-test --url=http://google.com --requests=30 --concurrency=10
```

## Run remotely docker image

```
$ docker run diogomarins/stress-test:latest --url=http://example.com --requests=20 --concurrency=10
```