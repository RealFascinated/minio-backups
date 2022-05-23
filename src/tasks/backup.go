package tasks

import (
	"context"
	"github.com/go-co-op/gocron"
	"github.com/minio/minio-go/v7"
	"log"
	"minio_backups/src/data"
	"os"
	"time"
)

func Start() {
	dataDir := data.ConfigCache.DataDirectory
	buckets := data.ConfigCache.MinioSettings.Buckets
	minioClient := data.GetMinio()

	log.Println("Initializing backup task")
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(10).Minute().Do(func() {
		log.Println("Starting backup task")

		for _, bucket := range buckets {
			log.Println("Starting backup for bucket: ", bucket)

			// Check if the backup directory exists
			err := os.MkdirAll(dataDir+"/"+bucket, os.ModePerm)
			if err != nil {
				log.Println("Error creating data directory:", err)
				return
			}

			objects := minioClient.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{})
			for object := range objects {
				// Check if the object already is backed up
				if _, err := os.Stat(dataDir + "/" + bucket + "/" + object.Key); err == nil {
					continue
				}

				// Backup the object
				log.Println("Starting backup for object: ", object.Key)
				err := minioClient.FGetObject(context.Background(), bucket, object.Key, dataDir+"/"+bucket+"/"+object.Key, minio.GetObjectOptions{})
				if err != nil {
					log.Println("Error getting object:", err)
					return
				}
			}
		}
	})
	if err != nil {
		log.Println(err)
		return
	}
	scheduler.StartBlocking()
}
