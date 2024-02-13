FROM golang:alpine as build
LABEL authors="taosu"
WORKDIR /app

ARG GOPROXY="https://goproxy.cn,direct"
COPY . .
RUN cd server/relationsvr && go mod init relationsvr && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM scratch as deploy
WORKDIR /app/run
COPY --from=build /app/server/relationsvr/main ./main
COPY --from=build /app/server/relationsvr/main.go ./main.go

CMD ["./main"]
