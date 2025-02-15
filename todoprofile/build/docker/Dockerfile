FROM golang:1.23.1-alpine3.20 AS init

RUN apk --no-cache add curl ca-certificates jq tzdata libssl3 make bash

WORKDIR /go/src/github.com/nvdovnar/todo-api

COPY todolib/go.mod todolib/go.sum ./todolib/
COPY todo/go.mod todo/go.sum ./todo/
COPY todoauth/go.mod todoauth/go.sum ./todoauth/
COPY todoprofile/go.mod todoprofile/go.sum ./todoprofile/
COPY go.work ./

RUN go mod download

COPY . .

WORKDIR /go/src/github.com/nvdovnar/todo-api/todoprofile

FROM init as build

RUN make build

FROM scratch AS final

COPY --from=build /go/src/github.com/nvdovnar/todo-api/todoprofile/output/bin/todoprofile /bin/todoprofile

ENTRYPOINT ["/bin/todoprofile"]    