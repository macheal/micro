
all:
	protoc --proto_path=. --micro_out=. --go_out=. *.proto
	sed -i 's|github.com/micro/go-micro|gitee.com/smartsteps/go-micro|g' *.go