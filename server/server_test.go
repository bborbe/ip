package server

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/server"
)

func TestImplementsServer(t *testing.T) {
	object := NewServer(8080)
	var expected *server.Server
	err := AssertThat(object, Implements(expected))
	if err != nil {
		t.Fatal(err)
	}
}
