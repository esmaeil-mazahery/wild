## Build stage
FROM golang AS build-env

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /src

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy the code into the container
COPY . .

RUN go generate

RUN mkdir /app

# RUN go install
RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /app/api

## Production stage
FROM scratch

# Copy binary from build stage to Production stage
COPY --from=build-env /app/. /app

# ENV MONGO_PASSWORD=
ENV MONGO_USERNAME=admin
ENV MONGO_ADDRESS=mongo:27017
ENV MONGO_DBNAME=wild
ENV MONGO_TIMEOUT=60000

ENV REDIS_URI=192.168.1.101:6379
# ENV REDIS_PASSWORD
ENV REDIS_TIMEOUT=60000
ENV REDIS_PERFIX=wild_

ENV HTTP=true
ENV GRPC_PORT=90
ENV GETEWAY_PORT=80

# ENV SERVICE_SECRET_KEY=
ENV SERVICE_TOKEN_LOGIN_DURATION=2592000
ENV SERVICE_TOKEN_VERIFY_DURATION=300

# ENV EMAIL_FROM=
# ENV EMAIL_PASSWORD=
ENV EMAIL_SMTPHOST=smtp.gmail.com
ENV EMAIL_SMTPPORT=587

# Export necessary port
EXPOSE 80

# Command to run when starting the container
ENTRYPOINT ["/app/api"]