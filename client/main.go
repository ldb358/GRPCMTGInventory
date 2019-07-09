package main
import (
  "log"
  "github.com/ldb358/GRPCMTGInventory/api"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials"
)

// Authentication holds the login/password
type Authentication struct {
	Token string
  }
  // GetRequestMetadata gets the current request metadata
  func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
	  "token": a.Token,
	}, nil
  }
  // RequireTransportSecurity indicates whether the credentials requires transport security
  func (a *Authentication) RequireTransportSecurity() bool {
	return true
  }

func makeCard(c api.MTGCardServiceClient, name string) {
	response, err := c.AddMTGCard(context.Background(), &api.AddMTGCardParams{Name: name, Qty: 1})
	if err != nil {
		log.Fatalf("Error when calling AddCard: %s", err)
	}
	for _, card := range response.Cards {  
		log.Printf("Card Created: %s", card.Name)
	}
}

func main() {
	// Create the client TLS credentials
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	// Setup the token auth THIS IS A TEMP PLACE HOLDER FOR JWT LATER
	auth := Authentication{
		Token: "098f6bcd4621d373cade4e832627b4f6",
  	}
	
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(
		"localhost:7777",
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewMTGCardServiceClient(conn)
	makeCard(c, "test")
	makeCard(c, "hello")
	makeCard(c, "world")
	response, err := c.GetMTGCard(context.Background(), &api.GetMTGCardParams{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	for _, card := range response.Cards {  
		log.Printf("Card Fetched: %s", card.Name)
	}
}