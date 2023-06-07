package db

import (
	"context"
	"fmt"
	"proj/internal/user"
	"proj/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("---> creating user...")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("FAILED: creating user failed due to: %v", err)
	}	

	d.logger.Debug("---> converting inserted ID to object ID...")
	oid, ok := result.InsertedID.(primitive.ObjectID) // oid === objectID
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("FAILED: failed to convert object ID to hex | probably oid: %s", result)
}


func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id) // oid === objectID
	if err != nil {
		return u, fmt.Errorf("failed to convert hex tp objectID | hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		// TODO: 404
		return u, fmt.Errorf("FAILED: failed to find user by id: %s due to error: %v", id, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("FAILED: failed to decode user (id: %s) from DB due to error: %v", id, err)
	}

	return u, nil

}

// -----------------------------------------------------------------------------------------------------

func (d *db) Update(ctx context.Context, user user.User) error {
	d.collection.UpdateOne()
}

// -----------------------------------------------------------------------------------------------------

func (d *db) Delete(ctx context.Context, id string) error {
	d.collection.DeleteOne()
}

// -----------------------------------------------------------------------------------------------------

func NewStorage(collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger: logger,
	}
}