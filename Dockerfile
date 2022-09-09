FROM golang:1.19-alpine AS build

# ENV GOPROXY=https://https://goproxy.cn

# Move to working directory (/build).
WORKDIR /app

# Copy and download dependency using go mod.
COPY go.mod ./
COPY go.sum ./
COPY . .

# RUN go mod tidy
RUN go mod vendor
RUN go build -o server

FROM scratch

WORKDIR /
# Copy binary and config files from /build to root folder of scratch container.
# TODO: remover .env
COPY --from=build ["/app/server", "/.env"]

EXPOSE 5000

USER nonroot:nonroot

# Command to run when starting the container.
ENTRYPOINT ["/server"]
