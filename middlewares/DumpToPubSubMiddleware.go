package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Good-Will/google-cloud/util"
	"github.com/Good-Will/middlewares"
)

// NewDumpToPubSubMiddleware TODO
func NewDumpToPubSubMiddleware() func(next http.Handler) http.Handler {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID != "" {
		return middlewares.NewDumpMiddleware(func(dump *middlewares.RoundtripDump) {
			marshalledDump, _ := json.Marshal(dump)

			g := util.GooglePubSub{}

			err := g.InProject(projectID).Topic("GATEWAY_ROUNDTRIPS").Pub(marshalledDump)
			if err != nil {
				log.Println(err)
			}
		})
	} else {
		return middlewares.NewDumpToLogMiddleware()
	}
}
