## Build
FROM golang:1.24-bookworm AS build

WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

COPY . ./

ENV CGO_ENABLED=0
ENV GOFLAGS='-buildvcs=false -trimpath -mod=vendor'
RUN go build -o /serve ./cmd/serve.go

## Deploy
FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /serve /serve
COPY --from=build /app/kodata/windrose_base.svg.tmpl /kodata/windrose_base.svg.tmpl
COPY --from=build /app/kodata/windrose_arrow.svg.tmpl /kodata/windrose_arrow.svg.tmpl

EXPOSE 8080

USER nonroot:nonroot
# IMAGE_ID=$(docker build -q -t foo . 2>/dev/null | awk '/Successfully built/{print $NF}')
ENTRYPOINT ["/serve"]