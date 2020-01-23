FROM golang:1.13-alpine as dep_builder
ENV APP_NAME app
WORKDIR /opt/${APP_NAME}
COPY . .
RUN go mod download

FROM dep_builder as app_builder
ENV APP_NAME app
ENV CGO_ENABLED 0
WORKDIR /opt/${APP_NAME}
COPY --from=dependency-builder /opt/app .
RUN go build -o ./bin/app .

FROM scratch
WORKDIR /opt/app
COPY --from=app-builder /opt/abf-guard/bin/app ./bin/
CMD ["./bin/app"]
