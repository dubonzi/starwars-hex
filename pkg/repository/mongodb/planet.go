package mongodb

import (
	"context"
	"errors"
	"starwars-hex/pkg/errs"
	"starwars-hex/pkg/planets"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type planetDB struct {
	col *mongo.Collection
}

// NewPlanetDB creates a new accessor for the Planet collection.
func NewPlanetDB(col *mongo.Collection) planets.Repository {
	return planetDB{col}
}

func (p planetDB) List() ([]planets.Planet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cursor, err := p.col.Find(ctx, bson.M{}, &options.FindOptions{
		Sort: bson.M{"name": 1},
	})
	if err != nil {
		return nil, err
	}

	pls := make([]planets.Planet, 0, 5)

	err = cursor.All(ctx, &pls)
	if err != nil {
		return nil, err
	}

	cursor.Close(ctx)
	return pls, nil
}

func (p planetDB) Exists(name string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result := p.col.FindOne(ctx, bson.M{"name": name})
	if err := result.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p planetDB) Insert(planet planets.Planet) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := p.col.InsertOne(ctx, planet)
	return err

}

func (p planetDB) FindByName(name string) (planets.Planet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	cursor := p.col.FindOne(ctx, bson.M{"name": name})
	if err := cursor.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return planets.Planet{}, errs.NoDBResults
		}
		return planets.Planet{}, err
	}
	var planet planets.Planet
	return planet, cursor.Decode(&planet)
}

func (p planetDB) Delete(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := p.col.FindOneAndDelete(ctx, bson.M{"name": name}).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return errs.NoDBResults
	}
	return err
}
