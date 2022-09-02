package api

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"net/url"
	"time"
)

// Get the infomation of file in form of list
func ListObjects(minioClient *minio.Client, bucketname, prefix string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}
	opts := minio.ListObjectsOptions{
		UseV1:     true,
		Prefix:    prefix,
		Recursive: true,
	}
	countObjects := 0
	// List all objects from a bucket-name with a matching prefix.
	for object := range minioClient.ListObjects(context.Background(), bucketname, opts) {
		if object.Err != nil {
			log.Fatalln(object.Err)
			return
		}
		countObjects++
		fmt.Println(fmt.Sprintf("%+v", object))
	}
	log.Println(fmt.Sprintf("%d files in bucket { %s }.", countObjects, bucketname))
}

//Download file:bucketname/objectname to path.
func GetObjectFile(minioClient *minio.Client, bucketname, objectname, path string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}
	if path == "" {
		log.Fatalln("Empty path.")
	}
	err := minioClient.FGetObject(context.Background(), bucketname, objectname, path, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fmt.Sprintf("Download %s:%s to %s successful! ", bucketname, objectname, path))
}
//Upload path:file to bucketname/objectname .
func PutObjectFile(minioClient *minio.Client, bucketname, objectname, path string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}
	info, err := minioClient.FPutObject(context.Background(), bucketname, objectname, path, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fmt.Sprintf("Upload %s, to %s:%s successful! It's details are :", path, bucketname, objectname))
	fmt.Println(info)
}

//Show status of bucketname/objectname.
func StatObject(minioClient *minio.Client, bucketname, objectname string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}
	stat, err := minioClient.StatObject(context.Background(), bucketname, objectname, minio.StatObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(fmt.Sprintf("Detailes of object :\n %+v", stat))
}

//Remove an object : bucketname/objectname.
func RemoveObject(minioClient *minio.Client, bucketname, objectname string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := minioClient.RemoveObject(context.Background(), bucketname, objectname, opts)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fmt.Sprintf("Remove %s:%s successful! It's details are :", bucketname, objectname))
}
//Remove all objects in bucketname.
func RemoveObjects(minioClient *minio.Client, bucketname, prefix string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}
	objectsCh := make(chan minio.ObjectInfo)

	// Send object names that are needed to be removed to objectsCh
	go func() {
		defer close(objectsCh)
		// List all objects from a bucket-name with a matching prefix.
		opts := minio.ListObjectsOptions{Prefix: prefix, Recursive: true}
		for object := range minioClient.ListObjects(context.Background(), bucketname, opts) {
			if object.Err != nil {
				log.Fatalln(object.Err)
			}
			objectsCh <- object
		}
	}()

	// Call RemoveObjects API
	errorCh := minioClient.RemoveObjects(context.Background(), bucketname, objectsCh, minio.RemoveObjectsOptions{})

	// Print errors received from RemoveObjects API
	for e := range errorCh {
		log.Fatalln("Failed to remove " + e.ObjectName + ", error: " + e.Err.Error())
	}
	log.Println(fmt.Sprintf("Remove all objects in bucket:%s successful!", bucketname))
}

//todo: add to cli
//Generates a presigned URL for HTTP GET operations : bucketname/objectname.
func PresignedGetObject(minioClient *minio.Client, bucketname, objectname string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}
	// Set request parameters
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", objectname))

	// Gernerate presigned get object url.
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), bucketname, objectname, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(presignedURL)
}

//todo: add to cli
//Generates a presigned URL for HTTP Put operations : bucketname/objectname.
func PresignedPutObject(minioClient *minio.Client, bucketname, objectname string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}

	// Gernerate presigned Put object url.
	presignedURL, err := minioClient.PresignedPutObject(context.Background(), bucketname, objectname, time.Duration(1000)*time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(presignedURL)
}
//todo: add to cli
//Generates a presigned URL for HTTP Head operations : bucketname/objectname.
func PresignedHeadObject(minioClient *minio.Client, bucketname, objectname string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}

	// Set request parameters
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", objectname))
	// Gernerate presigned Put object url.
	presignedURL, err := minioClient.PresignedHeadObject(context.Background(), bucketname, objectname, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(presignedURL)
}
//todo: add to cli
//Generates a presigned URL for HTTP Post operations : bucketname/objectname.
func PresignedPostObject(minioClient *minio.Client, bucketname, objectname string) {
	if _, err := minioClient.BucketExists(context.Background(), bucketname); err != nil {
		log.Fatalln(err)
	}
	policy := minio.NewPostPolicy()
	policy.SetBucket(bucketname)
	policy.SetKey(objectname)
	// Expires in 10 days.
	policy.SetExpires(time.Now().UTC().AddDate(0, 0, 10))
	// Returns form data for POST form request.
	presignedURL, formData, err := minioClient.PresignedPostPolicy(context.Background(), policy)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("curl ")
	for k, v := range formData {
		fmt.Printf("-F %s=%s ", k, v)
	}
	fmt.Printf("-F file=@/etc/bash.bashrc ")
	fmt.Printf("%s\n", presignedURL)
}