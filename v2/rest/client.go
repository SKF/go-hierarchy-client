package gohierarchyclient

import (
	"context"
	"fmt"

	"github.com/SKF/go-hierarchy-client/v2/rest/models"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/SKF/go-utility/v2/uuid"
)

type TreeFilter struct {
	Depth         int
	Limit         int
	Offset        int
	MetadataKey   string
	MetadataValue string
	NodeTypes     []string
	ModifiedAfter string
}

type HierarchyClient interface {
	GetNode(ctx context.Context, id uuid.UUID) (models.GetNodeResponse, error)
	CreateNode(ctx context.Context, node models.CreateNodeRequest) (models.CreateNodeResponse, error)
	UpdateNode(ctx context.Context, id uuid.UUID, node models.UpdateNodeRequest) (models.UpdateNodeResponse, error)
	DeleteNode(ctx context.Context, id uuid.UUID) error
	DuplicateNode(ctx context.Context, source uuid.UUID, destination uuid.UUID, suffix string) (models.DuplicateNodeResponse, error)

	GetAncestors(ctx context.Context, id uuid.UUID, height int, nodeTypes ...string) (models.GetAncestorsResponse, error)
	GetCompany(ctx context.Context, id uuid.UUID) (models.GetCompanyResponse, error)
	GetSubtree(ctx context.Context, id uuid.UUID, filter TreeFilter, continuationToken string) (models.GetSubtreeResponse, error)
	GetSubtreeCount(ctx context.Context, id uuid.UUID, nodeTypes ...string) (models.GetSubtreeCountResponse, error)

	LockNode(ctx context.Context, id uuid.UUID, recursive bool) error
	UnlockNode(ctx context.Context, id uuid.UUID, recursive bool) error

	GetOrigins(ctx context.Context, provider, continuationToken string, limit int) (models.GetOriginsResponse, error)
	GetOriginsByType(ctx context.Context, provider, originType, continuationToken string, limit int) (models.GetOriginsResponse, error)
	GetProviderNodeIDs(ctx context.Context, provider, continuationToken string, limit int) (models.GetNodesByPartialOriginResponse, error)
	GetProviderNodeIDsByType(ctx context.Context, provider, originType, continuationToken string, limit int) (models.GetNodesByPartialOriginResponse, error)
	GetOriginNodeID(ctx context.Context, origin models.Origin) (models.GetNodeByOriginResponse, error)
}

type client struct {
	*rest.Client
	clientID string
}

func WithStage(stage string) rest.Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://api.hierarchy.enlight.skf.com/v2/")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://api.%s.hierarchy.enlight.skf.com/v2/", stage))
}

func NewClient(clientID string, opts ...rest.Option) HierarchyClient {
	restClient := rest.NewClient(
		append([]rest.Option{
			// Defaults to production stage if no option is supplied
			WithStage(stages.StageProd),
		}, opts...)...,
	)

	return &client{Client: restClient, clientID: clientID}
}

func (c *client) GetNode(ctx context.Context, id uuid.UUID) (models.GetNodeResponse, error) {
	request := rest.Get("nodes/{node}").
		Assign("node", id).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetNodeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetNodeResponse{}, err
	}

	return response, nil
}

func (c *client) CreateNode(ctx context.Context, node models.CreateNodeRequest) (models.CreateNodeResponse, error) {
	request := rest.Post("nodes").
		WithJSONPayload(node).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.CreateNodeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.CreateNodeResponse{}, err
	}

	return response, nil
}

func (c *client) UpdateNode(ctx context.Context, id uuid.UUID, node models.UpdateNodeRequest) (models.UpdateNodeResponse, error) {
	request := rest.Patch("nodes/{node}").
		Assign("node", id).
		WithJSONPayload(node).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.UpdateNodeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.UpdateNodeResponse{}, err
	}

	return response, nil
}

func (c *client) DeleteNode(ctx context.Context, id uuid.UUID) (err error) {
	request := rest.Delete("nodes/{node}").
		Assign("node", id).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	_, err = c.Do(ctx, request)

	return
}

func (c *client) DuplicateNode(ctx context.Context, source uuid.UUID, destination uuid.UUID, suffix string) (models.DuplicateNodeResponse, error) {
	request := rest.Post("nodes/{node}/duplicate{?dstParentNodeId,label_suffix}").
		Assign("node", source).
		Assign("dstParentNodeId", destination).
		Assign("label_suffix", suffix).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.DuplicateNodeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.DuplicateNodeResponse{}, err
	}

	return response, nil
}

func (c *client) GetAncestors(ctx context.Context, id uuid.UUID, height int, nodeTypes ...string) (models.GetAncestorsResponse, error) {
	request := rest.Get("nodes/{node}/ancestors{?height,type*}").
		Assign("node", id).
		Assign("height", height).
		Assign("type", nodeTypes).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetAncestorsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetAncestorsResponse{}, err
	}

	return response, nil
}

func (c *client) GetCompany(ctx context.Context, id uuid.UUID) (models.GetCompanyResponse, error) {
	request := rest.Get("nodes/{node}/company").
		Assign("node", id).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetCompanyResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetCompanyResponse{}, err
	}

	return response, nil
}

func (c *client) GetSubtree(ctx context.Context, id uuid.UUID, filter TreeFilter, continuationToken string) (models.GetSubtreeResponse, error) {
	request := rest.Get("nodes/{node}/subtree{?depth,limit,offset,metadata_key,metadata_value,continuation_token,type*}").
		Assign("node", id).
		Assign("depth", filter.Depth).
		Assign("limit", filter.Limit).
		Assign("offset", filter.Offset).
		Assign("metadata_key", filter.MetadataKey).
		Assign("metadata_value", filter.MetadataValue).
		Assign("type", filter.NodeTypes).
		Assign("continuation_token", continuationToken).
		Assign("modified_after", filter.ModifiedAfter).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetSubtreeResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetSubtreeResponse{}, err
	}

	return response, nil
}

func (c *client) GetSubtreeCount(ctx context.Context, id uuid.UUID, nodeTypes ...string) (models.GetSubtreeCountResponse, error) {
	request := rest.Get("nodes/{node}/subtree/count{?type*}").
		Assign("node", id).
		Assign("type", nodeTypes).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetSubtreeCountResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetSubtreeCountResponse{}, err
	}

	return response, nil
}

func (c *client) GetOrigins(ctx context.Context, provider, continuationToken string, limit int) (models.GetOriginsResponse, error) {
	request := rest.Get("origin/{provider,continuation_token,limit}").
		Assign("provider", provider).
		Assign("continuation_token", continuationToken).
		Assign("limit", limit).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetOriginsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetOriginsResponse{}, err
	}

	return response, nil
}

func (c *client) GetOriginsByType(ctx context.Context, provider, originType, continuationToken string, limit int) (models.GetOriginsResponse, error) {
	request := rest.Get("origin/{provider}/{type}").
		Assign("provider", provider).
		Assign("type", originType).
		Assign("continuation_token", continuationToken).
		Assign("limit", limit).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetOriginsResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetOriginsResponse{}, err
	}

	return response, nil
}

func (c *client) GetProviderNodeIDs(ctx context.Context, provider, continuationToken string, limit int) (models.GetNodesByPartialOriginResponse, error) {
	request := rest.Get("origin/{provider}/nodes").
		Assign("provider", provider).
		Assign("continuation_token", continuationToken).
		Assign("limit", limit).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetNodesByPartialOriginResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetNodesByPartialOriginResponse{}, err
	}

	return response, nil
}

func (c *client) GetProviderNodeIDsByType(ctx context.Context, provider, originType, continuationToken string, limit int) (models.GetNodesByPartialOriginResponse, error) {
	request := rest.Get("origin/{provider}/{type}/nodes").
		Assign("provider", provider).
		Assign("type", originType).
		Assign("continuation_token", continuationToken).
		Assign("limit", limit).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetNodesByPartialOriginResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetNodesByPartialOriginResponse{}, err
	}

	return response, nil
}

func (c *client) GetOriginNodeID(ctx context.Context, origin models.Origin) (models.GetNodeByOriginResponse, error) {
	request := rest.Get("origin/{provider}/{type}/{id}/nodes").
		Assign("provider", origin.Provider).
		Assign("type", origin.Type).
		Assign("id", origin.ID).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	var response models.GetNodeByOriginResponse
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.GetNodeByOriginResponse{}, err
	}

	return response, nil
}

func (c *client) LockNode(ctx context.Context, id uuid.UUID, recursive bool) error {
	request := rest.Put("nodes/{node}/lock?recursive={recursive}").
		Assign("node", id).
		Assign("recursive", recursive).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	_, err := c.Do(ctx, request)

	return err
}
func (c *client) UnlockNode(ctx context.Context, id uuid.UUID, recursive bool) error {
	request := rest.Delete("nodes/{node}/lock?recursive={recursive}").
		Assign("node", id).
		Assign("recursive", recursive).
		SetHeader("Accept", "application/json").
		SetHeader("X-Client-Id", c.clientID)

	_, err := c.Do(ctx, request)

	return err
}
