# go-api-example

Example of a go API with a password validation endpoint and tests. The service runs on port 8000.

The correct credentials are:

**username**: technicallyjosh<br>
**password**: testing123

## Dependencies

- [Go 1.14+](https://golang.org/doc/install)

## Local

**Install Dependencies**

```console
$ go mod download
```

**Build and Run**

```console

$ go build && ./go-api-example
```

**Run**

```console
$ go run .
```

**Test**

```console
$ go test
```

## In Docker

**Build**

```console
$ docker build -t go-api-example .
```

**Run server**

```console
$ docker run --rm -p 8000:8000 go-api-example
```

## Running Tests

**Local**

```console
$ go mod download
$ go test
```

**Docker**

```console
$ docker run --rm go-api-example /bin/bash -c "go test"
```

## Routes

### `PUT /users/verify`

Verifies a password for a user. In this example it's more like a login without assigning a token or
session etc... Normally when just _verifying_ a user's password, we'd accept a signed token in the
header and get the username or id from it.

#### Request Body (JSON)

| Property | Type     | Required |
| -------- | -------- | -------- |
| username | `string` | Yes      |
| password | `string` | Yes      |

#### Example

```console
$ curl -i -XPUT localhost:8000/users/verify -d '
{
    "username": "technicallyjosh",
    "password": "testing123"
}'; echo
```
