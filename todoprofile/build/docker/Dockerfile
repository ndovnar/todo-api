FROM golang:1.23.1-alpine3.20 AS init

RUN apk --no-cache add make bash

WORKDIR /go/src/github.com/nvdovnar/todo-api

COPY todolib/go.mod todolib/go.sum ./todolib/
COPY todoprofile/go.mod todoprofile/go.sum ./todoprofile/

RUN go work init ./todolib ./todoprofile
RUN go mod download

COPY todolib ./todolib
COPY todoprofile ./todoprofile

WORKDIR /go/src/github.com/nvdovnar/todo-api/todoprofile
RUN make build

FROM scratch AS final

COPY --from=init /go/src/github.com/nvdovnar/todo-api/todoprofile/output/bin/todoprofile /bin/todoprofile

ENTRYPOINT ["/bin/todoprofile"]