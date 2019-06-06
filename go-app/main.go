package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/favclip/ucon"
	"github.com/favclip/ucon/swagger"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

var DS datastore.Client

func main() {
	Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	ucon.ListenAndServe(fmt.Sprintf(":%s", port))
}

func Initialize() {
	ctx := context.Background()

	SetUpDatastore(ctx)
	SetUpUcon()
}

func SetUpDatastore(ctx context.Context) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	ds, err := clouddatastore.FromContext(ctx, datastore.WithProjectID(projectID))
	if err != nil {
		log.Fatal("required ENV GOOGLE_CLOUD_PROJECT:", err)
		os.Exit(1)
	}
	DS = ds
}

func SetUpUcon() {
	ucon.Orthodox()
	ucon.Middleware(swagger.RequestValidator())

	swPlugin := swagger.NewPlugin(&swagger.Options{
		Object: &swagger.Object{
			Info: &swagger.Info{
				Title:   "GCPUG",
				Version: "v4",
			},
			Schemes: []string{"http", "https"},
		},
		DefinitionNameModifier: func(refT reflect.Type, defName string) string {
			if strings.HasSuffix(defName, "JSON") {
				return defName[:len(defName)-4]
			}
			return defName
		},
	})
	ucon.Plugin(swPlugin)

	SetupOrganizationAPI(swPlugin)

	ucon.DefaultMux.Prepare()
	http.Handle("/api/", ucon.DefaultMux)
	http.HandleFunc("/", StaticContentsHandler)
}
