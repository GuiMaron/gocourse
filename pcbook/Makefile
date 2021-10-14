clean:
	clear && rm -rf ./pb/*

gen:
	make clean && clear && protoc --proto_path=./proto --go_out=plugins=grpc:./pb ./proto/*.proto

init:
	clear && go mod init github.com/GuiMaron/gocourse

run:
	clear && go run *.go

tidy:
	clear && go mod tidy
