clean:
	clear && rm -rf ./pb/*

gen:
	make clean && clear && protoc --proto_path=proto --go_out=plugins=grpc:pb proto/*.proto

init:
	clear && rm -f go.mod && rm -f go.sum && go mod init github.com/GuiMaron/gocourse/pcbook

run:
	clear && go run *.go

test:
	clear && go test ./serializer

tidy:
	clear && go mod tidy
