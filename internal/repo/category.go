package repo

import (
	"context"

	"github.com/google/logger"
	mongo "github.com/renatoaguimaraes/ecommerce-pim/internal/driver"
	"github.com/renatoaguimaraes/ecommerce-pim/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CategoryRepo interface
type CategoryRepo interface {
	Insert(p *model.Category) error
	Update(p *model.Category) error
	FindByID(productID string) *model.Category
	ListAll() []model.Category
	List(path string) []model.Category
}

// CategoryRepoMongo repo mongo
type CategoryRepoMongo struct{}

// Insert save
func (cr CategoryRepoMongo) Insert(p *model.Category) error {
	coll := mongo.GetMongoClient().Database("pim").Collection("categories")
	res, err := coll.InsertOne(context.TODO(), *p)
	if err != nil {
		logger.Errorf("Error running Insert(p *model.Category): %v", err)
		return err
	}
	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		p.ID = id
	}
	return nil
}

// FindByID find by id
func (cr CategoryRepoMongo) FindByID(ID string) *model.Category {
	coll := mongo.GetMongoClient().Database("pim").Collection("categories")
	pid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		logger.Errorf("Error running FindByID(ID string): %v", err)
		return nil
	}
	Category := new(model.Category)
	filter := bson.M{"_id": pid}
	if err = coll.FindOne(context.Background(), filter).Decode(Category); err != nil {
		return nil
	}
	return Category
}

// FindByID find by id
func (cr CategoryRepoMongo) List(path string) []model.Category {
	var categories []model.Category
	collection := mongo.GetMongoClient().Database("pim").Collection("categories")
	filter := bson.D{{"path", primitive.Regex{Pattern: path, Options: ""}}}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logger.Errorf("Error running List(): %v", err)
		return categories
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		category := model.Category{}
		err := cur.Decode(&category)
		if err != nil {
			logger.Errorf("Error running List(): %v", err)
		} else {
			categories = append(categories, category)
		}
	}
	if err := cur.Err(); err != nil {
		logger.Errorf("Error running List(): %v", err)
	}
	return categories
}

//Update a Category
func (cr CategoryRepoMongo) Update(p *model.Category) error {
	coll := mongo.GetMongoClient().Database("pim").Collection("categories")
	if _, err := coll.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": p.ID}}, bson.M{"$set": *p}); err != nil {
		logger.Errorf("Error running Update(p *model.Category): %v", err)
		return err
	}
	return nil
}

// ListVariants variants
func (cr CategoryRepoMongo) ListAll() []model.Category {
	var categories []model.Category
	collection := mongo.GetMongoClient().Database("pim").Collection("categories")
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		logger.Errorf("Error running List(): %v", err)
		return categories
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		category := model.Category{}
		err := cur.Decode(&category)
		if err != nil {
			logger.Errorf("Error running List(): %v", err)
		} else {
			categories = append(categories, category)
		}
	}
	if err := cur.Err(); err != nil {
		logger.Errorf("Error running List(): %v", err)
	}
	return categories
}
