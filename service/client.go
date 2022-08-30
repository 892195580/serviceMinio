package service
import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"minio/config"
)
var selfMinioClient *minio.Client
func GetClient() *minio.Client {
	return selfMinioClient
}
func init() {
	endpoint, accessKeyID, secretAccessKey, useSSL := config.GetCredentials()
	newminioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("Failed to connect to server. \n", err)
	} else {
		selfMinioClient = newminioClient
	}
}
