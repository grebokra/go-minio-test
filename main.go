package main

import (
	"context"
	"log"
        "fmt"

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
		contentType := "image/jpeg"
		info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Fatalln(err)
		        break
                }

		log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	}
}
