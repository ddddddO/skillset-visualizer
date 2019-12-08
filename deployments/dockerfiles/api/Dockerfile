FROM golang:1.13-alpine

WORKDIR /api

COPY api ./

RUN ["go", "build"]

EXPOSE 8081
ENTRYPOINT ["./api"]
