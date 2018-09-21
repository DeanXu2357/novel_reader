FROM golang:1.11

ARG TIMEZONE

RUN \
    apt-get update &&\
    apt-get upgrade -y &&\
    apt-get install -y git vim curl

RUN apt-get -y install software-properties-common

# Set timezone
RUN ln -snf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime && echo ${TIMEZONE} > /etc/timezone
RUN "date"

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go run -o main .
CMD ["/app/main"]

EXPOSE 8080
