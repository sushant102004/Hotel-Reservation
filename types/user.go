package types

type User struct {
	/*
		Due to bson ID will be stored as _id in MongoDB
		omitempty will not show ID in JSON Response.
	*/
	ID        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
