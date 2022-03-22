FROM golang:1.18

ENV HOSTNAME=

ENV PORT=80

EXPOSE 80

WORKDIR /usr/src/app

COPY . .

# RUN  go mod init github.com/brtmvdl/docker-go-projects

RUN go build -v -o /usr/local/bin/app ./...

CMD app
