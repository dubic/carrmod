package models

import "time"

type TimeModel struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type AuditModel struct {
	CreatedBy string    `bson:"created_by"`
	Times     TimeModel `bson:"inline"`
}
