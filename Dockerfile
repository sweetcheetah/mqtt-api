FROM golang:1.16 as builder
COPY . /src
WORKDIR /src

RUN go test ./... -v -cover
RUN CGO_ENABLED=0 GOOS=linux go build -o mqtt-api .

FROM scratch
WORKDIR /

COPY --from=builder /src/env/dev.env /.env
COPY --from=builder /src/swagger-ui/ /swagger-ui
COPY --from=builder /src/mqtt-api /mqtt-api

EXPOSE 8081
ENV PORT=8081

CMD ["/mqtt-api"]
