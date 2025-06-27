#!/bin/bash
reso_addr='crpi-8mkzdazjkdk1gs4d.cn-hangzhou.personal.cr.aliyuncs.com/xin_go_zero/user-api-dev'
tag='latest'

pod_ip="192.168.1.247" # 用于设置外部访问服务ip

container_name="easy-chat-user-api-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-chat -v /easy-chat/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 9002:7002 -e POD_IP=${pod_ip} --name=${container_name} -d ${reso_addr}:${tag}