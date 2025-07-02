FROM golang:1.24 AS build
WORKDIR /go/src
COPY cmd ./cmd
COPY pkg ./pkg
COPY go.mod .
COPY go.sum .

WORKDIR /go/src/cmd
ENV CGO_ENABLED=0
RUN go build -o savac

FROM scratch AS runtime
COPY --from=build /go/src/cmd/savac /
ENTRYPOINT ["/savac"]
