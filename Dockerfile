FROM golang:1.12

WORKDIR /bookstore_users-api

ADD . /bookstore_users-api
ENV GO111MODULE=on
ENV mysql_users_username=root
ENV	mysql_users_password=1234
ENV mysql_users_host=mysql
ENV mysql_users_schema=users_db


RUN cd /bookstore_users-api
RUN go mod download
RUN go build
EXPOSE 8080

ENTRYPOINT ./bookstore_users-api