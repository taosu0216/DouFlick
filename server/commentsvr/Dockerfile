FROM golang:alpine as build
LABEL authors="taosu"
WORKDIR /app

ARG GOPROXY="https://goproxy.cn,direct"
COPY . .
RUN cd server/commentsvr && go mod init commentsvr && go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go
#RUN go mod init commentsvr
#RUN go mod tidy
RUN ls server/commentsvr
RUN ls
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch as deploy
WORKDIR /app/run
COPY --from=build /app/server/commentsvr/main ./main
#COPY --from=build /app/config/config.yaml ./config/config.yaml
COPY --from=build /app/server/commentsvr/main.go ./main.go

CMD ["./main"]
