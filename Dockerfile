FROM golang:1.12

WORKDIR /bookstore_users-api

ADD . /bookstore_users-api
ENV GO111MODULE=on
RUN cd /bookstore_users-api
RUN go mod download
RUN go build
EXPOSE 8080

ENTRYPOINT ./bookstore_users-api