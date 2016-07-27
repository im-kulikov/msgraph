// Package msgraphapi provides primitives for working with the Microsoft
// Graph API.
package msgraph

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft/clientcredentials"
)

type APIVersion int

const (
	APIVersionV1 APIVersion = iota
	APIVersionBeta
)

var (
	authURLPrefix = "https://login.microsoftonline.com"
	tokenURL      = "https://graph.microsoft.com"
)

// GraphAPI represents the Microsoft Graph API. It manages connections
// to Microsoft Graph.
type GraphAPI struct {
	TenantDomain string
	config       *clientcredentials.Config
	client       *http.Client
	token        *oauth2.Token
	httpDebug    bool
	log          *logrus.Logger
}

func (api *GraphAPI) SetDebug(debug bool) {
	if debug {
		api.log.Level = logrus.DebugLevel
		api.log.Debugf("Debug logging enabled")
	} else {
		api.log.Level = logrus.InfoLevel
	}
}

func (api *GraphAPI) SetHTTPDebug(debug bool) {
	if debug {
		api.log.Debugf("HTTP client debug logging enabled")
	}
	api.httpDebug = debug
}

func getAuthURL(tenantDomain string) string {
	return fmt.Sprintf("%s/%s/oauth2/token", authURLPrefix, tenantDomain)
}

func getTokenURL(tenantDomain string) string {
	return fmt.Sprintf("%s/%s", tokenURL, tenantDomain)
}

// New creates a new GraphAPI for the specified tenant domain.
func New(tenantDomain string) (api *GraphAPI) {
	log := logrus.New()
	log.Out = os.Stderr
	log.Debugf("Creating new GraphAPI for %s", tenantDomain)

	api = &GraphAPI{
		TenantDomain: tenantDomain,
		config: &clientcredentials.Config{
			Endpoint: oauth2.Endpoint{
				AuthURL:  getAuthURL(tenantDomain),
				TokenURL: tokenURL,
			},
		},
		log: log,
	}
	return
}

// SetClientID sets the OAuth2 "Client ID" to use for connections to the
// Microsoft Graph API.
func (api *GraphAPI) SetClientID(clientID string) {
	api.log.Debugf("Setting ClientID to %v", clientID)
	api.config.ClientID = clientID
}

// SetClientSecret sets the OAuth2 "Client Secret" to use for
// connections to the Microsoft Graph API.
func (api *GraphAPI) SetClientSecret(clientSecret string) {
	api.log.Debugf("Setting ClientSecret to %v", clientSecret)
	api.config.ClientSecret = clientSecret
}

func (api *GraphAPI) validate() error {
	if len(api.TenantDomain) == 0 {
		return &GraphAPIError{"Tenant domain must be set", nil}
	}
	if len(api.config.ClientID) == 0 {
		return &GraphAPIError{"ClientID must be set", nil}
	}
	if len(api.config.ClientSecret) == 0 {
		return &GraphAPIError{"Client secret must be set", nil}
	}
	api.log.Debug("validate(): GraphAPI validation successful.")
	return nil
}

func (api *GraphAPI) getContext() context.Context {
	ctx := oauth2.NoContext
	if api.httpDebug {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
			Transport: &logTransport{http.DefaultTransport},
		})
	}
	return ctx
}

func (api *GraphAPI) Client() (*http.Client, error) {
	if api.client == nil {
		api.log.Debugf("Creating a new http.Client")

		ctx := api.getContext()
		api.client = api.config.Client(ctx)
	}
	return api.client, nil
}
