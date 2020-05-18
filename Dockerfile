FROM golang:1.14.2 as builder
ARG SOURCE_LOCATION=/
WORKDIR ${SOURCE_LOCATION}
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
ARG SOURCE_LOCATION=/
RUN apk --no-cache add curl
EXPOSE 8080
WORKDIR /root/
COPY --from=builder ${SOURCE_LOCATION} .
ENTRYPOINT [ "./app" ]
