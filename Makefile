.PHONY: proto
proto:
	@protoc --proto_path=pb --go_out=:. pb/*.proto;