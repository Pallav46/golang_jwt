package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName    *string             `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName     *string             `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Username     *string             `bson:"username" json:"username" validate:"required,min=3,max=50"`
	Password     *string             `bson:"password" json:"password" validate:"required,min=8,max=100"`
	Email        *string             `bson:"email" json:"email" validate:"required,email"`
	Phone        *string             `bson:"phone,omitempty" json:"phone,omitempty" validate:"omitempty,e164"`
	Token        *string             `bson:"token,omitempty" json:"token,omitempty"`
	UserType     *string             `bson:"user_type" json:"user_type" validate:"required,oneof=USER ADMIN"`
	RefreshToken *string             `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	UserID       *string             `bson:"user_id" json:"user_id"`
}