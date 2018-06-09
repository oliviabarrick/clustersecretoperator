FROM golang:1.10.1-alpine3.7

WORKDIR /go/src/github.com/justinbarrick/clustersecretoperator

COPY . ./
RUN CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o clustersecretoperator cmd/clustersecretoperator/main.go

FROM scratch
COPY --from=0 /go/src/github.com/justinbarrick/clustersecretoperator/clustersecretoperator /clustersecretoperator
ENTRYPOINT ["/clustersecretoperator"]
