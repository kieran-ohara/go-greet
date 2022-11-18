package server

import (
	"fmt"
	"net"
	"net/http"
  "google.golang.org/grpc"
)

func Serve() {
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Println( "there was an error" )
  }
}

func RpcServe() {
  lis, err := net.Listen("tcp", fmt.Sprintf("localhost:1234"))
  if err != nil {
    fmt.Sprintln("there wasa grpc error")
  }
  var opts []grpc.ServerOption
  grpcServer := grpc.NewServer(opts...)
  grpcServer.Serve(lis)
}
