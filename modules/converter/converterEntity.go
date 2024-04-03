package converter

type (
	ShortUrl struct {
		shortenedUrl string `bson:"shortend_url"`
		endpoint     string `bson:"endpoint"`
		createdAt    string `bson:"created_at"`
	}
)
