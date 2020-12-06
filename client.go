package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/thaidzai285/dzai-mp3-protobuf/pkg/pb"
)

func main() {
	conn, _ := grpc.Dial(":3001", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	userClient := pb.NewSongsServiceClient(conn)

	listUsersRequest := pb.SongRequestParams{
		Genre:  "usuk",
	}

	res, err := userClient.GetSongs(context.Background(), &listUsersRequest)
	if err != nil {
		panic(err)
	}
	log.Println(res)
}
