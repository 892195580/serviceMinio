package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"minio/service"
)

func main() {

	// 路由
	//endpoint, _, _, _ := config.GetCredentials()
	minioClient := service.GetClient()
	//r := gin.Default()
	//r.GET(endpoint + "/getbuckets", service.GetBucketList)
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
	opts := minio.ListObjectsOptions{
		UseV1:     true,
		Recursive: true,
	}

	// List all objects from a bucket-name with a matching prefix.
	for object := range minioClient.ListObjects(context.Background(), "enmotech-test", opts) {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}

}