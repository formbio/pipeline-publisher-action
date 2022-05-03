package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
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

	fmt.Println(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	// gcs, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(*creds)))
	gcs, err := storage.NewClient(ctx)
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
