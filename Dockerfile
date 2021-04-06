FROM golang:1.12

WORKDIR /bookstore_users-api

ADD . /bookstore_users-api
ENV GO111MODULE=on
ARG MYSQL_USERS_USERNAME
ARG MYSQL_USERS_PASSWORD
ARG MYSQL_USERS_HOST
ARG MYSQL_USERS_SCHEMA
ENV mysql_users_username=$MYSQL_USERS_USERNAME \
    mysql_users_password=$MYSQL_USERS_PASSWORD \
    mysql_users_host=$MYSQL_USERS_HOST \
    mysql_users_schema=$MYSQL_USERS_SCHEMA
RUN cd /bookstore_users-api
RUN go mod download
RUN go build

EXPOSE 8080

ENTRYPOINT ./bookstore_users-api