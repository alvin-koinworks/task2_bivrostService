# FROM golang:latest as builder
# LABEL maintainer="tonny.adhi@koinworks.com"

# ENV GO111MODULE=on

# ENV GOPRIVATE=github.com/koinworks

# ARG GITHUB_USERNAME
# ARG GITHUB_ACCESS_TOKEN

# ENV GITHUB_USERNAME="alvin-koinworks"
# ENV GITHUB_ACCESS_TOKEN="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"

# RUN echo "machine github.com login $GITHUB_USERNAME password $GITHUB_ACCESS_TOKEN">~/.netrc

# ENV APP asgard-example-service

# WORKDIR /app

# COPY ..
# RUN go mod download 

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/${APP} main.go
# EXPOSE ${PORT}
# ENTRYPOINT /out/${APP}

FROM golang:latest as builder
LABEL maintainer="tonny.adhi@koinworks.com"

ENV GO111MODULE=on

ENV GOPRIVATE=github.com/koinworks

ARG GITHUB_USERNAME
ARG GITHUB_ACCESS_TOKEN

RUN echo "machine github.com login $GITHUB_USERNAME password $GITHUB_ACCESS_TOKEN" > ~/.netrc

ENV APP api-example-service

WORKDIR /app

COPY . .
RUN go mod download

EXPOSE ${PORT}
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/${APP} main.go
ENTRYPOINT /out/${APP}
