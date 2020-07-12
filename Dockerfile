FROM golang:latest

LABEL maintainer="abhaytiwari8109 <at8109555785@gmail.com>"

WORKDIR /app

COPY go.mod  .

COPY go.sum  .

RUN go mod download

COPY . .

ENV PORT 8000

RUN go build

CMD ["./golang-rest-api"]