FROM golang:1.21 AS build
COPY ./ /go/src/github.com/jake-willog/go-k8s-api/
WORKDIR /go/src/github.com/jake-willog/go-k8s-api/

ENV CGO_ENABLED=0
ARG GIT_AUTH_USR
ARG GIT_AUTH_PSW

RUN go build -a -installsuffix cgo -o go-k8s-api .

FROM alpine AS runtime
COPY --from=build /go/src/github.com/jake-willog/go-k8s-api/go-k8s-api ./
EXPOSE 8080/tcp
ENTRYPOINT ./go-k8s-api run
