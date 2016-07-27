package msgraph

import (
	"fmt"
	"net/url"
)

var resources map[string]Resource

func init() {
	resources = make(map[string]Resource, 10)
}

type Resource struct {
	Name       string
	APIVersion APIVersion
	Resource   string
}

func (api GraphAPI) GetResourceEndpoint(r Resource) *url.URL {
	var v string

	// TODO: Clean up the whole APIVersion thing.
	switch r.APIVersion {
	case APIVersionV1:
		v = "v1.0"
		break
	case APIVersionBeta:
		v = "beta"

	}

	u, err := url.Parse(fmt.Sprintf("%s/%s/%s", api.config.Endpoint.TokenURL, v, r.Resource))
	if err != nil {
		api.log.Fatal("Creating endpoint: %v", err)
	}

	return u
}
