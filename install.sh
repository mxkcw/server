#!/usr/bin/env bash
###
 # @creater: hc_wen kzh130@163.com
 # @since: 2024-05-23 15:57:45
 # @lastTime: 2024-05-24 14:44:36
 # @LastAuthor: hc_wen kzh130@163.com
 # @Path: \wynpay_website\config.sh
 # @message: headers
###
# 打包镜像名称变量
app_name='wynpay_website_statistics';
# 打包的版本号
app_version='latest';
echo '----开始执行sh命令----'
echo '判断是否存在docker运行的容器'
if [ $(docker ps -aq --filter name=${app_name}) ];
then
echo '----暂停docker容器----'
docker stop ${app_name};
echo '----移除docker容器-----'
docker rm -f ${app_name};
echo '----移除docker容器完成----'
echo '----移除images----'
docker rmi ${app_name}:${app_version};
echo '----移除images完成----'
fi
echo '----根据当前目录下Dockerfile制作镜像----'
docker build -t ${app_name}:${app_version} .;
echo '----运行镜像文件----'
docker run --cap-add SYS_TIME -d -p 8001:5001 --name ${app_name} -v /data/${app_name}:/data/webSite_statistics ${app_name};
echo '----完成部署----'