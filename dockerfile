FROM golang:1.22 AS base

# Install dependencies and build executable file
FROM base AS builder
RUN mkdir /app
WORKDIR /app
COPY . .

# Install dependencies based on the preferred package manager
RUN go mod tidy

# Build
RUN CGO_ENABLE=0 GOOS=linux go build -ldflags="-s -w" -o ./build/event-processor ./cmd

FROM ubuntu:latest as runner
RUN apt-get update && apt-get install -y tzdata
ENV TZ="America/Bahia"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN mkdir /app

COPY --from=builder /app/build/ /app
# # Run
CMD ["/app/event-processor"]