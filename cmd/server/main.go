package main

import (
	"github.com/Suhach/user-servicee/internal/database"
	"github.com/Suhach/user-servicee/internal/transport/grpc"
	"github.com/Suhach/user-servicee/internal/user"
	"log"
)

func main() {
	database.InitDB()

	repo := user.NewUserREPO(database.DB)

	userService := user.NewUserService(repo)

	if err := grpc.RunGRPC(userService); err != nil {
		log.Fatalf("failed to run grpc server: %v", err)
	}
}
