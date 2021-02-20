FROM golang:1.15-alpine as builder

ARG REVISION

RUN mkdir -p /chaos-arcade/

WORKDIR /chaos-arcade

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags "-s -w \
    -X github.com/ryneal/chaos-arcade/pkg/version.REVISION=${REVISION}" \
    -a -o bin/chaos-arcade cmd/chaos-arcade/*

RUN CGO_ENABLED=0 go build -ldflags "-s -w \
    -X github.com/ryneal/chaos-arcade/pkg/version.REVISION=${REVISION}" \
    -a -o bin/podcli cmd/podcli/*

FROM alpine:3.12

ARG BUILD_DATE
ARG VERSION
ARG REVISION

LABEL maintainer="ryneal"

RUN addgroup -S app \
    && adduser -S -G app app \
    && apk --no-cache add \
    ca-certificates curl netcat-openbsd

WORKDIR /home/app

COPY --from=builder /chaos-arcade/bin/chaos-arcade .
COPY --from=builder /chaos-arcade/bin/podcli /usr/local/bin/podcli
COPY ./ui ./ui
RUN chown -R app:app ./

USER app

CMD ["./chaos-arcade"]
