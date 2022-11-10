FROM golang:1.19-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/rpl-discordbot

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/rpl-discordbot .

# Start fresh from a smaller image
FROM alpine:3.16

COPY --from=build_base /tmp/rpl-discordbot/out/rpl-discordbot /app/rpl-discordbot

# Run the binary program produced by `go install`
CMD ["/app/rpl-discordbot"]