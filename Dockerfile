FROM golang:alpine as builder
ADD . /titlecasebot
WORKDIR /titlecasebot
RUN go build

FROM alpine:latest
WORKDIR /titlecasebot
COPY --from=builder /titlecasebot/titlecasebot /titlecasebot/
ENTRYPOINT [ "/titlecasebot/titlecasebot" ]