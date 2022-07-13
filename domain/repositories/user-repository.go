package repositories

import (
	"context"
	"fmt"

	"github.com/92Sam/ms-users/domain/models"
	"github.com/92Sam/ms-users/domain/persistence"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type User interface {
	Create(models.User) (*models.User, error)
	GetAll() ([]models.User, error)
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(id string) (*models.User, error)
	Delete(id string) error
}

type UserRepository struct {
	*persistence.DynamoDbContext
}

func NewUserRepository(dynamoPersistence *persistence.DynamoDbContext) *UserRepository {
	return &UserRepository{
		dynamoPersistence,
	}
}

func (ur *UserRepository) Create(userModel *models.User) (*models.User, error) {

	_, err := ur.PutItem(ur.Context, &dynamodb.PutItemInput{
		TableName:           persistence.USERS.GetTableNamePtr(),
		ConditionExpression: aws.String("attribute_not_exists(pk)"),
		Item: map[string]types.AttributeValue{
			"id":        &types.AttributeValueMemberS{Value: userModel.Id},
			"name":      &types.AttributeValueMemberS{Value: userModel.Name},
			"email":     &types.AttributeValueMemberS{Value: userModel.Email},
			"password":  &types.AttributeValueMemberS{Value: userModel.Password},
			"createdAt": &types.AttributeValueMemberS{Value: userModel.CreateAt.String()},
		},
	})

	if err != nil {
		return nil, err
	}

	return userModel, nil
}

func (ur *UserRepository) GetByEmail(email string) (*models.User, error) {
	results, err := ur.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:        persistence.USERS.GetTableNamePtr(),
		FilterExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
		},
	})
	if err != nil {
		return nil, err
	}

	var record models.User
	if results.Count > 0 {
		err = attributevalue.UnmarshalMap(results.Items[0], &record)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal Items, %w", err)
		}
	}

	return &record, nil
}

func (ur *UserRepository) GetUsers() ([]*models.User, error) {

	results, err := ur.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:       persistence.USERS.GetTableNamePtr(),
		AttributesToGet: []string{"id", "email", "name", "createdAt"},
	})
	if err != nil {
		return nil, err
	}

	var records []*models.User
	err = attributevalue.UnmarshalListOfMaps(results.Items, &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Items, %w", err)
	}

	return records, nil
}

func (ur *UserRepository) DeleteUserById(id string) (bool, error) {

	_, err := ur.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: persistence.USERS.GetTableNamePtr(),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ur *UserRepository) GetUsersById(id string) (*models.User, error) {
	result, err := ur.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: persistence.USERS.GetTableNamePtr(),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		AttributesToGet: []string{"id", "email", "name", "createdAt"},
	})
	if err != nil {
		return nil, err
	}

	var records *models.User
	err = attributevalue.UnmarshalMap(result.Item, &records)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Items, %w", err)
	}

	return records, nil
}

func (ur *UserRepository) UpdateUserById(id string, user *models.User) (bool, error) {

	_, err := ur.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: persistence.USERS.GetTableNamePtr(),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET #nm = :name, email = :email"),
		ExpressionAttributeNames: map[string]string{
			"#nm": "name",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name":  &types.AttributeValueMemberS{Value: user.Name},
			":email": &types.AttributeValueMemberS{Value: user.Email},
			// ":updatedAt": &types.AttributeValueMemberS{Value: user.UpdatedAt.String()},
		},

		ReturnValues: types.ReturnValueUpdatedNew,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
