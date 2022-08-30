package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
)

// Get the infomation of file in form of list
func GetObjects(minioClient *minio.Client, bucketname, prefix string) ([]minio.ObjectInfo , error){
	if e,err := IsBucketExists(minioClient, bucketname); !e || err != nil {
		return nil, err
	}
	opts := minio.ListObjectsOptions{
		UseV1:     true,
		Prefix:    prefix,
		Recursive: true,
	}
	objects := []minio.ObjectInfo{}
	// List all objects from a bucket-name with a matching prefix.
	for object := range minioClient.ListObjects(context.Background(), bucketname, opts) {
		if object.Err != nil {
			//fmt.Println(object.Err)
			return nil, object.Err
		}
		objects = append(objects, object)
		//fmt.Println(object)
	}
	log.Println(fmt.Sprintf("%d files in bucket %s.", len(objects), bucketname))
	return objects, nil
}

//Download file:bucketname/objectname to path.
func DownloadObject(minioClient *minio.Client, bucketname, objectname, path string) error {
	if e,err := IsBucketExists(minioClient, bucketname); e && err != nil {
		return err
	}
	if path == "" {
		log.Println("Invalid path.")
		return errors.New("Invalid path")
	}
	err := minioClient.FGetObject(context.Background(), bucketname, objectname, path, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
