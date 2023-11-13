package elastic

import (
	"context"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func ConnectWithElasticSearch(ctx context.Context) *elasticsearch.Client {
	newClient,err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELK_URL"),
		},
	})
	if err != nil {
		panic(err)
	}
	//inserir client no contexto
	// context.WithValue(ctx,"elc",newClient)
	return newClient
}