VERSION=latest

SERVER_NAME=social
SERVER_TYPE=api

# 测试环境配置
# docker的镜像发布地址
DOCKER_REPO_TEST=crpi-8mkzdazjkdk1gs4d.cn-hangzhou.personal.cr.aliyuncs.com/xin_go_zero/${SERVER_NAME}-${SERVER_TYPE}-dev
# 测试版本
VERSION_TEST=$(VERSION)
# 编译的程序名称
APP_NAME_TEST=easy-im-${SERVER_NAME}-${SERVER_TYPE}-test

# 测试下的编译文件
DOCKER_FILE_TEST=./deploy/dockerfile/Dockerfile_${SERVER_NAME}_${SERVER_TYPE}_dev

# 测试环境的编译发布
build-test:

	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${SERVER_NAME}-${SERVER_TYPE} ./apps/${SERVER_NAME}/${SERVER_TYPE}/${SERVER_NAME}.go
	docker build . -f ${DOCKER_FILE_TEST} --no-cache -t ${APP_NAME_TEST}

# 镜像的测试标签
tag-test:

	@echo 'create tag ${VERSION_TEST}'
	docker tag ${APP_NAME_TEST} ${DOCKER_REPO_TEST}:${VERSION_TEST}

publish-test:

	@echo 'publish ${VERSION_TEST} to ${DOCKER_REPO_TEST}'
	docker push $(DOCKER_REPO_TEST):${VERSION_TEST}

release-test: build-test tag-test publish-test