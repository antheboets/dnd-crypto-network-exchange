FROM golang:1.16-buster AS build

WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download

COPY *.go ./

#RUN go build testserver.go

RUN go build testserver.go
EXPOSE 8081
CMD ["/app/testserver"]