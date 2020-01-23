FROM golang:1.13-alpine as dep_builder
ENV APP_NAME app
WORKDIR /opt/${APP_NAME}
COPY . .
RUN go mod download

FROM dep_builder as app_builder
ENV APP_NAME app
WORKDIR /opt/${APP_NAME}
COPY --from=dep_builder /opt/app .
RUN CGO_ENABLED=0 go build -o ./bin/app ./cmd

FROM scratch
WORKDIR /opt/app
COPY --from=app_builder /opt/app/bin/app ./bin/
CMD ["./bin/app"]
