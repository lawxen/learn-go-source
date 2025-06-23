# 基础镜像
FROM alpine:3.18
# 镜像作者
MAINTAINER HYX
# 在容器内创建目录
RUN mkdir -p /home/servers/gin
# 设置容器内工作目录
WORKDIR /home/servers/gin
# 将当前目录gin复制到容器的/home/servers/gin
COPY ./gin /home/servers/gin
# 设置编译文件权限
RUN chmod 777 /home/servers/gin/main