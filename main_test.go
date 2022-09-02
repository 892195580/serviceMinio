package main

import (
	"context"
	"minio/api"
	"testing"
)

func Test(t *testing.T) {
	//Get client
	minioclient := api.GetClient()
	//Get BucketList
	api.GetBucketList(minioclient)
	//api.RemoveBucket(minioclient, "apptest")
	//make a new bucket
	if found,_ := minioclient.BucketExists(context.Background(),"apptest"); found  {
		api.RemoveBucket(minioclient, "apptest")
	}
	api.MakeBucket(minioclient, "apptest")
	//api.MakeBucket(minioclient, "apptest")
	//upload a file
	api.PutObjectFile(minioclient, "apptest", "test.jpg", "test.jpg")

	//Get BucketList
	api.GetBucketList(minioclient)
	//Get objects from bucketname
	api.ListObjects(minioclient, "apptest", "")
	//Download object to path.
	api.GetObjectFile(minioclient, "apptest", "test.jpg", "test1.jpg")
	api.StatObject(minioclient, "apptest", "test.jpg")
	api.RemoveObject(minioclient, "apptest", "test.jpg")
	api.RemoveBucket(minioclient, "apptest")
	api.GetBucketList(minioclient)
	//api.ListObjects(minioclient, "apptest", "")
}
