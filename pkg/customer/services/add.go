package services

import (
	"MyApp/pkg/Models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AddOne(db *dynamodb.DynamoDB, cst Models.Customer) {
	cstItem, err := dynamodbattribute.MarshalMap(cst)
	if err != nil {
		panic("Cannot map the values given in Customer struct for post request...")
	}
	params := &dynamodb.PutItemInput{
		TableName: aws.String("customers"),
		Item:      cstItem,
	}
	_, err = db.PutItem(params)
	if err != nil {
		panic("Error in putting the customer item")
	}
}