# ------------------------------------------------------------------
FROM golang:latest as dev
ENV GOMODULE111=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/github.com/kunaltaitkar/golang-template-project
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# --------------------------------------------------------------------
FROM dev as debug

# install debugging tools
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go get github.com/cespare/reflex
RUN which reflex

# install any project dependencies
RUN go get -d -v ./...

CMD reflex -R "__debug_bin" -s -- sh -c "dlv debug --headless --continue --accept-multiclient --listen :40000 --api-version=2 --log ./"