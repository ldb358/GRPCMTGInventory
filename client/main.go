package main
import (
  "log"
  "github.com/ldb358/GRPCMTGInventory/api"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

func MakeCard(c api.MTGCardServiceClient, name string) {
  response, err := c.AddMTGCard(context.Background(), &api.AddMTGCardParams{Name: name, Qty: 1})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  for _, card := range response.Cards {  
  	log.Printf("Card Created: %s", card.Name)
  }
}

func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial(":7777", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  c := api.NewMTGCardServiceClient(conn)
  MakeCard(c, "test")
  MakeCard(c, "hello")
  MakeCard(c, "world")
  response, err := c.GetMTGCard(context.Background(), &api.GetMTGCardParams{})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  for _, card := range response.Cards {  
  	log.Printf("Card Fetched: %s", card.Name)
  }
}