FROM golang

WORKDIR /go/src/Assignment-3
COPY . .

RUN go-wrapper download Assignment-3/src
RUN go-wrapper install Assignment-3/src

CMD ["go-wrapper", "run"]
