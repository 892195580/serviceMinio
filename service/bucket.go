package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"log"
)

type  H struct {
	mes string
}
func GetBuckets(minioClient *minio.Client) ([]minio.BucketInfo, error){
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		return nil, err
	}
	for _, bucket := range buckets {
		log.Println(bucket)
	}
	return buckets, err
}
func GetBucketList(c *gin.Context) {
	//c.JSON(http.StatusOK, H{
	//	"OK",
	//})
	reader, err := minioClient.GetObject(context.Background(), "my-bucketname", "my-objectname", minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()
	return
}
