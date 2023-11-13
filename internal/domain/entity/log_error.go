package entity

import (
	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
)

// entidade com informacoes do erro para armazenar.
// Service = micro service que aconteceu o erro, StatusCode = status code do erro,
// Date = data que o erro ocorreu, Message = menssagem do erro, UserInfo = informacoes do usuario ex:sistema operacional
type LogError struct {
	ID         string `json:"id" valid:"required"`
	Message    string `json:"message" valid:"required"`
	Service    string `json:"service" valid:"required"`
	StatusCode int    `json:"status_code" valid:"required"`
	Day        int    `json:"day" valid:"required"`
	Month      int    `json:"month" valid:"required"`
	Year       int    `json:"year" valid:"required"`
	UserInfo   string `json:"user_info" valid:"required"`
}

func NewLogError(message, service, userInfo string, statusCode,day,month,year int) (*LogError, error) {
	log := &LogError{
		ID:         xid.New().String(),
		Message:    message,
		Service:    service,
		StatusCode: statusCode,
		Day: day,
		Month: month,
		Year: year,
		UserInfo:   userInfo,
	}
	err := log.validate()
	if err != nil {
		return nil, err
	}
	return log, nil
}

func (log *LogError) validate() error {
	_, err := govalidator.ValidateStruct(log)
	if err != nil {
		return err
	}
	return nil
}
