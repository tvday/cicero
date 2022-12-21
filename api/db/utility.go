package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CollectionFindByID(col *mongo.Collection, id interface{}, obj interface{}) error {
	switch v := id.(type) {
	case string:
		oID, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return err
		}
		result := col.FindOne(context.TODO(), primitive.M{"_id": oID})
		err = result.Decode(obj)
		return err
	case primitive.ObjectID:
		result := col.FindOne(context.TODO(), primitive.M{"_id": v})
		err := result.Decode(&obj)
		return err
	default:
		return errors.New("invalid type")
	}
}
