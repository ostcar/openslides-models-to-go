package models_test

import (
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-modelsvalidate/models"
)

func TestUnmarshal(t *testing.T) {
	for _, tt := range []struct {
		name string
		yaml string
	}{
		{
			"With Relation",
			yamlWithRelation,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := models.Unmarshal(strings.NewReader(tt.yaml))
			if err != nil {
				t.Errorf("Can not unmarshal yaml: %v", err)
			}
		})
	}
}

const yamlWithRelation = `---
some_model:
  no_other_model:
    type: relation
    to: not_existing/field
  no_other_field:
    type: relation
    to: other_model/bar
other_model:
  foo: string
`
