FROM golang:1.18.3-alpine
WORKDIR /app
COPY ./src/app /app

RUN apk update \
  && apk add --no-cache git \
  && go get \
  && go mod tidy

CMD ["go", "run", "main.go"]

