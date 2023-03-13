FROM golang:1.19-alpine

RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/admin-server
RUN go build -o /server

CMD /server