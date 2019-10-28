package api

import "github.com/mitchellh/mapstructure"

type Route struct {
	ID                      string              `json:"id,omitempty"`
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
	Service                 map[string]string   `json:"service,omitempty"`
}

type RespRoute struct {
	Next interface{} `json:"next"`
	Data []Route     `json:"data"`
}

func (a *Api) CreateRoute(route Route) error {
	if err := a.CreateEntity(route, "routes"); err != nil {
		return err
	}

	return nil
}

func (a *Api) ListRoutes() (error, *[]Route) {
	var err error

	err, svc := a.ListEntity("routes")
	if err != nil {
		return err, nil
	}

	var routes []Route

	err = mapstructure.Decode(svc, &routes)
	if err != nil {
		return err, nil
	}

	return nil, &routes
}
