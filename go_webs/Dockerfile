FROM golang:1.17
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod init go_webs
RUN go build -o ws .
CMD ["/app/ws"]
