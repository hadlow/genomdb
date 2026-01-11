package endpoints

import (
	"net/http"

	"github.com/hadlow/genomdb/internal/database"
	"github.com/hadlow/genomdb/types"
)

type ServerInterface interface {
	GetDatabase() *database.Database
	GetConfig() *types.Config
}

func Get(s ServerInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		// Now you have access to s.GetDatabase(), s.GetConfig(), etc.
		// key := r.Form.Get("key")
		// value, err := s.GetDatabase().Get(key)

		// shard := helpers.GetShard(key, 2)

		// if shard != s.GetConfig().ShardId {
		// 	helpers.Route(w, r, shard)
		// 	return
		// }

		// if err != nil {
		// 	log.Fatal("Error getting value")
		// }
	}
}
