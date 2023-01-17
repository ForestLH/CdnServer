FROM golang:1.18
LABEL MAINTAINER = "lee 1838249551@qq.com"
ENV MYPATH /opt/CdnServer
WORKDIR $MYPATH
#COPY config/nginx.conf /etc/nginx/nginx.conf
COPY . .
EXPOSE 8000

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
		GOPROXY=https://goproxy.cn
VOLUME ["/opt/CdnServer/res", "/opt/CdnServer/log"]
RUN go build -v -o $MYPATH ./...

ENTRYPOINT ["/opt/CdnServer/CdnServer"]
