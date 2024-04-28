FROM golang:1.22-alpine

WORKDIR /app/

COPY /go.sum ./
COPY /go.mod ./

#RUN apk upgrade --update && apk --no-cache add git gcc musl-dev

RUN go mod download

#COPY backend  ./
COPY . .

EXPOSE 8080

#CMD ["go", "run", "./src"]

CMD ["./src"]





