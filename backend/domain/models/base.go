package models

import "time"

type BaseModel struct {
	CreatedBy string `bson:"created_by"`
	Created   time.Time
	Updated   time.Time
}
