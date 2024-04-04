package converterRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/tnp2004/url-shortener/modules/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IConverterRepository interface {
		InsertUrl(pctx context.Context, req *converter.Url) (primitive.ObjectID, error)
		SearchShortIdByObjectId(pctx context.Context, id primitive.ObjectID) (string, error)
	}

	converterRepository struct {
		db *mongo.Client
	}
)

func NewConverterRepository(db *mongo.Client) IConverterRepository {
	return &converterRepository{db: db}
}

func (r *converterRepository) InsertUrl(pctx context.Context, req *converter.Url) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.db.Database("converter_db")
	col := db.Collection("url")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertUrl failed: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert url failed")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *converterRepository) SearchShortIdByObjectId(pctx context.Context, id primitive.ObjectID) (string, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.db.Database("converter_db")
	col := db.Collection("url")

	result := new(converter.Url)

	if err := col.FindOne(ctx, bson.M{"_id": id}).Decode(result); err != nil {
		log.Printf("Error: SearchShortIdByObjectId failed: %s", err.Error())
		return "", errors.New("error: search short id failed")
	}

	return result.ShortId, nil
}
