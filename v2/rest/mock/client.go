package mock

import (
	"context"

	rest "github.com/SKF/go-hierarchy-client/v2/rest"
	"github.com/SKF/go-hierarchy-client/v2/rest/models"

	"github.com/SKF/go-utility/v2/uuid"
	"github.com/stretchr/testify/mock"
)

type HierarchyClientMock struct {
	*mock.Mock
}

func NewHierarchyClient() *HierarchyClientMock {
	client := &HierarchyClientMock{
		Mock: &mock.Mock{},
	}

	// Ensure the returned mock implements the HierarchyClient interface
	var _ rest.HierarchyClient = client

	return client
}

func (c *HierarchyClientMock) GetNode(ctx context.Context, id uuid.UUID) (models.GetNodeResponse, error) {
	args := c.Called(ctx, id)
	return args.Get(0).(models.GetNodeResponse), args.Error(1)
}

func (c *HierarchyClientMock) CreateNode(ctx context.Context, node models.CreateNodeRequest) (models.CreateNodeResponse, error) {
	args := c.Called(ctx, node)
	return args.Get(0).(models.CreateNodeResponse), args.Error(1)
}

func (c *HierarchyClientMock) UpdateNode(ctx context.Context, id uuid.UUID, node models.UpdateNodeRequest) (models.UpdateNodeResponse, error) {
	args := c.Called(ctx, id, node)
	return args.Get(0).(models.UpdateNodeResponse), args.Error(1)
}

func (c *HierarchyClientMock) DeleteNode(ctx context.Context, id uuid.UUID) error {
	args := c.Called(ctx, id)
	return args.Error(0)
}

func (c *HierarchyClientMock) DuplicateNode(ctx context.Context, source uuid.UUID, destination uuid.UUID, suffix string) (models.DuplicateNodeResponse, error) {
	args := c.Called(ctx, source, destination, suffix)
	return args.Get(0).(models.DuplicateNodeResponse), args.Error(1)
}

func (c *HierarchyClientMock) GetAncestors(ctx context.Context, id uuid.UUID, height int, nodeTypes ...string) (models.GetAncestorsResponse, error) {
	args := c.Called(ctx, id, height, nodeTypes)
	return args.Get(0).(models.GetAncestorsResponse), args.Error(1)
}

func (c *HierarchyClientMock) GetCompany(ctx context.Context, id uuid.UUID) (models.GetCompanyResponse, error) {
	args := c.Called(ctx, id)
	return args.Get(0).(models.GetCompanyResponse), args.Error(1)
}

func (c *HierarchyClientMock) GetSubtree(ctx context.Context, id uuid.UUID, filter rest.TreeFilter, continuationToken string) (models.GetSubtreeResponse, error) {
	args := c.Called(ctx, id, filter, continuationToken)
	return args.Get(0).(models.GetSubtreeResponse), args.Error(1)
}

func (c *HierarchyClientMock) GetSubtreeCount(ctx context.Context, id uuid.UUID, nodeTypes ...string) (models.GetSubtreeCountResponse, error) {
	args := c.Called(ctx, id, nodeTypes)
	return args.Get(0).(models.GetSubtreeCountResponse), args.Error(1)
}

func (c *HierarchyClientMock) LockNode(ctx context.Context, id uuid.UUID, recursive bool) error {
	args := c.Called(ctx, id, recursive)
	return args.Error(0)
}

func (c *HierarchyClientMock) UnlockNode(ctx context.Context, id uuid.UUID, recursive bool) error {
	args := c.Called(ctx, id, recursive)
	return args.Error(0)
}

func (c *HierarchyClientMock) GetOrigins(ctx context.Context, provider, continuationToken string, limit int) (models.GetOriginsResponse, error) {
	args := c.Called(ctx, provider, continuationToken, limit)
	return args.Get(0).(models.GetOriginsResponse), args.Error(1)
}

func (c *HierarchyClientMock) GetOriginsByType(ctx context.Context, provider, originType, continuationToken string, limit int) (models.GetOriginsResponse, error) {
	args := c.Called(ctx, provider, originType, continuationToken, limit)
	return args.Get(0).(models.GetOriginsResponse), args.Error(1)
}

func (c *HierarchyClientMock) GetProviderNodeIDs(ctx context.Context, provider, continuationToken string, limit int) (models.GetNodesByPartialOriginResponse, error) {
	args := c.Called(ctx, provider, continuationToken, limit)
	return args.Get(0).(models.GetNodesByPartialOriginResponse), args.Error(1)
}

func (c *HierarchyClientMock) GetProviderNodeIDsByType(ctx context.Context, provider, originType, continuationToken string, limit int) (models.GetNodesByPartialOriginResponse, error) {
	args := c.Called(ctx, provider, originType, continuationToken, limit)
	return args.Get(0).(models.GetNodesByPartialOriginResponse), args.Error(1)
}

func (c *HierarchyClientMock) GetOriginNodeID(ctx context.Context, origin models.Origin) (models.GetNodeByOriginResponse, error) {
	args := c.Called(ctx, origin)
	return args.Get(0).(models.GetNodeByOriginResponse), args.Error(1)
}
