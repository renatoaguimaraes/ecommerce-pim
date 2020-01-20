package repo

import (
	"context"

	"github.com/google/logger"
	mongo "github.com/renatoaguimaraes/ecommerce-pim/internal/driver"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VariantRepo interface
type VariantRepo interface {
	Insert(variants []model.Variant) error
	FindVariantByID(productID string, variantID string) *model.Variant
	Update(p *model.Variant) error
	ListVariants(productID string) []model.Variant
}

// VariantRepoMongo repo mongo
type VariantRepoMongo struct{}

// Insert a set of variants
func (vr VariantRepoMongo) Insert(variants []model.Variant) error {
	vars := make([]interface{}, len(variants))

	for i, s := range variants {
		vars[i] = s
	}

	coll := mongo.GetMongoClient().Database("pim").Collection("variants")

	if _, err := coll.InsertMany(context.TODO(), vars); err != nil {
		logger.Errorf("Error running Insert(variants []model.Variant): %v", err)
		return err
	}

	return nil
}

// FindVariantByID variant
func (vr VariantRepoMongo) FindVariantByID(productID string, variantID string) *model.Variant {
	coll := mongo.GetMongoClient().Database("pim").Collection("variants")

	pid, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		logger.Errorf("Error running FindVariantByID(productID string, variantID string): %v", err)
		return nil
	}

	vid, err := primitive.ObjectIDFromHex(variantID)
	if err != nil {
		logger.Errorf("Error running FindVariantByID(productID string, variantID string): %v", err)
		return nil
	}

	v := new(model.Variant)
	filter := bson.M{"_id": vid, "productId": pid}

	if err := coll.FindOne(context.Background(), filter).Decode(v); err != nil {
		logger.Errorf("Error running FindVariantByID(productID string, variantID string): %v", err)
		return nil
	}

	return v
}

//Update a product
func (vr VariantRepoMongo) Update(v *model.Variant) error {

	coll := mongo.GetMongoClient().Database("pim").Collection("variants")

	filter := bson.M{"_id": v.ID, "productId": v.ProductID}
	update := bson.M{"$set": *v}

	if _, err := coll.UpdateOne(context.TODO(), filter, update); err != nil {
		logger.Errorf("Error running Update(p *model.Variant): %v", err)
		return err
	}

	return nil
}

// ListVariants variants
func (vr VariantRepoMongo) ListVariants(productID string) []model.Variant {
	var variants []model.Variant

	collection := mongo.GetMongoClient().Database("pim").Collection("variants")

	oid, err := primitive.ObjectIDFromHex(productID)

	if err != nil {
		logger.Errorf("Error running ListVariants(productID string): %v", err)
		return variants
	}

	cur, err := collection.Find(context.Background(), bson.M{"productId": oid})

	if err != nil {
		logger.Errorf("Error running ListVariants(productID string): %v", err)
		return variants
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		variant := model.Variant{}
		err := cur.Decode(&variant)
		if err != nil {
			logger.Errorf("Error running ListVariants(productID string): %variant", err)
		} else {
			variants = append(variants, variant)
		}
	}

	if err := cur.Err(); err != nil {
		logger.Errorf("Error running ListVariants(productID string): %v", err)
	}

	return variants
}
