FROM golang:1.17.0-alpine3.14

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git && go get -u github.com/cosmtrek/air && go get -u github.com/go-delve/delve/cmd/dlv

COPY ./main.go ${ROOT}
COPY ./cmd/ ${ROOT}
COPY .air.toml ${ROOT}
COPY go.mod ${ROOT}

RUN go clean
RUN go mod tidy
CMD ["air", "-c", ".air.toml"]
