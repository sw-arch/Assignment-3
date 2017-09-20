FROM golang

RUN apt-get update && apt-get install sqlite3

WORKDIR /go/src/Assignment-3
COPY . .

RUN go-wrapper download
RUN go-wrapper install

RUN sqlite3 inventory.db<initialDBTables/inventory.txt && \
    sqlite3 users.db<initialDBTables/users.txt

CMD ["go-wrapper", "run"]
