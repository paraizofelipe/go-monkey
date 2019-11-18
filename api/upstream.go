package api

type HealthCheck struct {
	Active  interface{}
	Passive interface{}
}

type Upstream struct {
	Id               string `json:"id,omitempty"`
	CreatedAt        int64  `json:"create_at,omitempty"`
	Name             string `json:"name,omitempty"`
	Algorithm        string `json:"algorithm,omitempty"`
	HashOn           string `json:"hash_on,omitempty"`
	HashFallback     string `json:"hash_fallback,omitempty"`
	HashOnCookiePath string
	Slots            int64
	HealthCheck      HealthCheck
	Tags             []string
}
