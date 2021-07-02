FROM golang:latest

COPY server /app/server

CMD ["./app/","server"]
