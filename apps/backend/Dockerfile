
##############################
# Stage 1: Build (for both dev and prod)
FROM golang:1.24.2-alpine3.21 AS builder

ARG FRONTEND_URL
ENV FRONTEND_URL=$FRONTEND_URL

RUN apk add --no-cache git

ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest
RUN go build -o server .

##############################
# Stage 2: Development container (with Air)
FROM golang:1.24.2-alpine3.21 AS development

WORKDIR /app
RUN apk add --no-cache git

COPY --from=builder /app /app
COPY --from=builder /go/bin/air /usr/bin/air

EXPOSE 80
CMD ["air"]

##############################
# Stage 3: Production container (minimal image)
FROM alpine:3.21 AS production

WORKDIR /root/
COPY --from=builder /app/server .

EXPOSE 80
CMD ["./server"]
