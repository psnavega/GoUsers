FROM golang:1.20

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

EXPOSE 5001

COPY . .

CMD ["go", "run", "."]