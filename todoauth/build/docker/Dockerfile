FROM golang:1.23.1-alpine3.20 AS init

RUN apk --no-cache add curl ca-certificates jq tzdata libssl3 make bash

WORKDIR /go/src/github.com/nvdovnar/todo-api

COPY . .

WORKDIR /go/src/github.com/nvdovnar/todo-api/todoauth

FROM init as build

RUN make build

FROM alpine:3.19 AS final

COPY --from=build /go/src/github.com/nvdovnar/todo-api/todoauth/output/bin/todoauth /bin/todoauth

ENTRYPOINT ["/bin/todoauth"] 