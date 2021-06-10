
all:
	protoc --proto_path=. --micro_out=. --go_out=. *.proto
	sed -i 's|github.com/micro/go-micro|github.com/macheal/go-micro|g' *.go