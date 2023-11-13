package dto

import "github.com/ruhancs/listen-events/internal/domain/entity"

type RegisterEventInputDto struct {
	Type    string `json:"type"`
	Service string `json:"service"`
	Day     int    `json:"day"`
	Month   int    `json:"month"`
	Year    int    `json:"year"`
	Status  string `json:"status"`
}

type GetEventByIDElaticOutputDto struct {
	Index   string        `json:"_index"`
	ID      string        `json:"_id"`
	Version int           `json:"_version"`
	Source  *entity.Event `json:"_source"`
}

type SearchWithTypeServiceDateStatusInputDto struct {
	Service string `json:"service"`
	Type    string `json:"type"`
	Day     int    `json:"day"`
	Month   int    `json:"month"`
	Year    int    `json:"year"`
	Status  string `json:"status"`
}

type SearchWithTypeServiceDateStatusOutputDto struct {
	Hits struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		Hits []*struct {
			Source *entity.Event `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
