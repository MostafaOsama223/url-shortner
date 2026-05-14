package repo

import (
	"fmt"
	"shortner-service/database"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UrlRepo struct {
	collection *database.MongoDBCollection
}

type UrlRecord struct {
	OriginalUrl string `bson:"original_url"`
	ShortUrl    string `bson:"short_url"`
}

func (r *UrlRepo) InsertUrlRecord(record *UrlRecord) (UrlRecord, error) {
	_, err := r.collection.InsertOne(record)
	if err != nil {
		return UrlRecord{}, err
	}

	return *record, nil
}

func (r *UrlRepo) FindOriginalUrl(shortUrl string) (string, error) {
	filter := bson.D{{"short_url", shortUrl}}

	result := r.collection.FindOne(filter)
	if result == nil {
		return "", fmt.Errorf("no document found for short URL: %s", shortUrl)
	}

	return result["original_url"].(string), nil
}

func NewUrlRepo(collection *database.MongoDBCollection) *UrlRepo {
	return &UrlRepo{collection: collection}
}
