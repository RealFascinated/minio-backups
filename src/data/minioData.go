package data

import (
	"github.com/minio/minio-go/v7"
	"log"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func InitMinio() {
	log.Println("Initializing Minio")
	endpoint := ConfigCache.MinioSettings.Endpoint
	accessKey := ConfigCache.MinioSettings.AccessKey
	secretKey := ConfigCache.MinioSettings.SecretKey
	useSSL := ConfigCache.MinioSettings.UseSSL

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Connected to Minio")
	minioClient = client
}

func GetMinio() *minio.Client {
	return minioClient
}
