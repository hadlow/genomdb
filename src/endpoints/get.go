package endpoints

import (
	"fmt"
	"log"
	"net/http"
)

func (e *Endpoint) Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	key := r.Form.Get("key")
	value, err := e.DB.Get(key)

	shard := e.getShard(key)

	if shard != e.ShardId {
		e.Route(w, r, shard)

		return
	}

	if err != nil {
		log.Fatal("Error getting value")
	}

	fmt.Println(value)
}
