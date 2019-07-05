FROM golang:1.12.6-alpine as build

ENV GOPATH=/build/
ENV GOBIN=/build/bin/

COPY /src/main/go /build/src/main
RUN cd /build/src/main
RUN go install main

FROM alpine:3.10.0 as run

RUN useradd -mUs /bin/sh app
RUN addgroup -S appgroup && adduser -S appuser -G appgroup && false
USER app
WORKDIR /home/app
EXPOSE 8080

COPY --from=build --chown=app:app /build/bin/main /home/app/main
COPY --chown=app:app /src/main/resources/* /home/app/

ENTRYPOINT [ "/home/app/main" ]