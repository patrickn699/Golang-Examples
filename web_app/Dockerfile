FROM golang:1.17
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go get -u github.com/gorilla/mux
RUN go mod init go_webs
RUN go build -o webb .
CMD ["/app/webb"]


