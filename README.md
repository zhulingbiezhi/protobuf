# protobuf 练习

## 生成pb
```
protoc -I/usr/local/include \
-I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  \
--go_out=plugins=grpc:.  \
grpc/server.proto
```
- 第一个path是在安装protoc的时候自带一个include文件夹，cp过去即可
- 第二个path是本地
- 第三个是google的第三方api

## 生成tag
-  go get github.com/favadi/protoc-go-inject-tag
-  增加注释  // @inject_tag: gorm:"column:"phone""，这样会将注释生成到*.pd.go里面
-  在生成*.pd.go中可看到你写的注释
-  protoc-go-inject-tag -input *.pb.go可根据注释生成对应的tag
-  此处直接对go文件进行操作

## 生成api gateway
- 增加--grpc-gateway_out=
```
protoc -I/usr/local/include \
-I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=. \
server.proto
```