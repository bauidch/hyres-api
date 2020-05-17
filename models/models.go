package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Series struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Status    string             `bson:"status" json:"status"`
	StartDate time.Time          `bson:"start_at,omitempty" json:"start_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	EndDate   time.Time          `bson:"ended_at,omitempty" json:"ended_at,omitempty"`
}

type Seed struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SeriesID    primitive.ObjectID `json:"series_id,omitempty" bson:"series_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Type        string             `json:"type" bson:"type"`
	ResearchTag string             `json:"research_tag" bson:"research_tag"`
	CreateDate  time.Time          `bson:"created_at" json:"created_at"`
}

type Plant struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SeriesID    primitive.ObjectID `json:"series_id,omitempty" bson:"series_id,omitempty"`
	SeedID      primitive.ObjectID `json:"seed_id,omitempty" bson:"seed_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Type        string             `json:"type" bson:"type"`
	ResearchTag string             `json:"research_tag" bson:"research_tag"`
	SetAt       time.Time          `bson:"set_at" json:"set_at"`
}

type SeedJournal struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SeriesID    primitive.ObjectID `json:"series_id,omitempty" bson:"series_id,omitempty"`
	SeedID      primitive.ObjectID `json:"seed_id,omitempty" bson:"seed_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Type        string             `json:"type" bson:"type"`
	ResearchTag string             `json:"research_tag" bson:"research_tag"`
	CreateDate  time.Time          `bson:"created_at" json:"created_at"`
	Text        string             `json:"text" bson:"text"`
	Photo       string             `json:"photo" bson:"photo"`
}

type PlantJournal struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SeriesID    primitive.ObjectID `json:"series_id,omitempty" bson:"series_id,omitempty"`
	PlantID      primitive.ObjectID `json:"plant_id,omitempty" bson:"plant_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Type        string             `json:"type" bson:"type"`
	ResearchTag string             `json:"research_tag" bson:"research_tag"`
	CreateDate  time.Time          `bson:"created_at" json:"created_at"`
	Text        string             `json:"text" bson:"text"`
	Photo       string             `json:"photo" bson:"photo"`
}
