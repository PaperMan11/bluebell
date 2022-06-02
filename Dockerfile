FROM golang:alpine AS builder 

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GOPROXY="https://goproxy.cn,direct"

# 移动到工作目录
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum 文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bluebell
RUN go build -o bluebell .

# 分阶段构建
FROM scratch

# 配置文件
COPY ./conf /conf

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bluebell /

# 声明端口
EXPOSE 8080

# 需要运行的命令
ENTRYPOINT ["/bluebell"]
