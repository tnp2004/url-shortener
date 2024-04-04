package converter

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Url struct {
		Id          primitive.ObjectID `bson:"_id,omitempty"`
		ShortId     string             `bson:"short_id"`
		Destination string             `bson:"destination"`
		CreatedAt   time.Time          `bson:"created_at"`
	}
)
