package repository

import (
	"context"
	"errors"

	"github.com/anvi23mth/inventory-system/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 1. THE INTERFACE: This is the "Contract" that the Service uses.
// By defining this, your Service can accept either the REAL Mongo repo or a MOCK repo.
type ProductRepository interface {
	Create(ctx context.Context, p model.Product) error
	GetAll(ctx context.Context) ([]model.Product, error)
	GetByID(ctx context.Context, id string) (model.Product, error)
	Update(ctx context.Context, id string, p model.Product) error
	Delete(ctx context.Context, id string) error
}

// 2. THE IMPLEMENTATION: This struct handles the actual MongoDB logic.
// We keep it private (lowercase 'm') so other packages use the interface instead.
type mongoProductRepository struct {
	Col *mongo.Collection
}

// 3. THE CONSTRUCTOR: Returns the Interface type, not the struct pointer.
func NewProductRepository(db *mongo.Database) ProductRepository {
	return &mongoProductRepository{
		Col: db.Collection("products"),
	}
}

// --- Implementation of the Interface Methods ---

func (r *mongoProductRepository) Create(ctx context.Context, p model.Product) error {
	_, err := r.Col.InsertOne(ctx, p)
	return err
}

func (r *mongoProductRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	cursor, err := r.Col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []model.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *mongoProductRepository) GetByID(ctx context.Context, id string) (model.Product, error) {
	var p model.Product
	err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&p)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Product{}, errors.New("product not found")
		}
		return model.Product{}, err
	}
	return p, nil
}

func (r *mongoProductRepository) Update(ctx context.Context, id string, p model.Product) error {
	_, err := r.Col.ReplaceOne(ctx, bson.M{"_id": id}, p)
	return err
}

func (r *mongoProductRepository) Delete(ctx context.Context, id string) error {
	_, err := r.Col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
