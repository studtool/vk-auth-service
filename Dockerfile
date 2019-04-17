FROM golang:1.12-alpine3.9 as base
WORKDIR /tmp/vk-auth-service
COPY . .
RUN go build -mod vendor -o /tmp/service .

FROM ubuntu:18.04
WORKDIR /tmp
COPY --from=base /tmp/service ./service
ENTRYPOINT ./service
EXPOSE 80
