package endpoints

import (
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// key := r.Form.Get("key")
	// value, err := e.DB.Get(key)

	// shard := helpers.GetShard(key, 2)

	// if shard != e.ShardId {
	// 	helpers.Route(w, r, shard)

	// 	return
	// }

	// if err != nil {
	// 	log.Fatal("Error getting value")
	// }
}
