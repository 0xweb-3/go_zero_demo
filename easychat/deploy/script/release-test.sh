need_start_server_shell=(
  #user rpc
  user-rpc-test.sh
  #user api
  user-api-test.sh

  # social rpc
  social-rpc-test.sh
  # social api
  social-api-test.sh
)

for i in ${need_start_server_shell[*]} ; do
    chmod +x $i
    ./$i
done


docker ps # 打印所有容器

docker exec -it etcd etcdctl get --prefix "" # 打印etcd下所有key