FROM golang:1.21 AS build
WORKDIR /go/src/github.com/jake-willog/go-k8s-api/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o go-k8s-api .

FROM alpine AS runtime
COPY --from=build /go/src/github.com/jake-willog/go-k8s-api/go-k8s-api ./
EXPOSE 8080/tcp
ENTRYPOINT ["./go-k8s-api", "run"]
