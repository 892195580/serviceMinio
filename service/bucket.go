package service

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
)

func GetBucketList(minioClient *minio.Client) ([]minio.BucketInfo, error){
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		return nil, err
	}
	log.Println(fmt.Sprintf("%d buckets in all.", len(buckets)))
	return buckets, err
}

func IsBucketExists(minioClient *minio.Client, bucketname string) (bool, error){
	found, err := minioClient.BucketExists(context.Background(), bucketname)
	if err != nil {
		log.Fatalln(err)
	}

	if !found {
		log.Println("Bucket not found.")
		return found, err
	}
	log.Println("Bucket found.")
	return found, nil
}
