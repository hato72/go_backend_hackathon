FROM golang:1.22-alpine

#WORKDIR /app/

# ローカルの場合

# COPY backend/go.sum . 
# COPY backend/go.mod .

# RUN apk upgrade --update && apk --no-cache add git gcc musl-dev

# COPY ./backend .

# CMD ["go", "run", "./src"]


#本番環境の場合

WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod download \
&& go build -o main /src/main.go
CMD [ "/app/main" ]

# COPY go.mod ./

# RUN go mod download

# COPY . .

# #RUN go build -o main .

# EXPOSE 8080

# CMD go run \src\main.go

