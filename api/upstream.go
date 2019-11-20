package api

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Upstream struct {
	Id               string       `json:"id,omitempty"`
	CreatedAt        int64        `json:"created_at,omitempty"`
	Name             string       `json:"name,omitempty"`
	Algorithm        string       `json:"algorithm,omitempty"`
	HashOn           string       `json:"hash_on,omitempty"`
	HashFallback     string       `json:"hash_fallback,omitempty"`
	HashOnCookiePath string       `json:"hash_on_cookie_path,omitempty"`
	Slots            int64        `json:"slots,omitempty"`
	HealthChecks     HealthChecks `json:"healthchecks,omitempty"`
	Tags             []string     `json:"tags,omitempty"`
}

type Unhealthy struct {
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
	TCPFailures  int64   `json:"tcp_failures,omitempty"`
	Timeouts     int64   `json:"timeouts,omitempty"`
	HTTPFailures int64   `json:"http_failures,omitempty"`
	Interval     int64   `json:"interval,omitempty"`
}

type Healthy struct {
	Successes    int64   `json:"successes,omitempty"`
	Interval     int64   `json:"interval,omitempty"`
	HTTPStatuses []int64 `json:"http_statuses,omitempty"`
}

type Active struct {
	HTTPSVerifyCertificate bool      `json:"https_verify_certificate,omitempty"`
	Unhealthy              Unhealthy `json:"unhealthy,omitempty"`
	HTTPPath               string    `json:"http_path,omitempty"`
	Timeout                int64     `json:"timeout,omitempty"`
	Healthy                Healthy   `json:"healthy,omitempty"`
	HTTPSSni               string    `json:"https_sni,omitempty"`
	Concurrency            int64     `json:"concurrency,omitempty"`
	Type                   string    `json:"type,omitempty"`
}

type Passive struct {
	Unhealthy Unhealthy `json:"unhealthy,omitempty"`
	Type      string    `json:"type,omitempty"`
	Healthy   Healthy   `json:"healthy,omitempty"`
}

type HealthChecks struct {
	Active  Active  `json:"active,omitempty"`
	Passive Passive `json:"passive,omitempty"`
}

func (u *Upstream) GetValue(key string) interface{} {
	rt := reflect.ValueOf(u)
	k := strings.Title(strings.ToLower(key))
	f := reflect.Indirect(rt).FieldByName(k)
	return f
}

func (u *Upstream) ToMap() (error, map[string]interface{}) {
	b, err := json.Marshal(u)
	if err != nil {
		return err, nil
	}

	var ms map[string]interface{}
	if err := json.Unmarshal(b, &ms); err != nil {
		return err, nil
	}

	return nil, ms
}

func (a *Api) CreateUpstream(upstream Upstream) error {
	if err := a.CreateEntity(&upstream, "upstreams"); err != nil {
		return err
	}

	return nil
}

func (a *Api) Upstream(id string) (error, Upstream) {
	var err error
	var upstream Upstream

	err, svc := a.GetEntity("upstreams", id)
	if err != nil {
		return err, upstream
	}

	err = mapstructure.Decode(svc, &upstream)
	if err != nil {
		return err, upstream
	}

	return nil, upstream
}

func (a *Api) Upstreams() (error, []Upstream) {
	var err error

	err, rts := a.ListEntity("upstreams")
	if err != nil {
		return err, nil
	}

	var upstreams []Upstream

	err = mapstructure.Decode(rts, &upstreams)
	if err != nil {
		return err, nil
	}

	return nil, upstreams
}
