package common

import "sync"

// Pair the struct of the data stored in levelDB
type Pair struct {
	sync.Map
}
