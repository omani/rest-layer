package schema

import "errors"

// Reference validates time based values
type Reference struct {
	Path string
}

// Validate validates and normalize reference based value
func (v Reference) Validate(value interface{}) (interface{}, error) {
	// All the work is performed in rest.checkReferences()
	return value, nil
}

// ReferenceArray validates time based values
type ReferenceArray struct {
	Path string
}

// Validate validates and normalize reference based value
func (v ReferenceArray) Validate(values interface{}) (interface{}, error) {
	// Check if value is an array
	_, ok := values.([]interface{})
	if !ok {
		return nil, errors.New("not an array")
	}
	// All the work is performed in rest.checkReferences()
	return values, nil
}
