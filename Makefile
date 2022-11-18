gen: 
	protoc --go_out=proto/gen proto/$(gen).proto 

clean:
	rm proto/gen/**/*.go 

run:
	go run server/main.go 