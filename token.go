package msgraph

import (
	"fmt"

	"golang.org/x/oauth2"
)

func (api GraphAPI) GetToken() (*oauth2.Token, error) {
	var err error

	if api.token == nil {
		if err := api.validate(); err != nil {
			return nil, err
		}

		api.log.Debugf("Retrieving an OAuth2 token for tenant domain %v", api.TenantDomain)
		ctx := api.getContext()
		api.token, err = api.config.Token(ctx)
		if err != nil {
			return nil, &GraphAPIError{
				fmt.Sprintf("Retrieving token: %v", err),
				err}
		}

		api.log.Debugf("retrieveToken(): OAuth2 token for %s sucessfully retrieved", api.TenantDomain)
	}
	return api.token, nil
}

func (api *GraphAPI) SetToken(token *oauth2.Token) error {
	return &GraphAPIError{
		Message: "Unimplemented",
	}
}
