ARG APP_NAME=hello
ARG USER=appuser

############################
# STEP 1 build executable binary
############################
FROM golang:1.20-alpine as builder
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
ENV UID=10001
ARG USER
ARG APP_NAME
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/bin/main cmd/${APP_NAME}/main.go
############################
# STEP 2 build a small image
############################
FROM scratch
ARG USER
ARG APP_NAME
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /app/bin/main /app/bin/main
USER ${USER}:${USER}
ENTRYPOINT ["/app/bin/main"]
