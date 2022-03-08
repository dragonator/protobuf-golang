
.PHONY: pb
pb:
	rm -rf internal/pb/*
	protoc -I=pb/ --go_out=. pb/message.proto
