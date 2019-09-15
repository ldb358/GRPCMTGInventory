package handlers

import (
	"log"
	"golang.org/x/net/context"
	"github.com/ldb358/GRPCMTGInventory/api"
)


//AddInventory Create a new inventory
func (s *Server) AddInventory(ctx context.Context, newInventory *api.AddInventoryParams) (*api.InventoryResponse, error) {
	log.Printf("Received new inventory %s", newInventory.Name)
	inventory, err := s.DB.CreateInventory(newInventory.Name)
	if err != nil {
		return nil, err
	}
	return &api.InventoryResponse{Id: inventory.Id, Name: inventory.Name}, nil
}