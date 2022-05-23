package main

import (
	"log"
	"minio_backups/src/data"
	"minio_backups/src/tasks"
)

func main() {
	log.Println("Starting minio-backups")
	data.InitConfig()
	data.InitMinio()
	tasks.Start()
}
