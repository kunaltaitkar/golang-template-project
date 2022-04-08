# ---------------------------------------------------------------------------------
FROM golang:latest as dev
ENV GOMODULE111=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /golang-template-project
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# --------------------------------For development------------------------------------
FROM dev as debug
# install debugging tools
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cespare/reflex@latest
CMD reflex -R "__debug_bin" -s -- sh -c "dlv debug --headless --continue --accept-multiclient --listen :40000 --api-version=2 --log ./"
#------------------------------------Build-------------------------------------------
FROM dev as build
RUN go build -o  ./golang-template-project ./
# -------------------------------Test---------------------------------------------------
FROM dev as test
RUN go test -coverprofile=coverage.out -v $(go list ./src/... | grep -v test/)
RUN cat coverage.out | grep -v mock > /tmp/go-test-coverage.out \
   && go tool cover -func=/tmp/go-test-coverage.out
RUN cp /tmp/go-test-coverage.out /tmp/coverage.out
# ---------------------------------------production--------------------
FROM scratch as release
EXPOSE 8080
WORKDIR /app/
COPY --from=build /golang-template-project/golang-template-project ./
COPY --from=build /golang-template-project/app.env ./

ENTRYPOINT [ "/app/golang-template-project" ]
