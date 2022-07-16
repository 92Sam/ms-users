package repositories

import (
	"fmt"
	"time"

	"github.com/92Sam/ms-users/domain/models"
	"github.com/92Sam/ms-users/domain/persistence"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type IProductRepository interface {
	Create(productModel *models.Product) (*models.Product, error)
}

type ProductRepository struct {
	*persistence.DynamoDbContext
}

func NewProductRepository(dynamoPersistence *persistence.DynamoDbContext) IProductRepository {
	return &ProductRepository{
		dynamoPersistence,
	}
}

func (ps *ProductRepository) Create(productModel *models.Product) (*models.Product, error) {

	out, err := ps.PutItem(ps.Context, &dynamodb.PutItemInput{
		TableName:           persistence.USERS.GetTableNamePtr(),
		ConditionExpression: aws.String("attribute_not_exists(pk)"),
		Item: map[string]types.AttributeValue{
			"id":          &types.AttributeValueMemberS{Value: uuid.NewString()},
			"name":        &types.AttributeValueMemberS{Value: productModel.Name},
			"description": &types.AttributeValueMemberS{Value: productModel.Description},
			"createdAt":   &types.AttributeValueMemberS{Value: time.Now().String()},
		},
	})

	if err != nil {
		return nil, err
	}

	fmt.Println(out.Attributes)

	return productModel, nil
}
