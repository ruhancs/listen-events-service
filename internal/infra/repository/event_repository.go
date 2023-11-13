package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/ruhancs/listen-events/internal/application/dto"
	"github.com/ruhancs/listen-events/internal/domain/entity"
)

type EventRepository struct {
	Client *elasticsearch.Client
}

func NewEventRepository(client *elasticsearch.Client) *EventRepository {
	return &EventRepository{
		Client: client,
	}
}

//_mapping = schema
//_count = contagem de items

func (repo *EventRepository) Register(ctx context.Context, event *entity.Event) (string, error) {

	//cria a map automaticamente, melhor pratica é criar map manualmente no elasticsearch, para evitar erros
	_, err := repo.Client.Index("events", esutil.NewJSONReader(event), repo.Client.Index.WithDocumentID(event.ID))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return "event registered", nil
}

func (repo *EventRepository) BulkRegister(ctx context.Context, events []*entity.Event) (string, error) {

	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      "events",
		Client:     repo.Client,
		NumWorkers: 5, //numero de goroutines, processamento em paralelo
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer bulkIndexer.Close(ctx)

	for _, event := range events {
		body,err := json.Marshal(event)
		if err != nil {
			return "",err
		}
		err = bulkIndexer.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: event.ID,
			Body:       bytes.NewReader(body),
		},
		)
		if err != nil {
			log.Println(err)
			return "", err
		}
	}
	return "all events registered", nil
}

func (repo *EventRepository) GetByID(ctx context.Context,id string) (dto.GetEventByIDElaticOutputDto, error) {
	resp,err := repo.Client.Get("events",id)
	if err != nil {
		log.Println(err)
		return dto.GetEventByIDElaticOutputDto{},err
	}
	defer resp.Body.Close()
	
	var respOutput dto.GetEventByIDElaticOutputDto
	err = json.NewDecoder(resp.Body).Decode(&respOutput)
	if err != nil {
		log.Println(err)
		return dto.GetEventByIDElaticOutputDto{},err
	}
	
	return respOutput, nil
}

func (repo *EventRepository) SearchByTypeServiceDateAndStatus(ctx context.Context,input dto.SearchWithTypeServiceDateStatusInputDto) (dto.SearchWithTypeServiceDateStatusOutputDto, error) {
	var searchBuffer bytes.Buffer
	search := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]any{
						"service.en": input.Service,
						"type.en": input.Type,
						"status.en": input.Status,
						"day.en": input.Day,
						"month.en": input.Month,
						"year.en": input.Year,
					},
				},
			},
		},
	}
	err := json.NewEncoder(&searchBuffer).Encode(search)
	if err != nil {
		log.Println(err)
		return dto.SearchWithTypeServiceDateStatusOutputDto{},err
	}
	resp,err := repo.Client.Search(
		repo.Client.Search.WithContext(ctx),
		repo.Client.Search.WithIndex("events"),
		repo.Client.Search.WithBody(&searchBuffer),
		repo.Client.Search.WithTrackTotalHits(true),
		repo.Client.Search.WithPretty(),
	)
	defer resp.Body.Close()

	var respOutput dto.SearchWithTypeServiceDateStatusOutputDto
	err = json.NewDecoder(resp.Body).Decode(&respOutput)
	if err != nil {
		log.Println(err)
		return dto.SearchWithTypeServiceDateStatusOutputDto{},err
	}

	return respOutput,nil
}

func (repo *EventRepository) List(ctx context.Context) (string, error) {
	
	
	return "all events registered", nil
}
