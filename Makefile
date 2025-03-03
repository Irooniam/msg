build_protos:
	rm -f protos/connections.pb.go
	protoc --go_out=. --go_opt=paths=source_relative  protos/connections.proto
