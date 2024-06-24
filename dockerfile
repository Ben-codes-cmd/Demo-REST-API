# https://docs.docker.com/language/golang/build-images/

FROM golang:1.22.3

WORKDIR /app

COPY main/go.mod main/go.sum ./
RUN go mod download

COPY main/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /member-api

EXPOSE 8081

CMD ["/member-api"]

