package services_test

import (
	"github.com/stretchr/testify/require"
	"github.com/vvelikodny/ff-go-test/api/services"
	"testing"
)

func TestNewInMemSessionService(t *testing.T) {
	s := services.NewInMemSessionService()

	require.True(t, s.Register("123"))
	require.False(t, s.Register("123"))
	require.False(t, s.Register("123"))
}
