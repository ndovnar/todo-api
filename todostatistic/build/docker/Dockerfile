FROM golang:1.23.1-alpine AS init

# Required for confluent-kafka-go on Alpine Linux
ENV EXTRA_LINT_FLAGS=--build-tags=musl
ENV EXTRA_BUILD_FLAGS=-tags=musl

RUN apk --no-cache add make bash gcc musl-dev

WORKDIR /go/src/github.com/nvdovnar/todo-api

COPY todolib/go.mod todolib/go.sum ./todolib/
COPY todostatistic/go.mod todo/go.sum ./todostatistic/

RUN go work init ./todolib ./todostatistic
RUN go mod download

COPY todolib ./todolib
COPY todostatistic ./todostatistic

WORKDIR /go/src/github.com/nvdovnar/todo-api/todostatistic
RUN make build

FROM alpine:3 AS final

COPY --from=init /go/src/github.com/nvdovnar/todo-api/todostatistic/output/bin/todostatistic /bin/todostatistic

ENTRYPOINT ["/bin/todostatistic"]