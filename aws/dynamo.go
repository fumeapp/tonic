package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoIndexQuery struct {
	Region      string
	Table       string
	Wheres      DynamoIndexWheres
	OrderField  string
	Direction   string
	SetLimit    *int32
	ExposeQuery bool
}

type DynamoIndexWhere struct {
	Field string
	Type  string
	Value interface{}
}

type DynamoIndexWheres []DynamoIndexWhere

// Add - append to the w
func (w *DynamoIndexWheres) Add(field string, fType string, value interface{}) {
	*w = append(*w, DynamoIndexWhere{Field: field, Type: fType, Value: value})
}

// hide query if not specified
type DynamoIndexResults struct {
	NextToken *string     `json:"next_token,omitempty"`
	Items     interface{} `json:"items"`
	Query     *string     `json:"query,omitempty"`
}

func DynamoIndex(region string, table string) DynamoIndexQuery {
	var query DynamoIndexQuery
	query.Region = region
	query.Table = table
	return query
}

func (q DynamoIndexQuery) ShowQuery(value bool) DynamoIndexQuery {
	q.ExposeQuery = value
	return q
}

func (q DynamoIndexQuery) WhereStr(field string, value string) DynamoIndexQuery {
	q.Wheres.Add(field, "string", value)
	return q
}

func (q DynamoIndexQuery) WhereBool(field string, value bool) DynamoIndexQuery {
	q.Wheres.Add(field, "bool", value)
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

func (w DynamoIndexWhere) ToString() string {
	switch w.Type {
	case "string":
		return fmt.Sprintf(`"%s" = '%s'`, w.Field, w.Value)
	case "bool":
		if w.Value == true {
			return fmt.Sprintf(`"%s" = %s`, w.Field, "true")
		} else {
			return fmt.Sprintf(`"%s" = %s`, w.Field, "false")
		}
	default:
		return ""
	}
}

func (q DynamoIndexQuery) Get(out interface{}) (*DynamoIndexResults, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s"`, q.Table)
	for index, where := range q.Wheres {
		if index == 0 {
			query = fmt.Sprintf(`%s WHERE %s`, query, where.ToString())
		} else {
			query = fmt.Sprintf(`%s AND %s`, query, where.ToString())
		}
	}
	if q.OrderField != "" && q.Direction != "" {
		query = fmt.Sprintf(`%s ORDER BY "%s" %s`,
			query,
			q.OrderField,
			q.Direction,
		)
	}

	p, err := dynamodb.NewFromConfig(Config()).ExecuteStatement(context.Background(),
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
	if q.ExposeQuery {
		results.Query = &query
	}

	return &results, nil
}

func DynamoInsert(region string, table string, params any) error {
	item, err := attributevalue.MarshalMap(params)
	if err != nil {
		return err
	}

	if _, err := dynamodb.NewFromConfig(Config()).PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      item,
	}); err != nil {
		return err
	}

	return nil
}
