package gohierarchyclient

import (
	"context"
	"fmt"

	"github.com/SKF/go-hierarchy-client/rest/models"

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
}

type HierarchyClient interface {
	GetNode(ctx context.Context, id uuid.UUID) (models.Node, error)
	CreateNode(ctx context.Context, node models.WebmodelsNodeInput) (uuid.UUID, error)
	UpdateNode(ctx context.Context, id uuid.UUID, node models.WebmodelsNodeInput) (models.Node, error)
	DeleteNode(ctx context.Context, id uuid.UUID) error
	DuplicateNode(ctx context.Context, source uuid.UUID, destination uuid.UUID) (uuid.UUID, error)

	GetAncestors(ctx context.Context, id uuid.UUID, height int, nodeTypes ...string) ([]models.Node, error)
	GetCompany(ctx context.Context, id uuid.UUID) (models.Node, error)
	GetSubtree(ctx context.Context, id uuid.UUID, filter TreeFilter) ([]models.Node, error)
	GetSubtreeCount(ctx context.Context, id uuid.UUID, nodeTypes ...string) (int64, error)
}

type client struct {
	*rest.Client
}

func WithStage(stage string) rest.Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://api.hierarchy.enlight.skf.com")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://api.%s.hierarchy.enlight.skf.com", stage))
}

func NewClient(opts ...rest.Option) HierarchyClient {
	restClient := rest.NewClient(
		append([]rest.Option{
			// Defaults to production stage if no option is supplied
			WithStage(stages.StageProd),
		}, opts...)...,
	)

	return &client{Client: restClient}
}

func (c *client) GetNode(ctx context.Context, id uuid.UUID) (models.Node, error) {
	request := rest.Get("nodes/{node}").
		Assign("node", id).
		SetHeader("Accept", "application/json")

	var response models.WebmodelsNode
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Node{}, err
	}

	return response.Node, nil
}

func (c *client) CreateNode(ctx context.Context, node models.WebmodelsNodeInput) (uuid.UUID, error) {
	request := rest.Put("nodes").
		WithJSONPayload(node).
		SetHeader("Accept", "application/json")

	var response models.WebmodelsNodeID
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return uuid.EmptyUUID, err
	}

	return response.NodeID, nil
}

func (c *client) UpdateNode(ctx context.Context, id uuid.UUID, node models.WebmodelsNodeInput) (models.Node, error) {
	request := rest.Post("nodes/{node}").
		Assign("node", id).
		WithJSONPayload(node).
		SetHeader("Accept", "application/json")

	var response models.WebmodelsNode
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Node{}, err
	}

	return response.Node, nil
}

func (c *client) DeleteNode(ctx context.Context, id uuid.UUID) (err error) {
	request := rest.Delete("nodes/{node}").
		Assign("node", id).
		SetHeader("Accept", "application/json")

	_, err = c.Do(ctx, request)

	return
}

func (c *client) DuplicateNode(ctx context.Context, source uuid.UUID, destination uuid.UUID) (uuid.UUID, error) {
	request := rest.Post("nodes/{node}/duplicate{?dstParentNodeId}").
		Assign("node", source).
		Assign("dstParentNodeId", destination).
		SetHeader("Accept", "application/json")

	var response models.WebmodelsNodeID
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return uuid.EmptyUUID, err
	}

	return response.NodeID, nil
}

func (c *client) GetAncestors(ctx context.Context, id uuid.UUID, height int, nodeTypes ...string) ([]models.Node, error) {
	request := rest.Get("nodes/{node}/ancestors{?height,type*}").
		Assign("node", id).
		Assign("height", height).
		Assign("type", nodeTypes).
		SetHeader("Accept", "application/json")

	var response models.WebmodelsNodes
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return nil, err
	}

	return response.Nodes, nil
}

func (c *client) GetCompany(ctx context.Context, id uuid.UUID) (models.Node, error) {
	request := rest.Get("nodes/{node}/company").
		Assign("node", id).
		SetHeader("Accept", "application/json")

	var response models.WebmodelsNode
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.Node{}, err
	}

	return response.Node, nil
}

func (c *client) GetSubtree(ctx context.Context, id uuid.UUID, filter TreeFilter) ([]models.Node, error) {
	request := rest.Get("nodes/{node}/subtree{?depth,limit,offset,metadata_key,metadata_value,type*}").
		Assign("node", id).
		Assign("depth", filter.Depth).
		Assign("limit", filter.Limit).
		Assign("offset", filter.Offset).
		Assign("metadata_key", filter.MetadataKey).
		Assign("metadata_value", filter.MetadataValue).
		Assign("type", filter.NodeTypes).
		SetHeader("Accept", "application/json")

	var response models.WebmodelsNodes
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return nil, err
	}

	return response.Nodes, nil
}

func (c *client) GetSubtreeCount(ctx context.Context, id uuid.UUID, nodeTypes ...string) (int64, error) {
	request := rest.Get("nodes/{node}/subtree/count{?type*}").
		Assign("node", id).
		Assign("type", nodeTypes).
		SetHeader("Accept", "application/json")

	var response models.WebmodelsNodeCount
	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return 0, err
	}

	return response.NodeCount, nil
}
