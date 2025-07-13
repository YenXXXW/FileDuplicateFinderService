gen:
	@protoc \
		--proto_path=protobuf "protobuf/fileDuplicateFinder.proto" \
		--go_out=genproto/fileduppb --go_opt=paths=source_relative \
  	--go-grpc_out=genproto/fileduppb --go-grpc_opt=paths=source_relative
