FROM golang:1.12.5-alpine

WORKDIR /go/src/nats-stream-cqrs-sales-discount
COPY . .

RUN go install


CMD [ "nats-stream-cqrs-sales-discount" ]

