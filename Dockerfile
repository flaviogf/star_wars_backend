FROM golang:1.16
WORKDIR /go/src/app
COPY . .
RUN go build ./cmd/server
EXPOSE 80
ENTRYPOINT [ "./server" ]
