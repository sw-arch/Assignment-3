FROM golang

WORKDIR /go/src/Assignment-3
COPY src .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]

