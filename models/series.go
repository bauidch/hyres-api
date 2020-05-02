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
