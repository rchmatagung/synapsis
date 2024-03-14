#build stage
FROM golang:alpine AS builder

# Install dependencies to be able to build our code
RUN apk add --no-cache --update git && apk add build-base

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY ./go.mod ./go.sum ./
COPY ./cmd/api/main.go ./

RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./ .

# Compile the code
RUN go build -o main

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates && apk --no-cache add tzdata

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/config/config-dev.yaml ./config/config-dev.yaml
# COPY ./migrations ./migrations

# RUN apk add --no-cache postgresql-client

CMD ["./main"]
# CMD ["sh", "-c", "psql -h db -U postgres -d synapsis -a -f migrations/migrate.sql && ./main"]
