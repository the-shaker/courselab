FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM scratch
COPY --from=build /app/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]
