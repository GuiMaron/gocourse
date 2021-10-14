clean:
	clear && rm -rf ./pb/*

gen:
	make clean && clear && protoc --proto_path=./proto --go_out=plugins=grpc:./pb ./proto/*.proto && mv ./pb/github.com/GuiMaron/gocourse/*.go ./pb && rm -rf ./pb/github.com

init:
	clear && go mod init github.com/GuiMaron/gocourse

run:
	clear && go run *.go

tidy:
	clear && go mod tidy
