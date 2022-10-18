package main

import "upvote/grpc/proto/pb"

type Server struct {
	pb.CryptoServiceServer
}