package schema

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/qri-io/jsonschema"
)

func ValidateJSONWithSchema(schemaData []byte, jsonData []byte) error {
	rs := &jsonschema.Schema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		return fmt.Errorf("error unmarshaling schema: %w", err)
	}

	errs, err := rs.ValidateBytes(context.Background(), jsonData)
	if err != nil {
		return fmt.Errorf("error validating schema: %w", err)
	}
	if len(errs) > 0 {
		return fmt.Errorf("schema validation failed: %s", errs[0].Message)
	}

	return nil
}
