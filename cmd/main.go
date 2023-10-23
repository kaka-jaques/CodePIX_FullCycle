package main

import (
	"github.com/jinzhu/gorm"
	"github.com/kaka-jaques/CodePIX_FullCycle/codepix-go/application/grpc"
	"github.com/kaka-jaques/CodePIX_FullCycle/codepix-go/infrastructure/db"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
