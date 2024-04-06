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
		FindOneDestination(pctx context.Context, url string) (*converter.Url, error)
		FindOneDestinationByShortId(pctx context.Context, id string) (string, error)
		UpdateOneExpiration(pctx context.Context, id primitive.ObjectID, expiration time.Time) error
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

func (r *converterRepository) FindOneDestinationByShortId(pctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.db.Database("converter_db")
	col := db.Collection("url")

	result := new(converter.Url)

	if err := col.FindOne(ctx, bson.M{"short_id": id}).Decode(result); err != nil {
		log.Printf("Error: FindOneDestinationByShortId failed: %s", err.Error())
		return "", errors.New("error: short id not found")
	}

	return result.Destination, nil
}

func (r *converterRepository) FindOneDestination(pctx context.Context, url string) (*converter.Url, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.db.Database("converter_db")
	col := db.Collection("url")

	result := new(converter.Url)

	if err := col.FindOne(ctx, bson.M{"destination": url}).Decode(result); err != nil {
		log.Printf("Error: FindOneDestination failed: %s", err.Error())
		return nil, errors.New("error: search destination not found")
	}

	return result, nil
}

func (r *converterRepository) UpdateOneExpiration(pctx context.Context, id primitive.ObjectID, expiration time.Time) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.db.Database("converter_db")
	col := db.Collection("url")

	if _, err := col.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"expires_at": expiration}}); err != nil {
		log.Printf("Error: UpdateOneExpiration failed: %s", err.Error())
		return errors.New("error: update expiration failed")
	}

	return nil
}
