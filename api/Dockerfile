FROM golang:1.26.3-alpine3.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o api .

FROM scratch
COPY --from=build /app/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
