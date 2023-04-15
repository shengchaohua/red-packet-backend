FROM golang:1.18

MAINTAINER shengchaohua

COPY . /

EXPOSE 8080

ENTRYPOINT ["go", "run", "./bin/red-packet-backend-admin"]
