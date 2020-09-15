# builder image
FROM golang:1.15.2-alpine3.12 as builder
RUN mkdir /build
ADD / /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o superclog github.com/lwaddicor/superclog/cmd/superclog
# generate clean, final image for end users
FROM alpine:3.12.0

COPY --from=builder /build/superclog .

# executable
ENTRYPOINT [ "ls" ]
# arguments that can be overridden
CMD [ "--help" ]