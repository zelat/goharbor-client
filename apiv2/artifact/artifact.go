package artifact

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime"
	modelv1 "github.com/zelat/goharbor-client/apiv1/model"
	"github.com/zelat/goharbor-client/apiv2/config"
	"github.com/zelat/goharbor-client/apiv2/errors"
	v2client "github.com/zelat/goharbor-client/apiv2/internal/api/client"
	"github.com/zelat/goharbor-client/apiv2/internal/api/client/artifact"
	"github.com/zelat/goharbor-client/apiv2/internal/legacyapi/client"
	"github.com/zelat/goharbor-client/apiv2/model"
)

// RESTClient is a subclient for handling artifact related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
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
	AddArtifactLabel(ctx context.Context, projectName, repositoryName, reference string, label *model.Label) error
	CopyArtifact(ctx context.Context, from *CopyReference, projectName, repositoryName string) error
	CreateTag(ctx context.Context, projectName, repositoryName, reference string, tag *model.Tag) error
	DeleteTag(ctx context.Context, projectName, repositoryName, reference, tagName string) error
	GetArtifact(ctx context.Context, projectName, repositoryName, reference string) (*model.Artifact, error)
	ListArtifacts(ctx context.Context, projectName, repositoryName string) ([]*model.Artifact, error)
	ListTags(ctx context.Context, projectName, repositoryName, reference string) ([]*model.Tag, error)
	RemoveLabel(ctx context.Context, projectName, repositoryName, reference string, id int64) error
	// TODO: Introduce this, once https://github.com/goharbor/harbor/issues/13468 is resolved.
	GetAddition(ctx context.Context, projectName, repositoryName, reference string, addition Addition) (string, error)
	GetVulnerabilitiesAddition(ctx context.Context, projectName, repositoryName, reference string) (*modelv1.TopReport, error)
}

// ToString returns a string representation of a CopyReference.
// Possible formats are "project/repository:tag" or "project/repository@digest".
// Returns an error if neither tag nor digest is set.
func (in CopyReference) toString() (string, error) {
	var suffix string

	if in.Digest == "" && in.Tag == "" {
		return "", fmt.Errorf("no tag or digest specified")
	}

	if in.Digest != "" {
		suffix = "@" + in.Digest
	}
	if in.Tag != "" {
		suffix = ":" + in.Tag
	}

	return in.ProjectName + "/" + in.RepositoryName + suffix, nil
}

// CopyReference defines the parameters needed for an artifact copy.
type CopyReference struct {
	ProjectName    string
	RepositoryName string
	Tag            string
	Digest         string
}

type Addition string

const (
	AdditionBuildHistory Addition = "build_history"
	AdditionValuesYAML   Addition = "values.yaml"
	AdditionReadme       Addition = "readme.md"
	AdditionDependencies Addition = "dependencies"
)

func (c *RESTClient) ListArtifacts(ctx context.Context, projectName, repositoryName string) ([]*model.Artifact, error) {
	params := &artifact.ListArtifactsParams{
		Page:           &c.Options.Page,
		PageSize:       &c.Options.PageSize,
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Q:              &c.Options.Query,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)
	artifact, err := c.V2Client.Artifact.ListArtifacts(params, c.AuthInfo)
	if err != nil {
		fmt.Print(err.Error())
		return nil, handleSwaggerArtifactErrors(err)
	}

	if artifact.Payload != nil {
		return artifact.Payload, nil
	}

	return nil, &errors.ErrNotFound{}
}

func (c *RESTClient) AddArtifactLabel(ctx context.Context, projectName, repositoryName, reference string, label *model.Label) error {
	params := &artifact.AddLabelParams{
		Label:          label,
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Artifact.AddLabel(params, c.AuthInfo)

	return handleSwaggerArtifactErrors(err)
}

func (c *RESTClient) CopyArtifact(ctx context.Context, from *CopyReference, projectName, repositoryName string) error {
	f, err := from.toString()
	if err != nil {
		return err
	}

	params := &artifact.CopyArtifactParams{
		From:           f,
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err = c.V2Client.Artifact.CopyArtifact(params, c.AuthInfo)

	return handleSwaggerArtifactErrors(err)
}

func (c *RESTClient) CreateTag(ctx context.Context, projectName, repositoryName, reference string, tag *model.Tag) error {
	params := &artifact.CreateTagParams{
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		Tag:            tag,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Artifact.CreateTag(params, c.AuthInfo)

	return handleSwaggerArtifactErrors(err)
}

func (c *RESTClient) DeleteTag(ctx context.Context, projectName, repositoryName, reference, tagName string) error {
	params := &artifact.DeleteTagParams{
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		TagName:        tagName,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Artifact.DeleteTag(params, c.AuthInfo)
	return handleSwaggerArtifactErrors(err)
}

func (c *RESTClient) GetArtifact(ctx context.Context, projectName, repositoryName, reference string) (*model.Artifact, error) {
	params := &artifact.GetArtifactParams{
		Page:           &c.Options.Page,
		PageSize:       &c.Options.PageSize,
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Reference:      reference,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Artifact.GetArtifact(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerArtifactErrors(err)
	}

	if resp.Payload != nil {
		return resp.Payload, nil
	}

	return nil, &errors.ErrNotFound{}
}

func (c *RESTClient) ListTags(ctx context.Context, projectName, repositoryName, reference string) ([]*model.Tag, error) {
	params := artifact.NewListTagsParams()
	params.Page = &c.Options.Page
	params.PageSize = &c.Options.PageSize
	params.WithProjectName(projectName)
	params.WithRepositoryName(repositoryName)
	params.WithReference(reference)
	params.Q = &c.Options.Query
	params.WithContext(ctx)
	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Artifact.ListTags(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerArtifactErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) RemoveLabel(ctx context.Context, projectName, repositoryName, reference string, id int64) error {
	params := &artifact.RemoveLabelParams{
		LabelID:        id,
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Reference:      reference,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Artifact.RemoveLabel(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerArtifactErrors(err)
	}

	return nil
}

func (c *RESTClient) GetAddition(ctx context.Context, projectName, repositoryName, reference string, addition Addition) (interface{}, error) {
	params := &artifact.GetAdditionParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Reference:      reference,
		Addition:       string(addition),
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Artifact.GetAddition(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerArtifactErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) GetVulnerabilitiesAddition(ctx context.Context, projectName, repositoryName, reference string) (*modelv1.TopReport, error) {
	params := &artifact.GetVulnerabilitiesAdditionParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Reference:      reference,
		Context:        ctx,
	}
	params.WithTimeout(c.Options.Timeout)
	xAcceptVulnerabilities := "application/vnd.security.vulnerability.report; version=1.1, application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0"
	params.WithXAcceptVulnerabilities(&xAcceptVulnerabilities)

	resp, err := c.V2Client.Artifact.GetVulnerabilitiesAddition(params, c.AuthInfo)

	if err != nil {
		return nil, handleSwaggerArtifactErrors(err)
	}

	return resp.Payload, nil
}
