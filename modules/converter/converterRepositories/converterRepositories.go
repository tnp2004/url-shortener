package converterRepositories

import "go.mongodb.org/mongo-driver/mongo"

type (
	IConverterRepository interface{}

	converterRepository struct {
		db *mongo.Client
	}
)

func NewConverterRepository(db *mongo.Client) IConverterRepository {
	return &converterRepository{db: db}
}
