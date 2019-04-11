FROM golang:alpine as builder

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN apk --no-cache add gcc g++ make ca-certificates

COPY . /home/gistdirect

RUN set -x \
    && cd /home/gistdirect/main \
    && go build \
    && mv main /usr/bin/gistdirect \
    && rm -rf /go \
    && echo "Build complete."

FROM alpine:latest

RUN	apk add --no-cache \
    ca-certificates

COPY --from=builder /usr/bin/gistdirect /usr/bin/gistdirect

EXPOSE 8080
ENTRYPOINT [ "gistdirect" ]
CMD [ "" ]