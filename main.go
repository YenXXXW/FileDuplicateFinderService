package main

func main() {
	grpcServer := NewGrpcServer(":8000")
	grpcServer.Run()
}
