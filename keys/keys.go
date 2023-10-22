package keys

import (
	"context"
	"io"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func UploadImageToGoogleStorage(image *multipart.FileHeader) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : Cannot load config file,", err.Error())
	}

	ctx := context.Background()

	// Gantilah credentialFile dengan path ke credential file JSON Anda
	credentialFile := os.Getenv("GOOGLE_CLOUD_CREDENTIALS_PATH")
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFile))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Gantilah bucketName sesuai dengan nama bucket Google Cloud Storage Anda
	bucketName := "garvice"
	imagePath := "file_pickup/" + uuid.New().String() + ".jpg"

	wc := client.Bucket(bucketName).Object(imagePath).NewWriter(ctx)
	defer wc.Close()
	file, err := image.Open()
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	imageURL := "https://storage.googleapis.com/" + bucketName + "/" + imagePath

	return imageURL, nil
}
