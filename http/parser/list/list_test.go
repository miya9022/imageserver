package list

import (
	"testing"

	imageserver_http "github.com/pierrre/imageserver/http"
)

func TestTODO(t *testing.T) {
	t.Log("TODO")
}

func TestInterface(t *testing.T) {
	var _ imageserver_http.Parser = ListParser{}
	var _ imageserver_http.Resolver = ListParser{}
}
