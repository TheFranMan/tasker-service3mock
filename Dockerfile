FROM golang:1.23 AS base-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /service3

FROM alpine:3.14 AS release
WORKDIR /
COPY --from=base-stage /service3 /service3
EXPOSE 3003
ENTRYPOINT [ "/service3" ]