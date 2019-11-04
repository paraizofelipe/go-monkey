package api

import (
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Service struct {
	Host           string      `json:"host,omitempty"`
	CreatedAt      int64       `json:"created_at,omitempty"`
	ConnectTimeout int64       `json:"connect_timeout,omitempty"`
	Id             string      `json:"id,omitempty"`
	Protocol       string      `json:"protocol,omitempty"`
	Name           string      `json:"name,omitempty"`
	ReadTimeout    int64       `json:"read_timeout,omitempty"`
	Port           int64       `json:"port,omitempty"`
	Path           interface{} `json:"path,omitempty"`
	UpdatedAt      int64       `json:"updated_at,omitempty"`
	Retries        int64       `json:"retries,omitempty"`
	WriteTimeout   int64       `json:"write_timeout,omitempty"`
}

type RespService struct {
	Next interface{} `json:"next"`
	Data []Service   `json:"data"`
}

func (s *Service) GetValue(key string) interface{} {
	r := reflect.ValueOf(s)
	k := strings.Title(strings.ToLower(key))
	f := reflect.Indirect(r).FieldByName(k)
	return f
}

func (a *Api) CreateServices(service Service) error {
	if err := a.CreateEntity(service, "services"); err != nil {
		return err
	}

	return nil
}

func (a *Api) ListServices() (error, []Service) {
	var err error

	err, svc := a.ListEntity("services")
	if err != nil {
		return err, nil
	}

	var services []Service

	err = mapstructure.Decode(svc, &services)
	if err != nil {
		return err, nil
	}

	return nil, services
}
