FROM golang:1.12.6-alpine3.10

RUN adduser -D go_runner
WORKDIR /go/src/sensor_board

COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go" "run" "."]


