FROM golang:1.17
WORKDIR /app
COPY . .
RUN go build -o fruit-operator
CMD ["./fruit-operator"]