echo "开始构建"
export GOPATH=$WORKSPACE/..
export PATH=$PATH:$GOROOT/bin

# Print Go version
go version

export GO111MODULE=on
export GOPROXY=https://goproxy.io
export ENV=local

echo "current: ${USER}"
#拷贝配置文件到target下
mkdir -vp goods_srv/target/goods_srv
cp goods_srv/config-pro.yaml goods_srv/target/goods_srv/config-pro.yaml

cd $WORKSPACE
go build -o goods_srv/target/main goods_srv/main.go
echo "构建结束2"