# Compile stage
FROM golang:1.19.2 AS build-env
ADD . /dockerdev
WORKDIR /dockerdev
RUN go mod download
RUN go build -o /gservice cmd/server/main.go

FROM debian:buster
EXPOSE 8080
WORKDIR /
COPY --from=build-env /gservice /
CMD ["/gservice"]
