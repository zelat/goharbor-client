package scan

import (
	"context"
	"github.com/go-openapi/runtime"
	"github.com/zelat/goharbor-client/apiv2/config"
	v2client "github.com/zelat/goharbor-client/apiv2/internal/api/client"
	"github.com/zelat/goharbor-client/apiv2/internal/api/client/scan"
	"github.com/zelat/goharbor-client/apiv2/internal/legacyapi/client"
)

// RESTClient is a subclient for handling registry related actions.
type RESTClient struct {
	Options *config.Options

	LegacyClient *client.Harbor

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:      opts,
		LegacyClient: legacyClient,
		V2Client:     v2Client,
		AuthInfo:     authInfo,
	}
}

type Client interface {
	ScanArtifact(ctx context.Context, projectName, repositoryName, reference string) (string, error)
	GetReportLog(ctx context.Context, projectName, repositoryName, reference, reportID string) (string, error)
}

func (c *RESTClient) ScanArtifact(ctx context.Context, projectName, repositoryName, reference string) (interface{}, error) {
	params := &scan.ScanArtifactParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Reference:      reference,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Scan.ScanArtifact(params, c.AuthInfo)
	if err != nil {
		return nil, err
	}

	return resp.XRequestID, nil

}

func (c *RESTClient) GetReportLog(ctx context.Context, projectName, repositoryName, reference, reportID string) (interface{}, error) {
	params := &scan.GetReportLogParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Reference:      reference,
		ReportID:       reportID,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Scan.GetReportLog(params, c.AuthInfo)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil

}
