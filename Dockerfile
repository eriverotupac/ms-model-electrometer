FROM golang:1.20-alpine AS build_base

RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/electro_app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -o ./out/electro_app ./cmd/main.go

# Start fresh from a smaller image
FROM gcr.io/distroless/static

COPY --from=build_base /tmp/electro_app/out/electro_app /app/elec_data_fetch

#COPY /vault/.env /app/elec_data_fetch/.env

# Expose the container to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/elec_data_fetch"]
