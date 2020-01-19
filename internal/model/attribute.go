package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Attribute model
type Attribute struct {
	ID             primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name           string               `json:"name" bson:"name"`
	Group          string               `json:"-" bson:"group"`
	AllowVariation bool                 `json:"allowVariation" bson:"allowVariation"`
	Categories     []primitive.ObjectID `json:"categories" bson:"categories"`
	Metadata       map[string]string    `json:"metadata" bson:"metadata"`
}

// AttributeValue model
type AttributeValue struct {
	AttributeID primitive.ObjectID           `json:"-" bson:"attributeId"`
	ID          primitive.ObjectID           `json:"id" bson:"_id,omitempty"`
	Value       string                       `json:"value" bson:"value"`
	DataType    string                       `json:"dataType" bson:"dataType"`
	ParentID    *primitive.ObjectID          `json:"parentId" bson:"parentId"`
	Composition []primitive.ObjectID         `json:"composition" bson:"composition"`
	Metadata    map[string]map[string]string `json:"metadata" bson:"metadata"`
}

//SetAttributeID convert string into ObjectID
func (attributeValue *AttributeValue) SetAttributeID(atrributeID string) error {
	oid, err := primitive.ObjectIDFromHex(atrributeID)
	if err != nil {
		return err
	}
	attributeValue.AttributeID = oid
	return nil
}
