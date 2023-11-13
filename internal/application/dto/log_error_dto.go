package dto

import "github.com/ruhancs/listen-events/internal/domain/entity"

type RegisterLogErrortInputDto struct {
	Message    string `json:"message"`
	Service    string `json:"service"`
	Day        int    `json:"day"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	StatusCode int    `json:"status_code"`
	UserInfo   string `json:"user_info"`
}

type GetLogErrorByIDElaticOutputDto struct {
	Index   string           `json:"_index"`
	ID      string           `json:"_id"`
	Version int              `json:"_version"`
	Source  *entity.LogError `json:"_source"`
}

type SearchWithServiceAndDateInputDto struct {
	Service string `json:"service"`
	Day     int    `json:"day"`
	Month   int    `json:"month"`
	Year    int    `json:"year"`
}

type SearchWithServiceAndDateOutputDto struct {
	Hits struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		Hits []*struct {
			Source *entity.LogError `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
