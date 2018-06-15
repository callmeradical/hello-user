FROM golang:1.9.3 as builder

COPY . /go/src/github.com/callmeradical/hello-user
WORKDIR /go/src/github.com/callmeradical/hello-user
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

COPY --from=builder /go/src/github.com/callmeradical/hello-user/app .

CMD ["./hello-user"]
