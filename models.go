package resource

import "time"

type Source struct {
	URI      string `json:"uri"`
	Apikey   string `json:"apikey"`
	Insecure bool   `json:"insecure"`
}

type Version struct {
	Timestamp time.Time `json:"timestamp"`
}

type MetadataPair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
