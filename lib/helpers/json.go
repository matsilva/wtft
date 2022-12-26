package helpers

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

// ToJSON converts a struct to JSON
func ToJSON(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// ToYAML converts a struct to YAML
func ToYAML(v interface{}) ([]byte, error) {
	b, err := yaml.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, nil
}
