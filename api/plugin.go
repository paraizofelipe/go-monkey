package api

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Plugin struct {
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	CreatedAt int      `json:"created_at,omitempty"`
	Route     Route    `json:"route,omitempty"`
	Service   Service  `json:"service,omitempty"`
	Consumer  Consumer `json:"consumer,omitempty"`
	Config    Config   `json:"config,omitempty"`
	RunOn     string   `json:"run_on,omitempty"`
	Protocols []string `json:"protocols,omitempty"`
	Enabled   bool     `json:"enabled,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type Config struct {
	Hour   int `json:"hour,omitempty"`
	Minute int `json:"minute,omitempty"`
}

func (r *Plugin) GetValue(key string) interface{} {
	rt := reflect.ValueOf(r)
	k := strings.Title(strings.ToLower(key))
	f := reflect.Indirect(rt).FieldByName(k)
	return f
}

func (r *Plugin) ToMap() (error, map[string]interface{}) {
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

func (a *Api) CreatePlugin(plugin Plugin) error {
	if err := a.CreateEntity(&plugin, "routes"); err != nil {
		return err
	}

	return nil
}

func (a *Api) Plugin(id string) (error, Plugin) {
	var err error
	var plugin Plugin

	err, svc := a.GetEntity("routes", id)
	if err != nil {
		return err, plugin
	}

	err = mapstructure.Decode(svc, &plugin)
	if err != nil {
		return err, plugin
	}

	err, plugin.Service = a.Service(plugin.Service.Id)
	if err != nil {
		return err, plugin
	}

	err, plugin.Route = a.Route(plugin.Route.Id)
	if err != nil {
		return err, plugin
	}

	err, plugin.Consumer = a.Consumer(plugin.Consumer.Id)
	if err != nil {
		return err, plugin
	}

	return nil, plugin
}

func (a *Api) Plugins() (error, []Plugin) {
	var err error

	err, rts := a.ListEntity("plugins")
	if err != nil {
		return err, nil
	}

	var plugins []Plugin
	err = mapstructure.Decode(rts, &plugins)
	if err != nil {
		return err, nil
	}

	for index := 0; index < len(plugins); index++ {
		err, plugins[index].Service = a.Service(plugins[index].Service.Id)
		if err != nil {
			return err, nil
		}

		err, plugins[index].Route = a.Route(plugins[index].Route.Id)
		if err != nil {
			return err, nil
		}

		err, plugins[index].Consumer = a.Consumer(plugins[index].Consumer.Id)
		if err != nil {
			return err, nil
		}
	}

	return nil, plugins
}
