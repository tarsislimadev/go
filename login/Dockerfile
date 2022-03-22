FROM golang:1.18

WORKDIR /usr/src/app

COPY . .

# RUN  go mod init github.com/brtmvdl/docker-go-projects

RUN go build -v -o /usr/local/bin/app ./...

CMD app
