FROM golang AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download


COPY resources ./resources/
COPY *.go ./

#RUN go build testserver.go

RUN go build server.go
EXPOSE 8081
CMD ["/app/server"]