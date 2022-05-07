package retention

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/zelat/goharbor-client/v4/apiv2/internal/api/client/retention"
	modelv2 "github.com/zelat/goharbor-client/v4/apiv2/model"

	"github.com/go-openapi/runtime"
	v2client "github.com/zelat/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/zelat/goharbor-client/v4/apiv2/internal/legacyapi/client"
	projectapi "github.com/zelat/goharbor-client/v4/apiv2/project"
)

const (
	// AlgorithmOr is the default algorithm when operating on harbor retention rules
	AlgorithmOr string = "or"

	// Key for defining matching repositories
	ScopeSelectorRepoMatches ScopeSelector = "repoMatches"

	// Key for defining excluded repositories
	ScopeSelectorRepoExcludes ScopeSelector = "repoExcludes"

	// Key for defining matching tag expressions
	TagSelectorMatches TagSelector = "matches"

	// Key for defining excluded tag expressions
	TagSelectorExcludes TagSelector = "excludes"

	// The kind of the retention selector, _always_ defaults to 'doublestar'
	SelectorTypeDefault string = "doublestar"

	// Retain the most recently pushed n artifacts - count
	PolicyTemplateLatestPushedArtifacts PolicyTemplate = "latestPushedK"

	// Retain the most recently pulled n artifacts - count
	PolicyTemplateLatestPulledArtifacts PolicyTemplate = "latestPulledN"

	// Retain the artifacts pushed within the last n days
	PolicyTemplateDaysSinceLastPush PolicyTemplate = "nDaysSinceLastPush"

	// Retain the artifacts pulled within the last n days
	PolicyTemplateDaysSinceLastPull PolicyTemplate = "nDaysSinceLastPull"

	// Retain always
	PolicyTemplateRetainAlways PolicyTemplate = "always"
)

type Client interface {
	NewRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error
	GetRetentionPolicyByProject(ctx context.Context, project *modelv2.Project) (*modelv2.RetentionPolicy, error)
	GetRetentionPolicyByID(ctx context.Context, id int64) (*modelv2.RetentionPolicy, error)
	DisableRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error
	UpdateRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error
}

// RESTClient is a subclient for handling retention related actions.
type RESTClient struct {
	// The swagger client
	LegacyClient *client.Harbor

	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		LegacyClient: legacyClient,
		V2Client:     v2Client,
		AuthInfo:     authInfo,
	}
}

// ScopeSelector is the retention selector decoration used for operations on retention objects.
type ScopeSelector string

func (r ScopeSelector) String() string {
	return string(r)
}

// PolicyTemplate defines the possible values used for the policy matching mechanism.
type PolicyTemplate string

func (p PolicyTemplate) String() string {
	return string(p)
}

// TagSelector defines the possible values used for the tag matching mechanism. Valid values are: "matches, excludes".
type TagSelector string

// String returns the string value of a TagSelector.
func (t TagSelector) String() string {
	return string(t)
}

// NewRetentionPolicy creates a new tag retention policy for a project.
func (c *RESTClient) NewRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	if ret == nil {
		return &ErrRetentionNotProvided{}
	}

	_, err := c.V2Client.Retention.CreateRetention(&retention.CreateRetentionParams{
		Policy:  ret,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return handleSwaggerRetentionErrors(err)
	}

	return nil
}

// GetRetentionPolicyByProject returns a retention policy that is fetched by the metadata value contained in a project's metadata.
func (c *RESTClient) GetRetentionPolicyByProject(ctx context.Context, project *modelv2.Project) (*modelv2.RetentionPolicy, error) {
	pc := projectapi.NewClient(c.LegacyClient, c.V2Client, c.AuthInfo)

	val, err := pc.GetProjectMetadataValue(ctx, fmt.Sprint(project.ProjectID), projectapi.ProjectMetadataKeyRetentionID)
	if err != nil {
		return nil, err
	}

	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("could not convert retention id %q to int64, project: %d", val, project.ProjectID)
	}

	return c.GetRetentionPolicyByID(ctx, id)
}

// GetRetentionPolicyByID returns a retention policy identified by it's id.
func (c *RESTClient) GetRetentionPolicyByID(ctx context.Context, id int64) (*modelv2.RetentionPolicy, error) {
	resp, err := c.V2Client.Retention.GetRetention(&retention.GetRetentionParams{
		ID:      id,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRetentionErrors(err)
	}

	return resp.Payload, nil
}

// DisableRetentionPolicy replaces the rules of a retention policy with an empty set of rules.
// As of Harbor v2.2.1, the swagger specifications do not contain a DELETE route for
// retention policies, but instead PUT a retention policy with a dummy retention rule.
// This function provides the same functionality as "Action -> Delete" when editing retention rules in the GUI.
// TODO: replace this method with a new one 'DeleteRetentionPolicy'
// once the delete route for the retention API [^1] has been released.
// [^1]: https://github.com/goharbor/harbor/pull/14747/commits/81e5aa715b98b39fbf729048c34fe46c3af31505
func (c *RESTClient) DisableRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	if ret == nil {
		return &ErrRetentionNotProvided{}
	}

	resp, err := c.V2Client.Retention.UpdateRetention(&retention.UpdateRetentionParams{
		ID: ret.ID,
		Policy: &modelv2.RetentionPolicy{
			Algorithm: ret.Algorithm,
			ID:        ret.ID,
			Rules:     []*modelv2.RetentionRule{},
			Scope:     ret.Scope,
			Trigger:   ret.Trigger,
		},
		Context: ctx,
	}, c.AuthInfo)

	if resp == nil {
		return &ErrRetentionDoesNotExist{}
	}

	if err != nil {
		return handleSwaggerRetentionErrors(err)
	}

	return nil
}

// UpdateRetentionPolicy updates the specified retention policy ret.
func (c *RESTClient) UpdateRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	if ret == nil {
		return &ErrRetentionNotProvided{}
	}

	resp, err := c.V2Client.Retention.UpdateRetention(&retention.UpdateRetentionParams{
		ID:      ret.ID,
		Policy:  ret,
		Context: ctx,
	}, c.AuthInfo)

	if resp == nil {
		return &ErrRetentionDoesNotExist{}
	}

	if err != nil {
		return handleSwaggerRetentionErrors(err)
	}

	return nil
}

// ToTagSelectorExtras converts a boolean to the representative string value used by Harbor.
// Represents the functionality of the 'untagged artifacts' checkbox when editing tag retention rules in the Harbor UI.
func ToTagSelectorExtras(untagged bool) string {
	if untagged {
		return `{"untagged":"true"}`
	}
	return `{"untagged":"false"}`
}

// evaluateRetentionRuleParams evaluates the provided map of PolicyTemplate by comparing the keys to the pre-defined PolicyTemplates.
// Returns an error if the provided or resulting map is empty.
func evaluateRetentionRuleParams(params map[PolicyTemplate]interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})

	if len(params) == 0 {
		return nil, errors.New("no retention rule params provided")
	}

	for k, v := range params {
		if _, ok := params[k]; ok {
			switch k {
			case PolicyTemplateDaysSinceLastPull:
				res[k.String()] = v
			case PolicyTemplateDaysSinceLastPush:
				res[k.String()] = v
			case PolicyTemplateLatestPulledArtifacts:
				res[k.String()] = v
			case PolicyTemplateLatestPushedArtifacts:
				res[k.String()] = v
			default:
				continue
			}
		}
	}

	if len(res) == 0 {
		return nil, errors.New("invalid params provided")
	}

	return res, nil
}
