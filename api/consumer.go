package api

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Consumer struct {
	Id        string   `json:"id,omitempty"`
	CreatedAt int64    `json:"create_at,omitempty"`
	Username  string   `json:"username,omitempty"`
	CustomId  string   `json:"custom_id,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

func (c *Consumer) GetValue(key string) interface{} {
	r := reflect.ValueOf(c)
	k := strings.Title(strings.ToLower(key))
	f := reflect.Indirect(r).FieldByName(k)
	return f
}

func (c *Consumer) ToMap() (error, map[string]interface{}) {
	b, err := json.Marshal(c)
	if err != nil {
		return err, nil
	}

	var ms map[string]interface{}
	if err := json.Unmarshal(b, &ms); err != nil {
		return err, nil
	}

	return nil, ms
}

func (a *Api) CreateConsumer(consumer Consumer) error {
	if err := a.CreateEntity(&consumer, "consumers"); err != nil {
		return err
	}

	return nil
}

func (a *Api) GetConsumer(id string) (error, Consumer) {
	var err error
	var consumer Consumer

	err, cs := a.GetEntity("consumers", id)
	if err != nil {
		return err, consumer
	}

	err = mapstructure.Decode(cs, &consumer)
	if err != nil {
		return err, consumer
	}

	return nil, consumer
}

func (a *Api) ListConsumer() (error, []Consumer) {
	var err error

	err, svc := a.ListEntity("consumers")
	if err != nil {
		return err, nil
	}

	var consumers []Consumer

	err = mapstructure.Decode(svc, &consumers)
	if err != nil {
		return err, nil
	}

	return nil, consumers
}
