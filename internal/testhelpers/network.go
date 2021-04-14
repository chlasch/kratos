package testhelpers

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"

	"github.com/ory/kratos/persistence"
	"github.com/ory/x/networkx"
)

func NewNetwork(t *testing.T, p persistence.Persister) (uuid.UUID, persistence.Persister) {
	n := networkx.NewNetwork()
	require.NoError(t, p.GetConnection(context.Background()).Create(n))
	return n.ID, p.WithNetworkID(n.ID)
}

func ExistingNetwork(t *testing.T, p persistence.Persister, id uuid.UUID) persistence.Persister {
	require.NoError(t, p.GetConnection(context.Background()).Save(&networkx.Network{ID: id}))
	return p.WithNetworkID(id)
}