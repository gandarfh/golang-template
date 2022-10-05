FROM golang:1.19-alpine AS build

# ENV GOPROXY=https://https://goproxy.cn

# Move to working directory (/build).
WORKDIR /app

# Copy and download dependency using go mod.
# COPY go.mod ./
# COPY go.sum ./
COPY . .

# RUN go mod tidy
RUN go mod tidy
RUN go build -o server

FROM alpine:latest

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=build /app/server /server

EXPOSE 5000

# Command to run when starting the container.
ENTRYPOINT [ "/server" ]
