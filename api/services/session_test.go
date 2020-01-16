package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vvelikodny/ff-go-test/api/services"
)

func TestNewInMemSessionService(t *testing.T) {
	s := services.NewInMemSessionService()

	require.True(t, s.Register("123"))
	require.False(t, s.Register("123"))
	require.False(t, s.Register("123"))
}
