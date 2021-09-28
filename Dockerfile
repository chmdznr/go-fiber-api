##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .
COPY env.docker .env

RUN go build -o /api

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /api /api
COPY --from=build /app/.env /.env

EXPOSE 5001

USER nonroot:nonroot

ENTRYPOINT ["/api"]