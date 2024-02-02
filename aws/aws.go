package aws

import (
	"bytes"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UploadData struct {
	FileLocation string
	FileName     string
	Key          string
}

type AwsCredentials struct {
	Region     string `mapstructure:"region"`
	AccessKey  string `mapstructure:"access_key"`
	SecretKey  string `mapstructure:"secret_key"`
	BucketName string `mapstructure:"bucket_name"`
}

type AwsClient interface {
	UploadFile() error
}

type awsClient struct {
	session     *session.Session
	credentials AwsCredentials
}

func NewAwsClient(c AwsCredentials) *awsClient {
	// Create an AWS session
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(c.Region),
		Credentials: credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, ""),
	})
	if err != nil {
		return nil
	}

	return &awsClient{
		session:     awsSession,
		credentials: c,
	}
}

func (a *awsClient) UploadFile(uploadData UploadData) error {
	// Check for parameters integrity
	if a.credentials.BucketName == "" || uploadData.FileLocation == "" || uploadData.FileName == "" {
		return errors.New("missing file parameters, could not upload to s3 bucket.")
	}

	// Open the MongoDB dump file
	file, err := os.Open(uploadData.FileLocation)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(a.session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(a.credentials.BucketName),
		Key:    aws.String(uploadData.Key + uploadData.FileName),
		Body:   bytes.NewReader(buffer),
	})
	return err
}
