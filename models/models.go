package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Report struct {
	ID              primitive.ObjectID  `bson:"_id, omitempty"`
	PublicKey       string              `bson:"publickey"`
	TransactionHash string              `bson:"transactionhash"`
	Description     string              `bson:"description"`
	Time            primitive.Timestamp `bson:"time"`
}
