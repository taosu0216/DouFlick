FROM golang:alpine as build
LABEL authors="taosu"
WORKDIR /app

ARG GOPROXY="https://goproxy.cn,direct"
COPY . .
RUN cd server/favoritesvr && go mod init favoritesvr && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
RUN ls server/favoritesvr
RUN ls

FROM scratch as deploy
WORKDIR /app/run
COPY --from=build /app/server/favoritesvr/main ./main
COPY --from=build /app/server/favoritesvr/main.go ./main.go

CMD ["./main"]
