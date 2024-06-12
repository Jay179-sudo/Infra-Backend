package data

import "time"

type Specification struct {
	RAM        int32     `json:"RAM"`
	Storage    int32     `json:"Storage"`
	ExpiryTime time.Time `json:"ExpiryTime"`
	PublicKey  string    `json:"PublicKey"`
}
type VMRequest struct {
	Email string        `json:"Email"`
	Spec  Specification `json:"Specification"`
}
