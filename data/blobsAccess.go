package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Presigner struct {
	PresignClient *s3.PresignClient
}

func (presigner Presigner) GetObject(
	bucketName string, objectKey string, lifeTimeSecs int) (*v4.PresignedHTTPRequest, error) {
	request, err := presigner.PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifeTimeSecs * int(time.Second))
	})

	if err != nil {
		log.Printf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}

	return request, err
}

func generateS3Url(objectKey string) string {
	bucketName := "baum-image-data"
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		log.Fatal(err)
	}

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)
	presigner := Presigner{PresignClient: presignClient}

	presignedGetReq, err := presigner.GetObject(bucketName, objectKey, 600)

	if err != nil {
		fmt.Println("error getting presignedURL:", err)
	}

	fmt.Println(presignedGetReq.URL)
	blobUrl := presignedGetReq.URL
	return blobUrl
}
