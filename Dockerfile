FROM golang:1.21 as builder
ENV ROOT /go/src/urlmap-front
WORKDIR ${ROOT}
COPY go.mod go.sum ./
RUN go mod download
COPY pb/ ./pb/
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main

FROM golang:1.21 as main
# FROM scratch as main
ENV TZ Asia/Tokyo
COPY --from=builder /go/bin/main /main
USER nobody

CMD ["/main"]