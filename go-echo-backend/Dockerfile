FROM golang:1.11.4

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV MAXSTREAM=1000

CMD go run go_echo_backend.go -version $VERSION -maxstream $MAXSTREAM -certfile $CERTFILE -keyfile $KEYFILE
