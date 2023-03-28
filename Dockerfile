FROM golang:1.20-alpine
RUN mkdir /projectSync
WORKDIR /projectSync
COPY . .
RUN go mod vendor && go build -mod=vendor .
CMD ["./app"]