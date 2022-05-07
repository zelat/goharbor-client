//go:build integration
// +build integration

package system

import (
	"context"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
	v2client "github.com/zelat/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/zelat/goharbor-client/v4/apiv2/internal/legacyapi/client"
	integrationtest "github.com/zelat/goharbor-client/v4/apiv2/testing"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
)

func TestAPIHealth(t *testing.T) {
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	_, err := c.Health(ctx)
	require.NoError(t, err)
}
