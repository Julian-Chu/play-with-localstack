package main

import (
	"context"
	b64 "encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

func main() {
	mockS3EndpointResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		//if service == s3.ServiceID {
		return aws.Endpoint{
			PartitionID:       "aws",
			URL:               "http://localhost:4566",
			SigningRegion:     "eu-central-1",
			HostnameImmutable: true, // don't rewrite aws-s3:8080/<bucket> to <bucket>.aws-s3:8080
		}, nil
		//}
		//return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithEndpointResolver(mockS3EndpointResolver))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := kms.NewFromConfig(cfg)
	text := "test"
	keyId := "arn:aws:kms:eu-central-1:000000000000:key/2091c2f4-87d8-441e-8e63-cb4850cf6be6"
	input := &kms.EncryptInput{
		KeyId:     &keyId,
		Plaintext: []byte(text),
	}

	//result, err := EncryptText(context.TODO(), client, input)
	output, err := client.Encrypt(context.TODO(), input)
	if err != nil {
		fmt.Println("Got error encrypting data:")
		fmt.Println(err)
		return
	}

	fmt.Println(output)
	blobString := b64.StdEncoding.EncodeToString(output.CiphertextBlob)

	fmt.Println(blobString)
}
