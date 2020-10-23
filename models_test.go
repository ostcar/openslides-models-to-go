package models_test

import (
	"strings"
	"testing"

	models "github.com/OpenSlides/openslides-models-to-go"
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
			yml := strings.ReplaceAll(tt.yaml, "\t", " ")
			_, err := models.Unmarshal(strings.NewReader(yml))
			if err != nil {
				t.Errorf("Can not unmarshal yaml: %v", err)
			}
		})
	}
}

func TestRelation(t *testing.T) {
	yml := strings.ReplaceAll(yamlWithRelation, "\t", " ")
	got, err := models.Unmarshal(strings.NewReader(yml))
	if err != nil {
		t.Errorf("Can not unmarshal yml: %v", err)
	}

	if got["model"].Attributes["other_id"].Relation().List() {
		t.Errorf("model/other_id is a list")
	}

	if !got["model"].Attributes["other_ids"].Relation().List() {
		t.Errorf("model/other_ids is not a list")
	}

}

const yamlWithRelation = `---
model:
	other_id:
		type: relation
		to: not_existing/field
	other_ids:
		type: relation-list
		to: other/name
other:
	name: string
`
