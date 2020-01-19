package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Variant model
type Variant struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID       primitive.ObjectID `json:"-" bson:"productId"`
	Combination     Combination        `json:"combination" bson:"combination"`
	AttributesGroup AttributesGroup    `json:"attributes,omitempty" bson:"attributes,omitempty"`
	Compositions    Compositions       `json:"compositions,omitempty" bson:"compositions,omitempty"`
	Tags            []Tag              `json:"tags,omitempty" bson:"tags,omitempty"`
}
