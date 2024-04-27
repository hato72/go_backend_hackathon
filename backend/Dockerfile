FROM golang:1.22-alpine

WORKDIR /app/

COPY go.sum ./
COPY go.mod ./

RUN apk upgrade --update && apk --no-cache add git gcc musl-dev

COPY . .

CMD ["go", "run", "./backend/src"]





