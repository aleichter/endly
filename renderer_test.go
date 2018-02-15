package endly_test

import (
	"github.com/viant/endly"
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func Test_Renderer_PrintTable(t *testing.T) {

	var buf = new(bytes.Buffer)
	render := endly.NewRenderer(buf, 80)

	var data = [][]string{
		{
			"Id", "true", "Lorem Ipsum is simply dummy text ",
		},
		{
			"Name", "false", "Lorem Ipsum is simply dummy text of the printing and typesetting ",
		},
	}
	render.PrintTable("table1", []string{"name", "required", "description"}, data, 80)
	assert.True(t, len(buf.String()) > 0)
}
