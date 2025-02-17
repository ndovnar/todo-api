FROM golang:1.23.1-alpine3.20 AS init

RUN apk --no-cache add make bash

WORKDIR /go/src/github.com/nvdovnar/todo-api

COPY todolib/go.mod todolib/go.sum ./todolib/
COPY todoauth/go.mod todoauth/go.sum ./todoauth/

RUN go work init ./todolib ./todoauth
RUN go mod download

COPY todolib ./todolib
COPY todoauth ./todoauth

WORKDIR /go/src/github.com/nvdovnar/todo-api/todoauth
RUN make build

FROM scratch AS final

COPY --from=init /go/src/github.com/nvdovnar/todo-api/todoauth/output/bin/todoauth /bin/todoauth

ENTRYPOINT ["/bin/todoauth"]