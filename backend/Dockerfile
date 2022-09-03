FROM golang:1.17.0-alpine3.14

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git && go get -u github.com/cosmtrek/air && go get -u github.com/go-delve/delve/cmd/dlv

COPY ./main.go ${ROOT}
COPY ./cmd/ ${ROOT}/cmd/
COPY .air.toml ${ROOT}
COPY go.mod ${ROOT}
COPY go.sum ${ROOT}

RUN go mod tidy
RUN ls

CMD ["air", "-c", ".air.toml"]
