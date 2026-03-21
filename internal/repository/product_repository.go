package repository

import (
	"context"
	"errors"

	"github.com/anvi23mth/inventory-system/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Col *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{Col: db.Collection("products")}
}

func (r *ProductRepository) Create(ctx context.Context, p model.Product) error {
	_, err := r.Col.InsertOne(ctx, p)
	return err
}

func (r *ProductRepository) GetAll(ctx context.Context) ([]model.Product, error) {
	cursor, err := r.Col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var products []model.Product
	err = cursor.All(ctx, &products)
	return products, err
}

//	func (r *ProductRepository) GetByID(ctx context.Context, id string) (model.Product, error) {
//		var p model.Product
//		err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&p)
//		return p, err
//	}
func (r *ProductRepository) GetByID(ctx context.Context, id string) (model.Product, error) {
	var p model.Product
	err := r.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&p)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Product{}, errors.New("product not found")
		}
		return model.Product{}, err
	}

	return p, nil
}
func (r *ProductRepository) Update(ctx context.Context, id string, p model.Product) error {
	_, err := r.Col.ReplaceOne(ctx, bson.M{"_id": id}, p)
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	_, err := r.Col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
