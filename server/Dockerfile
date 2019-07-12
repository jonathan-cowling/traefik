FROM golang:1.12.6-alpine as build

ENV GOPATH=/build/
ENV GOBIN=/build/bin/

COPY /src/main/go /build/src/main
RUN cd /build/src/main
RUN go install main

FROM alpine:3.10.0 as run

RUN addgroup -S app
RUN adduser -S -G app -h /home/app -s /bin/sh app
USER app
WORKDIR /home/app
EXPOSE 8080

ARG ARG_SUCCESS_THRESHOLD=0.5

COPY --from=build --chown=app:app /build/bin/main /home/app/main
COPY --chown=app:app /src/main/resources/ /home/app/

ENV SUCCESS_THRESHOLD=${ARG_SUCCESS_THRESHOLD}
ENTRYPOINT [ "/home/app/main" ]