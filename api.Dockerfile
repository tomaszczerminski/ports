FROM golang:1.15.0-alpine AS build
WORKDIR .
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o cmd/domain.go .

FROM scratch AS bin
COPY --from=build /out/example /
