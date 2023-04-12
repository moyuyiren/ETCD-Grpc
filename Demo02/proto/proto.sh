echo "生成 rpc server 代码"
protoc --go_out=:. *.proto
if [ $? -eq 0 ];then
  echo "数据模型创建完成"
fi
protoc --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false *.proto
if [ $? -eq 0 ];then
  echo "方法创建完成"
fi