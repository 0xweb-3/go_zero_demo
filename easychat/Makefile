user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-test # 对应easychat/deploy/mk/user-rpc.mk中最后指令设置

user-api-dev:
	@make -f deploy/mk/user-api.mk release-test

social-rpc-dev:
	@make -f deploy/mk/social-rpc.mk release-test

social-api-dev:
	@make -f deploy/mk/social-api.mk release-test

release-test: user-rpc-dev user-api-dev social-rpc-dev social-api-dev # 编译打包到镜像仓库

install-server:
	cd ./deploy/script && chmod +x release-test.sh && release-test.sh # 进行部署