FROM golang:alpine

WORKDIR  /app 

COPY . .   

RUN ["go","run","cmd/main.go"]