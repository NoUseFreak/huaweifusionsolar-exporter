FROM golang:1.22-alpine AS build

WORKDIR /app
COPY . /app
RUN go mod vendor \
    && CGO_ENABLED=0 go build -ldflags='-w -s -extldflags "-static"' -a main.go 

FROM gcr.io/distroless/static

COPY --from=build /app/main /app/main
USER 65532
CMD ["/app/main", "metrics"]
