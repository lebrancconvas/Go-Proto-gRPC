gen: 
	protoc --go_out=proto/gen proto/$(gen).proto 

clean:
	rm pb/*.go 

run:
	go run server/main.go 