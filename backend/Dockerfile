FROM golang:1.22-alpine

WORKDIR /app/

COPY backend/go.sum . 
COPY backend/go.mod .

RUN apk upgrade --update && apk --no-cache add git gcc musl-dev

COPY ./backend .

CMD ["go", "run", "./src"]




