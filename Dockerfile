FROM golang:1.25

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 80

CMD ["go", "run", "."]
