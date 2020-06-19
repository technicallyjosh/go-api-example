# go-api-example

Example of a go API with a password validation endpoint and tests.

## Dependencies

- [Go 1.14+](https://golang.org/doc/install)

## Local

_build and run binary..._

```console
$ go build && ./go-api-example
```

_...or you can run it without compiling it..._

```console
$ go run .
```

_...and now you can call it..._

```console
$ curl -i -XPUT localhost:8000/users/verify -d '
{
    "username": "technicallyjosh",
    "password": "testing123"
}'; echo
```

## Running Tests

```console
$ go mod download
$ go test
```

## With Docker

**Build it**

```console
$ docker build -t go-api-example .
```

**Run server**

```console
$ docker run --rm go-api-example
```

**Run tests**

```console
$ docker run --rm go-api-example /bin/bash -c "go test"
```
