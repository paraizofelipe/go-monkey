package api

type Entity interface {
	GetValue(key string) interface{}
}
