package repositories

import (
	"fmt"

	"github.com/92Sam/ms-users/domain/models"
	"github.com/92Sam/ms-users/domain/persistence"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type IProductRepository interface {
	Create(productModel *models.Product) (*models.Product, error)
	GetProducts() ([]*models.Product, error)
}

type ProductRepository struct {
	*persistence.Neo4jDbContext
}

func NewProductRepository(neo4jPersistence *persistence.Neo4jDbContext) IProductRepository {
	return &ProductRepository{
		neo4jPersistence,
	}
}

func (ps *ProductRepository) Create(productModel *models.Product) (*models.Product, error) {
	// values := productModel.GetMapString()
	// fmt.Printf("%#c", values
	session := ps.Neo4jDbContext.Session
	defer session.Close()
	query := "CREATE (n:Product { id: $id, name: $name, description: $description, createAt: $createAt }) RETURN n.id, n.name"
	records, err := session.Run(query, map[string]interface{}{
		"id":          productModel.Id,
		"name":        productModel.Name,
		"description": productModel.Description,
		"rating":      productModel.Rating,
		"createAt":    neo4j.DateOf(productModel.CreateAt),
	})

	// In face of driver native errors, make sure to return them directly.
	// Depending on the error, the driver may try to execute the function again.
	if err != nil {
		fmt.Println("Insert Error")
		return nil, err
	}

	record, err := records.Single()
	if err != nil {
		return nil, err
	}

	return &models.Product{
		Id:   record.Values[0].(string),
		Name: record.Values[1].(string),
	}, nil

}

func (ps *ProductRepository) GetProducts() ([]*models.Product, error) {
	session := ps.Neo4jDbContext.Session
	defer session.Close()
	records, err := session.Run("MATCH (n:Product) RETURN n", map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	recordColletion, err := records.Collect()
	if err != nil {
		return nil, err
	}
	totalRecords := len(recordColletion)
	var results = make([]*models.Product, totalRecords)

	for k, _ := range recordColletion {
		node := recordColletion[k].Values[0].(neo4j.Node)
		results[k] = &models.Product{
			Id:          node.Props["id"].(string),
			Name:        node.Props["name"].(string),
			Description: node.Props["description"].(string),
			// Rating:      node.Props["rating"].(int8),
			// CreateAt: node.Props["createAt"].(time.Time),
		}
	}

	return results, nil
}
