# go-protobuf-demo
## 第一步，在本机安装protobuf
### brew install protobuf 

## 第二步，安装protoc-gen-go
### go get -u github.com/golang/protobuf/protoc-gen-go

## 第三步，写proto文件

## 第四步生成go文件,在当前目录下执行如下命令
### protoc --go_out=. *.proto
