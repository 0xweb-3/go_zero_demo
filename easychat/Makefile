user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-test # 对应easychat/deploy/mk/user-rpc.mk中最后指令设置

release-test: user-rpc-dev

install-server:
	cd ./deploy/script && chmod +x release-test.sh && release-test.sh