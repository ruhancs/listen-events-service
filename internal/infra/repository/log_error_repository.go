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

type LogErrorRepository struct {
	Client *elasticsearch.Client
}

func NewLogErrorRepository(client *elasticsearch.Client) *LogErrorRepository{
	return &LogErrorRepository{
		Client: client,
	}
}

//_mapping = schema
//_count = contagem de items

func (repo *LogErrorRepository) Register(ctx context.Context, LogError *entity.LogError) (string, error) {

	//cria a map automaticamente, melhor pratica é criar map manualmente no elasticsearch, para evitar erros
	_, err := repo.Client.Index("log_erros", esutil.NewJSONReader(LogError), repo.Client.Index.WithDocumentID(LogError.ID))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return "LogError registered", nil
}

func (repo *LogErrorRepository) BulkRegister(ctx context.Context, LogErrors []*entity.LogError) (string, error) {

	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      "log_erros",
		Client:     repo.Client,
		NumWorkers: 5, //numero de goroutines, processamento em paralelo
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer bulkIndexer.Close(ctx)

	for _, LogError := range LogErrors {
		body,err := json.Marshal(LogError)
		if err != nil {
			return "",err
		}
		err = bulkIndexer.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: LogError.ID,
			Body:       bytes.NewReader(body),
		},
		)
		if err != nil {
			log.Println(err)
			return "", err
		}
	}
	return "all LogErrors registered", nil
}

func (repo *LogErrorRepository) GetByID(ctx context.Context,id string) (dto.GetLogErrorByIDElaticOutputDto, error) {
	resp,err := repo.Client.Get("log_erros",id)
	if err != nil {
		log.Println(err)
		return dto.GetLogErrorByIDElaticOutputDto{},err
	}
	defer resp.Body.Close()
	
	var respOutput dto.GetLogErrorByIDElaticOutputDto
	err = json.NewDecoder(resp.Body).Decode(&respOutput)
	if err != nil {
		log.Println(err)
		return dto.GetLogErrorByIDElaticOutputDto{},err
	}
	
	return respOutput, nil
}

func (repo *LogErrorRepository) SearchByServiceAndDate(ctx context.Context,input dto.SearchWithServiceAndDateInputDto) (dto.SearchWithServiceAndDateOutputDto, error) {
	var searchBuffer bytes.Buffer
	search := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]any{
						"service.en": input.Service,
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
		return dto.SearchWithServiceAndDateOutputDto{},err
	}
	resp,err := repo.Client.Search(
		repo.Client.Search.WithContext(ctx),
		repo.Client.Search.WithIndex("log_erros"),
		repo.Client.Search.WithBody(&searchBuffer),
		repo.Client.Search.WithTrackTotalHits(true),
		repo.Client.Search.WithPretty(),
	)
	defer resp.Body.Close()

	var respOutput dto.SearchWithServiceAndDateOutputDto
	err = json.NewDecoder(resp.Body).Decode(&respOutput)
	if err != nil {
		log.Println(err)
		return dto.SearchWithServiceAndDateOutputDto{},err
	}

	return respOutput,nil
}
