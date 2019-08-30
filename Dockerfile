# build
FROM golang:1.12.9-alpine3.10 as build

# default port of auto devops is 5000
ENV PORT 5000
EXPOSE 5000

RUN mkdir /app
ADD . /app

ENV GOPROXY https://goproxy.io
ENV GIN_MODE release

WORKDIR  /app
RUN go mod vendor
RUN go build -mod=vendor -tags=jsoniter -o golang-auto-devops .


# release
FROM alpine:3.10
RUN mkdir /app
COPY --from=build /app/golang-auto-devops /app/golang-auto-devops

WORKDIR  /app
CMD ["/app/golang-auto-devops"]
