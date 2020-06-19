# For simplicity I'm including source and built in here so you can run tests as well...
FROM golang:1.14
LABEL maintainer="technicallyjosh"

WORKDIR /app

ADD . /app

RUN go mod download
RUN go build .

EXPOSE 8000

CMD ["./go-api-example"]
