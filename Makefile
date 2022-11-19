gen: 
	protoc --go_out=proto/gen --go-grpc_out=proto/gen/services proto/$(gen).proto  

clean:
	rm proto/gen/**/*.go 

run:
	go run server/main.go 