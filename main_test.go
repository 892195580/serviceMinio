package main

import (
	"fmt"
	"minio/service"
	"testing"
)

func TestDownLoad(t *testing.T) {
	//Get client
	minioclient := service.GetClient()
	//Get BucketList
	buckets, err := service.GetBucketList(minioclient)
	if err != nil {
		t.Log("GetBucketList Failed")
	}
	t.Log("Load buckets successful!")
	for _, bucket := range buckets {
		t.Log(bucket)
	}
	//Get objects from bucketname
	objectsInfo, err := service.GetObjects(minioclient, "test", "")
	if err != nil {
		t.Log("GetObjects Failed")
	}
	for _,v := range objectsInfo {
		fmt.Println(fmt.Sprintf("%+v", v))
	}
	t.Log("Load Objects successful!")
	//Download object to path.
	service.DownloadObject(minioclient, "test", "20.jpg", "20.jpg")
}
