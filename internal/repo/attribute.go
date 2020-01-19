package repo

import (
	"context"

	"github.com/google/logger"
	"github.com/renatoaguimaraes/ecm-pim/internal/driver"
	"github.com/renatoaguimaraes/ecm-pim/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AttributeRepo interface
type AttributeRepo interface {
	Insert(attr *model.Attribute) error
	FindAttributeByID(group string, attributeID string) *model.Attribute
}

// AttributeValueRepo interface
type AttributeValueRepo interface {
	Insert(attr *model.AttributeValue) error
	FindAttributeValueByID(attributeID string, attributeValueID string) *model.AttributeValue
}

// AttributeRepoMongo repo mongo
type AttributeRepoMongo struct{}

// AttributeValueRepoMongo repo mongo
type AttributeValueRepoMongo struct{}

//Insert a product
func (ar AttributeRepoMongo) Insert(attr *model.Attribute) error {
	coll := driver.GetMongoClient().Database("pim").Collection("attributes")

	res, err := coll.InsertOne(context.TODO(), *attr)

	if err != nil {
		logger.Errorf("Error running Insert(variants []model.Variant): %v", err)
		return err
	}

	attr.ID = res.InsertedID.(primitive.ObjectID)

	return nil
}

// FindAttributeByID get an attribute by id
func (ar AttributeRepoMongo) FindAttributeByID(group string, attributeID string) *model.Attribute {
	coll := driver.GetMongoClient().Database("pim").Collection("attributes")

	oid, err := primitive.ObjectIDFromHex(attributeID)

	if err != nil {
		logger.Errorf("Error running FindAttributeByID(group string, attributeID string): %v", err)
		return nil
	}

	attr := new(model.Attribute)

	if err := coll.FindOne(context.Background(), bson.M{"_id": oid, "group": group}).Decode(attr); err != nil {
		logger.Errorf("Error running FindAttributeByID(group string, attributeID string): %v", err)
		return nil
	}

	return attr
}

//Insert a product
func (ar AttributeValueRepoMongo) Insert(attr *model.AttributeValue) error {
	coll := driver.GetMongoClient().Database("pim").Collection("attributes")

	res, err := coll.InsertOne(context.TODO(), *attr)

	if err != nil {
		logger.Errorf("Error running Insert(attr *model.AttributeValue): %v", err)
		return err
	}

	attr.ID = res.InsertedID.(primitive.ObjectID)

	return nil
}

// FindAttributeValueByID get an attribute values by id
func (ar AttributeValueRepoMongo) FindAttributeValueByID(attributeID string, attributeValueID string) *model.AttributeValue {
	coll := driver.GetMongoClient().Database("pim").Collection("attributes-values")

	attributeOID, err := primitive.ObjectIDFromHex(attributeID)
	attributeValueOID, err := primitive.ObjectIDFromHex(attributeValueID)

	if err != nil {
		logger.Errorf("Error running FindAttributeValueByID(attributeID string, attributeValueID string): %v", err)
		return nil
	}

	attributeValue := new(model.AttributeValue)

	if err := coll.FindOne(context.Background(), bson.M{"_id": attributeValueOID, "attributeId": attributeOID}).Decode(attributeValue); err != nil {
		logger.Errorf("Error running FindAttributeValueByID(attributeID string, attributeValueID string): %v", err)
		return nil
	}

	return attributeValue
}
