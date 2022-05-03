package main

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	ctx   = context.Background()
	creds = flag.String("creds", "", "")
)

func main() {
	// jwt, err := google.JWTConfigFromJSON([]byte(*creds))
	// if err != nil {
	// 	log.Fatalf("unable to read credentials: %s", err)
	// }

	// _ = jwt

	gcs, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(*creds)))
	if err != nil {
		log.Fatalf("unable to init gcs client: %s", err)
	}

	bucket := gcs.Bucket("bantha-pipelines")

	it := bucket.Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			log.Fatalf("Bucket(%q).Objects: %v", bucket, err)
		}
		log.Println(attrs.Name)
	}
}
