FROM golang:1.14.2
LABEL maintainer="ldclakmal@gmail.com"

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENV MAXSTREAM=1000

CMD go run go_echo_backend.go -version $VERSION -maxstream $MAXSTREAM
