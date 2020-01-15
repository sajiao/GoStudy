# 第一阶段Go语言编译环境，别名builder
FROM golang as builder

COPY * /build/
WORKDIR /build
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -ldflags '-w -s' -o server

# RUN go build -ldflags '-w -s' -o server.exe

# 运行阶段
FROM builder as final
WORKDIR /app
# 从编译阶段的中拷贝编译结果到当前镜像中 这个 Dockerfile 的玄妙之处就在于 COPY 指令的 --from=0 参数，从前边的阶段中拷贝文件到当前阶段中，多个FROM语句时，0代表第一个阶段。除了使用数字，我们还可以给阶段命名
# 更为强大的是，COPY --from 不但可以从前置阶段中拷贝，还可以直接从一个已经存在的镜像中拷贝。比如，
# COPY --from=quay.io/coreos/etcd:v3.3.9 /usr/local/bin/etcd /usr/local/bin
COPY --from=builder /build /app

# ENTRYPOINT ["/app/server.exe"]
CMD /app/server.exe
# Install beego and the bee dev tool
# RUN go get github.com/astaxie/beego && go get github.com/beego/bee

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
# CMD ["bee", "run"]
# RUN mkdir -p /root/gopath/src/server
# COPY * /root/gopath/src/server/
# WORKDIR /root/gopath/src/server
# RUN ReBuild.bat
# CMD /root/gopath/src/server/MathApp.exe

