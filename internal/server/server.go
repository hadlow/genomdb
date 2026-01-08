package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/hadlow/genomdb/internal/database"
	"github.com/hadlow/genomdb/internal/endpoints"
	"github.com/hadlow/genomdb/types"
)

func pingNodeSuccessful(shard types.Shard) bool {
	fmt.Println("Trying node: ", shard.Port)
	response, err := http.Get("http://" + shard.Host + ":" + strconv.Itoa(shard.Port))

	if err != nil {
		return false
	}

	defer response.Body.Close()

	return true
}

func findNextAvailableNode(shards []types.Shard) types.Shard {
	for _, shard := range shards {
		if !pingNodeSuccessful(shard) {
			return shard
		}
	}

	return types.Shard{}
}

type Server struct {
	database *database.Database
	close    func() error
	config   *types.Config
}

func NewServer(config *types.Config) (s *Server) {
	database, close, err := database.NewDatabase(config.Database)
	database.SetBucket("main")

	if err != nil {
		log.Fatal(err)
	}

	return &Server{database: database, close: close, config: config}
}

func (s *Server) Serve() error {
	shard := findNextAvailableNode(s.config.Shards)

	http.HandleFunc("/ping", endpoints.Ping)
	http.HandleFunc("/get", endpoints.Get)
	http.HandleFunc("/set", endpoints.Set)

	err := http.ListenAndServe(shard.Host+":"+strconv.Itoa(shard.Port), nil)

	if err != nil {
		return err
	}

	fmt.Println("Starting server on: " + shard.Host + ":" + strconv.Itoa(shard.Port))

	return nil
}

func (s *Server) Close() {
	s.close()
}
