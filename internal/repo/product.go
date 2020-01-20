package repo

import (
	"context"
	"github.com/google/logger"
	mongo "github.com/renatoaguimaraes/ecommerce-pim/internal/driver"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductRepo interface
type ProductRepo interface {
	Insert(p *model.Product) error
	Update(p *model.Product) error
	FindProductByID(productID string) *model.Product
	List(filter map[string]string) []model.Product
}

// ProductRepoMongo repo mongo
type ProductRepoMongo struct{}

// Insert save
func (pr ProductRepoMongo) Insert(p *model.Product) error {
	coll := mongo.GetMongoClient().Database("pim").Collection("products")
	p.Status = model.Draft
	res, err := coll.InsertOne(context.TODO(), *p)
	if err != nil {
		logger.Errorf("Error running Insert(p *model.Product): %v", err)
		return err
	}
	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		p.ID = id
	}
	return nil
}

// FindProductByID find a produc by id
func (pr ProductRepoMongo) FindProductByID(productID string) *model.Product {
	coll := mongo.GetMongoClient().Database("pim").Collection("products")
	pid, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		logger.Errorf("Error running FindProductByID(productID string): %v", err)
		return nil
	}
	product := new(model.Product)
	filter := bson.M{"_id": pid}
	if err = coll.FindOne(context.Background(), filter).Decode(product); err != nil {
		return nil
	}
	return product
}

//Update a product
func (pr ProductRepoMongo) Update(p *model.Product) error {
	coll := mongo.GetMongoClient().Database("pim").Collection("products")
	if _, err := coll.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": p.ID}}, bson.M{"$set": *p}); err != nil {
		logger.Errorf("Error running Update(p *model.Product): %v", err)
		return err
	}
	return nil
}

// FindByID find by id
func (pr ProductRepoMongo) List(filter map[string]string) []model.Product {
	var products []model.Product
	collection := mongo.GetMongoClient().Database("pim").Collection("products")
	// filters
	var params []bson.E
	for k, v := range filter {
		params = append(params, bson.E{k, v})
	}
	// query
	cur, err := collection.Find(context.Background(), params)
	if err != nil {
		logger.Errorf("Error running List(): %v", err)
		return products
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		p := model.Product{}
		err := cur.Decode(&p)
		if err != nil {
			logger.Errorf("Error running List(): %v", err)
		} else {
			products = append(products, p)
		}
	}
	if err := cur.Err(); err != nil {
		logger.Errorf("Error running List(): %v", err)
	}
	return products
}
