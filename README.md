# protobuf
protobuf练习，增加inject tag， like gorm
1、go get github.com/favadi/protoc-go-inject-tag
2、增加注释  // @inject_tag: gorm:"column:"phone""，这样会将注释生成到*.pd.go里面
3、protoc --go_out=. *.proto
4、在生成*.pd.go中可看到你写的注释
5、protoc-go-inject-tag -input *.pb.go可根据注释生成对应的tag
