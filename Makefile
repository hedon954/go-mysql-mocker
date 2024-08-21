main:
	go run main.go

genpb:
	protoc --go_out=. ./protos/*