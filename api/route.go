package api

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Route struct {
	Id                      string              `json:"id,omitempty"`
	CreatedAt               int64               `json:"create_at,omitempty"`
	UpdateAt                int64               `json:"update_at,omitempty"`
	Name                    string              `json:"name,omitempty"`
	Protocols               []string            `json:"protocols,omitempty"`
	Methods                 []string            `json:"methods,omitempty"`
	Hosts                   []string            `json:"hosts,omitempty"`
	Paths                   []string            `json:"paths,omitempty"`
	Headers                 map[string][]string `json:"headers,omitempty"`
	HttpsRedirectStatusCode int64               `json:"https_redirect_status_code,omitempty"`
	RegexPriority           int64               `json:"regex_priority,omitempty"`
	StripPath               bool                `json:"strip_path,omitempty"`
	PreserveHost            bool                `json:"preserve_host,omitempty"`
	Tags                    string              `json:"tags,omitempty"`
	Service                 Service             `json:"service,omitempty"`
}

func (r *Route) GetValue(key string) interface{} {
	rt := reflect.ValueOf(r)
	k := strings.Title(strings.ToLower(key))
	f := reflect.Indirect(rt).FieldByName(k)
	return f
}

func (r *Route) ToMap() (error, map[string]interface{}) {
	b, err := json.Marshal(r)
	if err != nil {
		return err, nil
	}

	var ms map[string]interface{}
	if err := json.Unmarshal(b, &ms); err != nil {
		return err, nil
	}

	return nil, ms
}

func (a *Api) CreateRoute(route Route) error {
	if err := a.CreateEntity(&route, "routes"); err != nil {
		return err
	}

	return nil
}

func (a *Api) Route(id string) (error, Route) {
	var err error
	var route Route

	err, svc := a.GetEntity("routes", id)
	if err != nil {
		return err, route
	}

	err = mapstructure.Decode(svc, &route)
	if err != nil {
		return err, route
	}

	err, route.Service = a.Service(route.Service.Id)
	if err != nil {
		return err, route
	}

	return nil, route
}

func (a *Api) Routes() (error, []Route) {
	var err error

	err, rts := a.ListEntity("routes")
	if err != nil {
		return err, nil
	}

	var routes []Route
	err = mapstructure.Decode(rts, &routes)
	if err != nil {
		return err, nil
	}

	for index := 0; index < len(routes); index++ {
		err, routes[index].Service = a.Service(routes[index].Service.Id)
		if err != nil {
			return err, nil
		}
	}

	return nil, routes
}
