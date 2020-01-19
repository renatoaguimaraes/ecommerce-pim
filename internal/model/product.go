package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tag model
type Tag string

// AttributeValueRef model
type AttributeValueRef string

// AttributeRef model
type AttributeRef string

// Attributes model
type Attributes map[AttributeRef]AttributeValueRef

// Amount model
type Amount string

// Material model
type Material string

// Materials model
type Materials map[Material]Amount

// Part model
type Part string

// Compositions model
type Compositions map[Part]Materials

// Group model
type Group string

// AttributesGroup model
type AttributesGroup map[Group]Attributes

// Variation model
type Variation map[AttributeRef][]AttributeValueRef

// Combination model
type Combination map[AttributeRef]AttributeValueRef

//Configuration model
type Configuration struct {
	Plp       bool `json:"plp" bson:"plp"`
	Pdp       bool `json:"pdp" bson:"pdp"`
	AddToCart bool `json:"addToCart" bson:"addToCart"`
}

// Status product status model
type Status string

const (
	// Draft product status
	Draft Status = "Draft"
	// Ready product status
	Ready Status = "Ready"
	// Published product status
	Published Status = "Published"
	// Deactivated product status
	Deactivated Status = "Deactivated"
)

// Product model
type Product struct {
	ID            primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	Schema        *primitive.ObjectID `json:"schema" bson:"schema"`
	Name          string              `json:"name" bson:"name"`
	Type          string              `json:"type" bson:"type"`
	Status        Status              `json:"status" bson:"status"`
	Categories    []string            `json:"categories" bson:"categories"`
	Attributes    AttributesGroup     `json:"attributes" bson:"attributes"`
	Compositions  Compositions        `json:"compositions" bson:"compositions,omitempty"`
	Variation     Variation           `json:"variation" bson:"variation"`
	Tags          []Tag               `json:"tags" bson:"tags,omitempty"`
	Configuration Configuration       `json:"configuration" bson:"configuration"`
}
