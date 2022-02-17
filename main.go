package main

import (
	"context"
	"fmt"
	"log"
	"os"

	guuid "github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)


func main() {
	ctx := context.Background()
	endpoint := "storage.cloud.croc.ru"
	accessKeyID := ""
	secretAccessKey := ""
	useSSL := false


	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		bucketName := "my_s3_minio_go_test"
		id := guuid.New()
		objectName := fmt.Sprintf("test" + id.String() + ".jpg")
		filePath := "test1.jpg"

		fileReader, errOpen := os.Open(filePath)
		if errOpen != nil {
			panic(1)
		}

		fileStat, errStat := fileReader.Stat()
		if errStat != nil {
			panic(1)
		}

		fileSize := fileStat.Size()

		contentType := "image/jpeg"
		info, err := minioClient.PutObject(ctx, bucketName, objectName, fileReader, fileSize, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Fatalln(err)
			break
		}

		log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	}
}
