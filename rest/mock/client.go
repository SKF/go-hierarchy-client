package mock

import (
	"context"

	rest "github.com/SKF/go-hierarchy-client/rest"
	"github.com/SKF/go-hierarchy-client/rest/models"

	"github.com/SKF/go-utility/v2/uuid"
	"github.com/stretchr/testify/mock"
)

type client struct {
	*mock.Mock
}

func NewHierarchyClient() rest.HierarchyClient {
	return &client{}
}

func (c *client) GetNode(ctx context.Context, id uuid.UUID) (models.Node, error) {
	args := c.Called(ctx, id)
	return args.Get(0).(models.Node), args.Error(1)
}

func (c *client) CreateNode(ctx context.Context, node models.WebmodelsNodeInput) (uuid.UUID, error) {
	args := c.Called(ctx, node)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (c *client) UpdateNode(ctx context.Context, id uuid.UUID, node models.WebmodelsNodeInput) (models.Node, error) {
	args := c.Called(ctx, id, node)
	return args.Get(0).(models.Node), args.Error(1)
}

func (c *client) DeleteNode(ctx context.Context, id uuid.UUID) error {
	args := c.Called(ctx, id)
	return args.Error(0)
}

func (c *client) DuplicateNode(ctx context.Context, source uuid.UUID, destination uuid.UUID) (uuid.UUID, error) {
	args := c.Called(ctx, source, destination)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (c *client) GetAncestors(ctx context.Context, id uuid.UUID, height int, nodeTypes ...string) ([]models.Node, error) {
	args := c.Called(ctx, id, height, nodeTypes)
	return args.Get(0).([]models.Node), args.Error(1)
}

func (c *client) GetCompany(ctx context.Context, id uuid.UUID) (models.Node, error) {
	args := c.Called(ctx, id)
	return args.Get(0).(models.Node), args.Error(1)
}

func (c *client) GetSubtree(ctx context.Context, id uuid.UUID, filter rest.TreeFilter) ([]models.Node, error) {
	args := c.Called(ctx, id, filter)
	return args.Get(0).([]models.Node), args.Error(1)
}

func (c *client) GetSubtreeCount(ctx context.Context, id uuid.UUID, nodeTypes ...string) (int64, error) {
	args := c.Called(ctx, id, nodeTypes)
	return args.Get(0).(int64), args.Error(1)
}
