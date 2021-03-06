package mapper

import (
	"fmt"

	"github.com/rancher/norman/types"
)

type Drop struct {
	Field string
}

func (d Drop) FromInternal(data map[string]interface{}) {
	delete(data, d.Field)
}

func (d Drop) ToInternal(data map[string]interface{}) {
}

func (d Drop) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	if _, ok := schema.ResourceFields[d.Field]; !ok {
		return fmt.Errorf("can not drop missing field %s on %s", d.Field, schema.ID)
	}

	delete(schema.ResourceFields, d.Field)
	return nil
}
