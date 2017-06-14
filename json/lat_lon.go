package json

import (
	"encoding/json"
	"fmt"
)

// LatLon struct
type LatLon struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

// Print json
func (ll LatLon) Print() {
	json, _ := json.Marshal(ll)
	fmt.Println(string(json))
}

// Bytes - return json
func (ll LatLon) Bytes() []byte {
	json, _ := json.Marshal(ll)
	return json
}
