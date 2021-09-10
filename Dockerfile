FROM golang:1.17 AS build

ARG GIT_DESC=undefined

WORKDIR /go/src/github.com/mysteriumnetwork/ethcheck
COPY . .
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-s -w -extldflags "-static" -X main.version='"$GIT_DESC" ./cmd/ethcheck

FROM scratch
COPY --from=build /go/src/github.com/mysteriumnetwork/ethcheck/ethcheck /
USER 9999:9999
ENTRYPOINT ["/ethcheck"]
