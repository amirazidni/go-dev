package test

import (
	"go-dev/dependency-injection/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databaes")
	assert.NotNil(t, connection)

	cleanup()
}
