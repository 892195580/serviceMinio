package api

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/replication"
	"github.com/minio/minio-go/v7/pkg/tags"
	"log"
	"os"
)

func GetBucketList(minioClient *minio.Client) {
	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatalln("Failed to get bucket list:\n", err)
		return
	}
	log.Println(fmt.Sprintf("%d buckets in all.", len(buckets)))
	for _, bucket := range buckets {
		fmt.Println(fmt.Sprintf("%+v", bucket))
	}
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

func MakeBucket(minioClient *minio.Client, bucketName string) {
	opts := minio.MakeBucketOptions{
		Region: "us-east-1",
	}

	err := minioClient.MakeBucket(context.Background(), bucketName, opts)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to make bucket : %s", bucketName),err)
		return
	}
	log.Println(fmt.Sprintf("Make a new bucket : %s successful!", bucketName))
}

func RemoveBucket(minioClient *minio.Client, bucketName string) {
	// This operation will only work if your bucket is empty.
	err := minioClient.RemoveBucket(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Failed to remove bucket : %s", bucketName), err)
		return
	}
	log.Println(fmt.Sprintf("Remove bucket : %s successful", bucketName))
}

//todo: add to cli
func ListIncompleteUploads(minioClient *minio.Client, bucketName, prefix string) {
	// List all multipart uploads from a bucket-name with a matching prefix.
	for multipartObject := range minioClient.ListIncompleteUploads(context.Background(), bucketName, prefix, true) {
		if multipartObject.Err != nil {
			fmt.Println(multipartObject.Err)
			return
		}
		fmt.Println(multipartObject)
	}
}

//todo: add to cli
func GetBucketTagging (minioClient *minio.Client, bucketName string) {
	t, err := minioClient.GetBucketTagging(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Fetched Object Tags: %s", t)
}

//todo: add to cli
func SetBucketTagging (minioClient *minio.Client, bucketName, tag, tagValue string) {
	t, err := tags.MapToBucketTags(map[string]string{
		tag: tagValue,
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = minioClient.SetBucketTagging(context.Background(), bucketName, t)
	if err != nil {
		log.Fatalln("Failed to set Bucket Tag ", err)
		return
	}
	log.Fatalln("Set Bucket Tag successful!")
}

//todo: add to cli
func RemoveBucketTagging(minioClient *minio.Client, bucketName string) {
	//remove all tags
	err := minioClient.RemoveBucketTagging(context.Background(), bucketName)
	if err != nil {
		log.Fatalln("Failed to remove Bucket Tags", err)
		return
	}
	log.Fatalln("Remove Bucket Tags successful!")
}

//todo: add to cli
func GetBucketReplication(minioClient *minio.Client, bucketName string){
	// Get bucket replication configuration from S3
	replicationCfg, err := minioClient.GetBucketReplication(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	// Create replication config file
	localReplicationCfgFile, err := os.Create("replication.xml")
	if err != nil {
		log.Fatalln(err)
	}
	defer localReplicationCfgFile.Close()

	replBytes, err := json.Marshal(replicationCfg)
	if err != nil {
		log.Fatalln(err)
	}
	localReplicationCfgFile.Write(replBytes)
}
//todo: add to cli
func SetBucketReplication (minioClient *minio.Client, bucketName string){
	//todo: complete it with json or file

	replicationStr := `<ReplicationConfiguration><Rule><ID>string</ID><Status>Enabled</Status><Priority>1</Priority><DeleteMarkerReplication><Status>Disabled</Status></DeleteMarkerReplication><Destination><Bucket>arn:aws:s3:::dest</Bucket></Destination><Filter><And><Prefix>Prefix</Prefix><Tag><Key>Tag-Key1</Key><Value>Tag-Value1</Value></Tag><Tag><Key>Tag-Key2</Key><Value>Tag-Value2</Value></Tag></And></Filter></Rule></ReplicationConfiguration>`
	var replCfg replication.Config
	err := xml.Unmarshal([]byte(replicationStr), &replCfg)
	if err != nil {
		log.Fatalln(err)
	}

	// This replication ARN should have been generated for replication endpoint using `mc admin bucket remote` command
	replCfg.Role = "arn:minio:replica::dadddae7-f1d7-440f-b5d6-651aa9a8c8a7:dest"
	// Set replication config on a bucket
	err = minioClient.SetBucketReplication(context.Background(), "my-bucketname", replCfg)
	if err != nil {
		log.Fatalln(err)
	}
}

//todo: add to cli
func RemoveBucketReplication(minioClient *minio.Client, bucketName string){
	// Remove replication configuration on a bucket
	err := minioClient.RemoveBucketReplication(context.Background(), bucketName)

	if err != nil {
		log.Fatalln(err)
	}
}

//todo: add to cli



//todo: add to cli




























