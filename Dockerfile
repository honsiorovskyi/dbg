# build stage
FROM golang:alpine AS build-env
ADD . /go/dbg
RUN cd /go/dbg && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' .

# final stage
FROM scratch
ENV LISTEN :8080
COPY --from=build-env /go/dbg/dbg /bin/
ENTRYPOINT ["/bin/dbg"]
