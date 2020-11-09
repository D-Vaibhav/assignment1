.PHONY = protos

protos: 
	protoc proto/blog.proto --go_out=plugins=grpc:.