package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/cswords/gcloud/util"
	"github.com/cswords/middlewares"
)

// NewDumpToPubSubMiddleware creates a new DumpMiddleware object which publish the data records including request and response of all HTTP roundtrips to a PubSub topic in GCP.
// The project id is fetched from environment variable GOOGLE_CLOUD_PROJECT which should be automatically set by GAE.
// If the project id is not found, this function returns a DumpToLogMiddleware.
func NewDumpToPubSubMiddleware(topicID string) func(next http.Handler) http.Handler {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID != "" {
		return middlewares.NewDumpMiddleware(func(dump *middlewares.RoundtripDump) {
			marshalledDump, _ := json.Marshal(dump)

			g := util.GooglePubSub{}

			err := g.InProject(projectID).Topic(topicID).Pub(marshalledDump)
			if err != nil {
				log.Println(err)
			}
		})
	}
	return middlewares.NewDumpToLogMiddleware()
}
