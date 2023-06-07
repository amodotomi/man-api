package db

import (
	"context"
	"errors"
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

func (d *db) FindAll(ctx context.Context) (u []user.User, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	
	if cursor.Err() != nil {
		return u, fmt.Errorf("FAILED: failed to find all users due to error: %v", err)
	}

	if err = cursor.All(ctx, &u); err != nil {
		return u, fmt.Errorf("FAILED: failed to read all documents form cursor | error %v", err)
	}


	return u, nil

}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id) // oid === objectID
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectID | hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		// TODO: handle error
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			// TODO: handle error 404
			return u, fmt.Errorf("FAILED: NOT FOUND 404")
		}

		return u, fmt.Errorf("FAILED: failed to find one user by id: #{id} due ro error: %v", err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("FAILED: failed to decode user (id: %s) from DB due to error: %v", id, err)
	}

	return u, nil

}

func (d *db) Update(ctx context.Context, user user.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("FAILED: failed to convert user ID to object ID | ID=%s", user.ID)
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("FAILED: failed to marshal user due to: error: %v", err)
	}

	var updateUserObj bson.M
	bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("FAILED: failed to unmarshal user due to: error: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"set": updateUserObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("FAILED: failed to execute update query | error %v", err)
	}

	if result.MatchedCount == 0 {
		// todo: err not found 404
		return fmt.Errorf("FAILED: NOT FOUND 404")
	}

	d.logger.Tracef("OK: MATCHED %d documents and MODIFIED %d", result.MatchedCount, result.ModifiedCount)
	
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("FAILED: failed to convert user ID to object ID | ID=%s", id)
	}

	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("FAILED: failed to delete | error: %s", err)
	}
	if result.DeletedCount == 0 {
		// TODO 404
		return fmt.Errorf("FAILED: NOT FOUND 404")
	}	
	d.logger.Tracef("SUCCESS: Deleted %d documents", result.DeletedCount)

	return nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger: 	logger,
	}
}