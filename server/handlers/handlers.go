package handlers

import "github.com/ldb358/GRPCMTGInventory/models"

//Server gRPC server
type Server struct {
	DB *models.DB
}