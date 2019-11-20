package api

import (
	"encoding/json"
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

func (s *Service) GetValue(key string) interface{} {
	r := reflect.ValueOf(s)
	k := strings.Title(strings.ToLower(key))
	f := reflect.Indirect(r).FieldByName(k)
	return f
}

func (s *Service) ToMap() (error, map[string]interface{}) {
	b, err := json.Marshal(s)
	if err != nil {
		return err, nil
	}

	var ms map[string]interface{}
	if err := json.Unmarshal(b, &ms); err != nil {
		return err, nil
	}

	return nil, ms
}

func (a *Api) CreateServices(service Service) error {
	if err := a.CreateEntity(&service, "services"); err != nil {
		return err
	}

	return nil
}

func (a *Api) Service(id string) (error, Service) {
	var err error
	var service Service

	err, svc := a.GetEntity("services", id)
	if err != nil {
		return err, service
	}

	err = mapstructure.Decode(svc, &service)
	if err != nil {
		return err, service
	}

	return nil, service
}

func (a *Api) Services() (error, []Service) {
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
