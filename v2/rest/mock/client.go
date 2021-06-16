package mock

import (
	"context"

	rest "github.com/SKF/go-hierarchy-client/rest"
	"github.com/SKF/go-hierarchy-client/rest/models"

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

func (c *HierarchyClientMock) GetNode(ctx context.Context, id uuid.UUID) (models.Node, error) {
	args := c.Called(ctx, id)
	return args.Get(0).(models.Node), args.Error(1)
}

func (c *HierarchyClientMock) CreateNode(ctx context.Context, node models.WebmodelsNodeInput) (uuid.UUID, error) {
	args := c.Called(ctx, node)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (c *HierarchyClientMock) UpdateNode(ctx context.Context, id uuid.UUID, node models.WebmodelsNodeInput) (models.Node, error) {
	args := c.Called(ctx, id, node)
	return args.Get(0).(models.Node), args.Error(1)
}

func (c *HierarchyClientMock) DeleteNode(ctx context.Context, id uuid.UUID) error {
	args := c.Called(ctx, id)
	return args.Error(0)
}

func (c *HierarchyClientMock) DuplicateNode(ctx context.Context, source uuid.UUID, destination uuid.UUID) (uuid.UUID, error) {
	args := c.Called(ctx, source, destination)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (c *HierarchyClientMock) GetAncestors(ctx context.Context, id uuid.UUID, height int, nodeTypes ...string) ([]models.Node, error) {
	args := c.Called(ctx, id, height, nodeTypes)
	return args.Get(0).([]models.Node), args.Error(1)
}

func (c *HierarchyClientMock) GetCompany(ctx context.Context, id uuid.UUID) (models.Node, error) {
	args := c.Called(ctx, id)
	return args.Get(0).(models.Node), args.Error(1)
}

func (c *HierarchyClientMock) GetSubtree(ctx context.Context, id uuid.UUID, filter rest.TreeFilter) ([]models.Node, error) {
	args := c.Called(ctx, id, filter)
	return args.Get(0).([]models.Node), args.Error(1)
}

func (c *HierarchyClientMock) GetSubtreeCount(ctx context.Context, id uuid.UUID, nodeTypes ...string) (int64, error) {
	args := c.Called(ctx, id, nodeTypes)
	return args.Get(0).(int64), args.Error(1)
}

func (c *HierarchyClientMock) GetOrigins(ctx context.Context, provider string) ([]models.Origin, error) {
	args := c.Called(ctx, provider)
	return args.Get(0).([]models.Origin), args.Error(1)
}

func (c *HierarchyClientMock) GetOriginsByType(ctx context.Context, provider, originType string) ([]models.Origin, error) {
	args := c.Called(ctx, provider, originType)
	return args.Get(0).([]models.Origin), args.Error(1)
}

func (c *HierarchyClientMock) GetProviderNodeIDs(ctx context.Context, provider string) ([]uuid.UUID, error) {
	args := c.Called(ctx, provider)
	return args.Get(0).([]uuid.UUID), args.Error(1)
}

func (c *HierarchyClientMock) GetProviderNodeIDsByType(ctx context.Context, provider, originType string) ([]uuid.UUID, error) {
	args := c.Called(ctx, provider, originType)
	return args.Get(0).([]uuid.UUID), args.Error(1)
}

func (c *HierarchyClientMock) GetOriginNodeID(ctx context.Context, origin models.Origin) (uuid.UUID, error) {
	args := c.Called(ctx, origin)
	return args.Get(0).(uuid.UUID), args.Error(1)
}
