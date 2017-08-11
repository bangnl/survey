package model

//Pair :
type Pair struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func newPair(key string, v interface{}) *Pair {
	return &Pair{Key: key, Value: v}
}
