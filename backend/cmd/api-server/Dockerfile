FROM golang:1.10-alpine as build
RUN apk update && apk upgrade && apk add git curl

# Copy source
WORKDIR /go/src/api-server
COPY . .

# Install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Install dependencies
RUN dep ensure

# Build application
RUN go build

FROM alpine:3.6 as run
WORKDIR /root/
COPY --from=build /go/src/chat/api-server api-server
CMD ["./api-server"]
