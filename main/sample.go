package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

const tableNameUser = "User"

type User struct {
	ID          string `dynamo:"ID,hash"`
	Name        string `dynamo:"Name"`
	Description string `dynamo:"Description"`
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String(os.Getenv("DYNAMO_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials("dummy-id", "dummy-secret", "dummy-token"),
	})
	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess)

	db.Table(tableNameUser).DeleteTable().Run()

	err = db.CreateTable(tableNameUser, User{}).Run()
	if err != nil {
		panic(fmt.Errorf("failed to CreateTable %w", err))
	}

	fmt.Println("create is successful.")

	user := db.Table(tableNameUser)
	id, _ := uuid.NewUUID()
	err = user.Put(&User{
		ID:          id.String(),
		Name:        "test-name",
		Description: "test-description",
	}).Run()
	if err != nil {
		panic(fmt.Errorf("failed to Put %w", err))
	}

	var registeredUser User
	err = user.Get("ID", id.String()).One(&registeredUser)
	if err != nil {
		panic(fmt.Errorf("failed to Get %w", err))
	}
	fmt.Printf("insert is successful. registeredUser: %+v\n", registeredUser)

	err = user.Update("ID", id.String()).Set("Name", "test-name2").Set("Description", "test-description2").Run()
	if err != nil {
		panic(fmt.Errorf("failed to Update %s", err))
	}

	err = user.Get("ID", id.String()).One(&registeredUser)
	if err != nil {
		panic(fmt.Errorf("failed to Get %w", err))
	}

	fmt.Printf("update is successful. updatedUser: %+v\n", registeredUser)

	err = user.Delete("ID", id.String()).Run()
	if err != nil {
		panic(fmt.Errorf("failed to Delete %w", err))
	}

	err = user.Get("ID", id.String()).One(&registeredUser)
	if err != nil {
		if err.Error() != "dynamo: no item found" {
			panic(fmt.Errorf("failed to Get %w", err))
		}

		fmt.Println("delete is successful.")
	}
}
