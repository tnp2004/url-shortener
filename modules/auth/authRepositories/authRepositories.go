package authRepositories

import "go.mongodb.org/mongo-driver/mongo"

type (
	IAuthRepository interface{}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) IAuthRepository {
	return &authRepository{db: db}
}
