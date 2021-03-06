package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/shashijangra22/bootcamp-project/pkg/populate"
)

type AWS_STRUCT struct {
	AWS_KEY_ID     string
	AWS_SECRET_KEY string
	REGION         string
}

var secret AWS_STRUCT

func createDBSession(filename string) *dynamodb.DynamoDB {
	secretsFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening secrets.json!")
		os.Exit(1)
	}
	defer secretsFile.Close()
	byteValue, _ := ioutil.ReadAll(secretsFile)
	json.Unmarshal(byteValue, &secret)
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(secret.REGION),
		Endpoint:    aws.String("http://localhost:8000"), // [for local should be http] remove this line to connect to cloud dynamodDB with creds in secrets.json file
		Credentials: credentials.NewStaticCredentials(secret.AWS_KEY_ID, secret.AWS_SECRET_KEY, ""),
	}))
	db := dynamodb.New(sess)
	return db
}

func main() {
	path, _ := os.Getwd()
	db := createDBSession(path + "/secrets.json")
	args := os.Args[1:]
	if len(args) == 1 {
		populate.Customers(db, path+"/"+args[0]+"/customers.csv", "Team2-CUSTOMERS")
		populate.Orders(db, path+"/"+args[0]+"/orders.csv", "Team2-ORDERS")
		populate.Restaurants(db, path+"/"+args[0]+"/restaurants.csv", "Team2-RESTAURANTS")
		fmt.Println("Done populating all the tables :)")
	} else {
		fmt.Println("Please give directory path of DB schema as an argument!")
	}
}
