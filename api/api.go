package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Api struct {
	BaseUrl   string
	EndPoints map[string]string
}

var paths = map[string]string{
	"services": "/services",
	"routes":   "/routes",
}

type RespEntity struct {
	Next interface{} `json:"next"`
	Data []Entity    `json:"data"`
}

func New(baseUrl string) *Api {
	return &Api{
		BaseUrl:   baseUrl,
		EndPoints: paths,
	}
}

func (a *Api) makeRequests(method string, url string, body io.Reader) (error, map[string]interface{}) {
	client := http.Client{}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return err, nil
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return errors.New(resp.Status), nil
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err, nil
	}
	return nil, result
}

func (a *Api) CreateEntity(entity Entity, entityName string) error {
	u, err := url.Parse(a.BaseUrl)
	if err != nil {
		log.Fatal("Base URL invalid")
	}
	u.Path = path.Join(u.Path, a.EndPoints[entityName])

	requestBody, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	if err, _ = a.makeRequests("POST", u.String(), bytes.NewBuffer(requestBody)); err != nil {
		return err
	}

	return nil
}

func (a *Api) GetEntity(entityName string, id string) (error, interface{}) {
	u, err := url.Parse(a.BaseUrl)
	if err != nil {
		log.Fatal("Base URL invalid")
	}
	u.Path = path.Join(u.Path, a.EndPoints[entityName], id)

	err, result := a.makeRequests("GET", u.String(), nil)
	if err != nil {
		return err, nil
	}

	return nil, result
}

func (a *Api) ListEntity(entityName string) (error, []interface{}) {
	u, err := url.Parse(a.BaseUrl)
	if err != nil {
		log.Fatal("Base URL invalid")
	}
	u.Path = path.Join(u.Path, a.EndPoints[entityName])

	err, result := a.makeRequests("GET", u.String(), nil)
	if err != nil {
		return err, nil
	}

	return nil, result["data"].([]interface{})
}
