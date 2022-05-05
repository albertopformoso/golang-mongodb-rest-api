FROM golang:1.18.1-alpine3.15
RUN mkdir /src
ADD . /src
WORKDIR /src
RUN go build -o main .
EXPOSE 8000
CMD ["/src/main"]