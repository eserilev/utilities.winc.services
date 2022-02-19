package winc_s3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadFile(bucket string, key string, content []byte) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default-winc"))
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)

	response, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(content),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Uploaded version " + *response.VersionId + " of " + key + " to bucket " + bucket)
}

func GetFile(bucket string, key string) *json.Decoder {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("default-winc"))
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)

	response, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	return decoder
}
