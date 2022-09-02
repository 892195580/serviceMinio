package main

import (
	"minio/api"
	"testing"
)

func TestDownLoad(t *testing.T) {
	//Get client
	minioclient := api.GetClient()
	//Get BucketList
	api.GetBucketList(minioclient)

	//Get objects from bucketname
	api.ListObjects(minioclient, "test", "")
	//Download object to path.
	api.GetObjectFile(minioclient, "test", "nginx.yaml", "nginx.yaml")
}
