package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoIndexQuery struct {
	Region     string
	Table      string
	Wheres     DynamoIndexWheres
	OrderField string
	Direction  string
	SetLimit   *int32
}

type DynamoIndexWhere struct {
	Field string
	Value string
}

type DynamoIndexWheres []DynamoIndexWhere

// Add - append to the w
func (w *DynamoIndexWheres) Add(field string, value string) {
	*w = append(*w, DynamoIndexWhere{Field: field, Value: value})
}

type DynamoIndexResults struct {
	NextToken *string
	Items     interface{}
}

func DynamoIndex(region string, table string) DynamoIndexQuery {
	var query DynamoIndexQuery
	query.Region = region
	query.Table = table
	return query
}

func (q DynamoIndexQuery) Where(field string, value string) DynamoIndexQuery {
	q.Wheres.Add(field, value)
	return q
}
func (q DynamoIndexQuery) Order(field string) DynamoIndexQuery {
	q.OrderField = field
	return q
}

func (q DynamoIndexQuery) Desc() DynamoIndexQuery {
	q.Direction = "DESC"
	return q
}
func (q DynamoIndexQuery) Asc() DynamoIndexQuery {
	q.Direction = "ASC"
	return q
}

func (q DynamoIndexQuery) Limit(limit int) DynamoIndexQuery {
	q.SetLimit = aws.Int32(int32(limit))
	return q
}

func (q DynamoIndexQuery) Get(out interface{}) (*DynamoIndexResults, error) {
	config, err := cfg(q.Region)

	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf(`SELECT * FROM "%s"`, q.Table)
	for index, where := range q.Wheres {
		if index == 0 {
			query = fmt.Sprintf(`%s WHERE "%s" = '%s'`, query, where.Field, where.Value)
		} else {
			query = fmt.Sprintf(`%s AND "%s" = '%s'`, query, where.Field, where.Value)
		}
	}
	query = fmt.Sprintf(`%s ORDER BY "%s" %s`,
		query,
		q.OrderField,
		q.Direction,
	)
	p, err := dynamodb.NewFromConfig(config).ExecuteStatement(context.Background(),
		&dynamodb.ExecuteStatementInput{
			Statement:      aws.String(query),
			ConsistentRead: aws.Bool(false),
			Limit:          q.SetLimit,
		},
	)

	if err != nil {
		return nil, err
	}

	if err := attributevalue.UnmarshalListOfMaps(p.Items, &out); err != nil {
		return nil, err
	}
	var results DynamoIndexResults
	results.Items = out
	results.NextToken = p.NextToken

	return &results, nil
}

func DynamoInsert(region string, table string, params any) error {

	config, err := cfg(region)
	if err != nil {
		return err
	}

	item, err := attributevalue.MarshalMap(params)
	if err != nil {
		return err
	}

	if _, err := dynamodb.NewFromConfig(config).PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      item,
	}); err != nil {
		return err
	}

	return nil
}
