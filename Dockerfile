FROM golang:buster
# Build our app
WORKDIR /usr/app
COPY . .

RUN go build main.go

CMD ["./main"]