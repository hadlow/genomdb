package endpoints

import (
	"fmt"
	"net/http"
)

func Set(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// key := r.Form.Get("key")
	// value := r.Form.Get("data")

	// shard := helpers.GetShard(key, 2)

	// err := e.DB.Set(key, []byte(value))

	// if shard != e.ShardId {
	// 	helpers.Route(w, r, shard)

	// 	return
	// }

	// if err != nil {
	// 	log.Fatal("Error setting value")
	// }

	fmt.Println("Key value set")
}
