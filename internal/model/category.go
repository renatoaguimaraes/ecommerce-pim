package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category model
type Category struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	ParentPath string             `json:"parentPath" bson:"parentPath"`
	Path       string             `json:"path" bson:"path"`
	Metadata   map[string]string  `json:"metadata,omitempty" bson:"metadata,omitempty"`
}
