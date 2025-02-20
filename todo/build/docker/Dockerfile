FROM golang:1.23.1-alpine3.20 AS init

RUN apk --no-cache add make bash

WORKDIR /go/src/github.com/nvdovnar/todo-api

COPY todolib/go.mod todolib/go.sum ./todolib/
COPY todo/go.mod todo/go.sum ./todo/

RUN go work init ./todolib ./todo
RUN go mod download

COPY todolib ./todolib
COPY todo ./todo

WORKDIR /go/src/github.com/nvdovnar/todo-api/todo
RUN make build

FROM scratch AS final

COPY --from=init /go/src/github.com/nvdovnar/todo-api/todo/output/bin/todo /bin/todo

ENTRYPOINT ["/bin/todo"]