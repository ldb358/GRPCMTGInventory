package handlers

import (
	"log"
	"golang.org/x/net/context"
	"github.com/ldb358/GRPCMTGInventory/api"
)


//AddMTGCard Add a magic card to the current inventory 
func (s *Server) AddMTGCard(ctx context.Context, newCard *api.AddMTGCardParams) (*api.MTGCardResponse, error) {
	log.Printf("Received new card %s", newCard.Name)
	card := &api.MTGCard{
		Name: newCard.Name,
		Qty: 1, 
	}
	s.Inventory = append(s.Inventory, card)
	return &MTGCardResponse{Cards: []*MTGCard{card}}, nil
}

//GetMTGCard Return all magic cards
func (s *Server) GetMTGCard(ctx context.Context, _ *GetMTGCardParams) (*MTGCardResponse, error) {
	log.Printf("get cards")
	return &MTGCardResponse{Cards: s.Inventory}, nil
}

//DeleteMTGCard Delete a magic card from the inventory
func (s *Server) DeleteMTGCard(ctx context.Context, card *DeleteMTGCardParams) (*DeleteResponse, error) {
	var updatedInventory []*MTGCard
	found := false
	for index, currentCard := range s.Inventory {
		if currentCard.Name == card.Name {
			updatedInventory = append(
				s.Inventory[:index], s.Inventory[index+1:]...)
			found = true
			break
		}
	}
	if found {
		s.Inventory = updatedInventory
		return &DeleteResponse{Message: "success"}, nil
	}
	return &DeleteResponse{Message: "not found"}, nil
}