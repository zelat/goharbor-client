//go:build integration
// +build integration

package gc

import (
	"context"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
	v2client "github.com/zelat/goharbor-client/apiv2/internal/api/client"
	"github.com/zelat/goharbor-client/apiv2/internal/legacyapi/client"
	modelv2 "github.com/zelat/goharbor-client/apiv2/model"
	integrationtest "github.com/zelat/goharbor-client/apiv2/testing"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
)

func TestAPINewGarbageCollection(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	err := c.NewGarbageCollection(ctx, &modelv2.Schedule{
		Schedule: &modelv2.ScheduleObj{
			Cron: "0 * * * * *",
			Type: "Hourly",
		},
	})

	defer c.ResetGarbageCollection(ctx)

	require.NoError(t, err)
}

func TestAPIUpdateGarbageCollection(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	err := c.UpdateGarbageCollection(ctx, &modelv2.Schedule{
		Schedule: &modelv2.ScheduleObj{
			Cron: "0 * * * * *",
			Type: "Hourly",
		},
	})

	defer c.ResetGarbageCollection(ctx)

	require.NoError(t, err)
}

func TestResetGarbageCollection(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	err := c.ResetGarbageCollection(ctx)

	require.NoError(t, err)
}

func TestGetGarbageCollectionSchedule(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	sched := &modelv2.ScheduleObj{
		Cron: "0 * * * * *",
		Type: "Hourly",
	}

	err := c.NewGarbageCollection(ctx, &modelv2.Schedule{
		Schedule: sched,
	})

	require.NoError(t, err)

	gc, err := c.GetGarbageCollectionSchedule(ctx)

	require.NoError(t, err)
	require.NotNil(t, gc)

	require.Equal(t, gc.Schedule, sched)

	defer c.ResetGarbageCollection(ctx)
}
