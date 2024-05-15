FROM golang:1.22.0 as build
WORKDIR /goservice
COPY api api
COPY ["cmd/http-server", "cmd/http-server"]
COPY docs docs
COPY config config
COPY domain domain
COPY dto dto
COPY persistence persistence
#COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
RUN echo 'Listing of all cmd/http-server files'
RUN ls -al cmd/http-server
RUN go clean -modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /publish/goservice ./cmd/http-server


FROM alpine:latest as publish
WORKDIR /publish

COPY --from=build /publish/goservice .
COPY --from=build goservice/docs/swagger.* ./docs
COPY --from=build goservice/cmd/http-server/config.yml .
EXPOSE 8080
ENTRYPOINT ["./goservice"]
