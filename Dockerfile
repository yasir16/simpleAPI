FROM golang:alpine

WORKDIR /build


COPY go.mod .
COPY go.sum .
RUM go mod download

COPY . . 

RUM go build -o main .


WORKDIR /dist

RUN cp /build/main .

EXPOSE 3000

CMD ["/dist/main"]
