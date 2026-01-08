package helpers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Route(w http.ResponseWriter, r *http.Request, shard int) {
	response, err := http.Get("http://" + e.Shards[shard].Host + ":" + strconv.Itoa(e.Shards[shard].Port) + r.RequestURI)

	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error at: " + err.Error())

		return
	}

	defer response.Body.Close()

	io.Copy(w, response.Body)
}
