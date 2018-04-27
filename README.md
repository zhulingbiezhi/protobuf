# protobuf
# protobuf练习，增加inject tag， like gorm
# go get github.com/favadi/protoc-go-inject-tag
#  增加注释  // @inject_tag: gorm:"column:"phone""，这样会将注释生成到*.pd.go里面
#  protoc --go_out=. *.proto
#  在生成*.pd.go中可看到你写的注释
#  protoc-go-inject-tag -input *.pb.go可根据注释生成对应的tag
