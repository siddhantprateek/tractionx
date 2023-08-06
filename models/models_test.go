package models_test

import (
	"encoding/json"
	"testing"

	models "tractionx/models"

	"github.com/stretchr/testify/assert"
)

func TestPropertyJSONSerialization(t *testing.T) {
	property := models.Property{
		ID:           "prop123",
		PlotNumber:   1,
		Builder:      "ABC Builders",
		OwnedBy:      "John Doe",
		CurrentPrice: 100000,
		SellingPrice: 120000,
	}

	propertyJSON, err := json.Marshal(property)
	assert.NoError(t, err, "JSON marshaling should not return an error")

	var deserializedProperty models.Property
	err = json.Unmarshal(propertyJSON, &deserializedProperty)
	assert.NoError(t, err, "JSON unmarshaling should not return an error")

	assert.Equal(t, property, deserializedProperty, "Deserialized Property should be equal to the original Property")
}
