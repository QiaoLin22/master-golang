package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CatFact struct {
	Fact   string `bson:"fact"`
	Length int    `bson:"length"`
}

type Storer interface {
	GetAll() ([]*CatFact, error)
	Put(*CatFact) error
}

type MongoStore struct {
	client     *mongo.Client
	database   string
	collection string
	coll       *mongo.Collection
}

func NewMongoStore() (*MongoStore, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	coll := client.Database("catfact").Collection("facts ")

	return &MongoStore{
		client:     client,
		database:   "catfact",
		collection: "facts",
		coll:       coll,
	}, nil
}

func (store *MongoStore) Put(fact *CatFact) error {
	_, err := store.coll.InsertOne(context.TODO(), fact)
	return err
}

func (store *MongoStore) GetAll() ([]*CatFact, error) {
	query := bson.M{}
	cursor, err := store.coll.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}

	results := []*CatFact{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

type Server struct {
	store Storer
}

func NewServer(s Storer) *Server {
	return &Server{
		store: s,
	}
}

func (s *Server) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	catFacts, err := s.store.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catFacts)
}

type CatFactWorker struct {
	store       Storer
	apiEndPoint string
}

func NewCatFactWorkert(store Storer, apiEndPoint string) *CatFactWorker {
	return &CatFactWorker{
		store:       store,
		apiEndPoint: apiEndPoint,
	}
}

func (cfw *CatFactWorker) start() error {
	ticker := time.NewTicker(2 * time.Second)
	for {
		resp, err := http.Get(cfw.apiEndPoint)
		if err != nil {
			return err
		}
		var catFact CatFact
		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}

		if err := cfw.store.Put(&catFact); err != nil {
			return err
		}
		<-ticker.C
	}
}

func main() {
	mongoStore, err := NewMongoStore()
	if err != nil {
		log.Fatal(err)
	}
	worker := NewCatFactWorkert(mongoStore, "https://catfact.ninja/fact")
	go worker.start()

	server := NewServer(mongoStore)
	http.HandleFunc("/facts", server.handleGetAllFacts)
	http.ListenAndServe(":3000", nil)
}
