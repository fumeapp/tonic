package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

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

func DynamoIndex(region string, table string, out interface{}) error {

	config, err := cfg(region)

	if err != nil {
		return err
	}

	p, err := dynamodb.NewFromConfig(config).ExecuteStatement(context.Background(),
		&dynamodb.ExecuteStatementInput{
			Statement:      aws.String(fmt.Sprintf("SELECT * FROM \"%s\"", table)),
			ConsistentRead: aws.Bool(false),
		},
	)

	if err != nil {
		return err
	}

	if err := attributevalue.UnmarshalListOfMaps(p.Items, &out); err != nil {
		return err
	}

	return nil

}
