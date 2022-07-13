package persistence

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDbContext struct {
	*dynamodb.Client
	context.Context
}

type DynamoTable int

// TABLES Enum
const (
	USERS DynamoTable = iota
	PRODUCTS
)

func (tableIndx DynamoTable) GetTableNamePtr() *string {
	return aws.String([]string{
		"users",
		"products",
	}[tableIndx])
}

func InitDynamoDb(ctx context.Context) *DynamoDbContext {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(os.Getenv("AWS_DYNAMODB_REGION")),
		config.WithCredentialsProvider(
			aws.NewCredentialsCache(
				credentials.NewStaticCredentialsProvider(
					os.Getenv("AWS_ACCESS_KEYID"),
					os.Getenv("AWS_SECRET_KEY"),
					os.Getenv("AWS_SESSION"),
				),
			),
		),
		// config.WithEndpointResolver(
		// 	aws.EndpointResolverFunc(
		// 		func(service, region string) (aws.Endpoint, error) {
		// 			return aws.Endpoint{URL: os.Getenv("AWS_DYNAMODB_HOST")}, nil
		// 		}),
		// ),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)

	return &DynamoDbContext{
		svc, ctx,
	}

}
