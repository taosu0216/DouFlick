FROM golang:alpine as build
LABEL authors="taosu"
WORKDIR /app

ARG GOPROXY="https://goproxy.cn,direct"
COPY . .
RUN cd server/videosvr && go mod init videosvr && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM scratch as deploy
WORKDIR /app/run
COPY --from=build /app/server/videosvr/main ./main
COPY --from=build /app/server/videosvr/main.go ./main.go

CMD ["./main"]
