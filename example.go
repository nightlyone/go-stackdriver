package main

import (
	"github.com/nightlyone/go-stackdriver/stackdriver"
	"log"
	"time"
)

func main() {
	now := time.Now().Unix()
	query := &stackdriver.Metrics{
		Timestamp:    now,
		ProtoVersion: 1,
		Data: []stackdriver.Datapoint{
			stackdriver.Datapoint{
				Name:        "one",
				Value:       2,
				CollectedAt: now,
			},
		},
	}

	if err := query.Submit(); err != nil {
		log.Fatalln(err)
	}
	log.Println("sent", query, "to apiendpoint")
	deploy := &stackdriver.Deploy{
		RevisionId: "87230611cdc7e5ff7723a91e715367c553ad1115",
		DeployedBy: "John Doe",
		DeployedTo: "production",
		Repository: "example_repo",
	}
	if err := deploy.Submit(); err != nil {
		log.Fatalln(err)
	}
	log.Println("sent", deploy, "to apiendpoint")
	annotate := &stackdriver.Annotation{
		Message:     "Started moving more services to the cloud",
		AnnotatedBy: "devops",
		Level:       "INFO",
	}
	if err := annotate.Submit(); err != nil {
		log.Fatalln(err)
	}
	log.Println("sent", annotate, "to apiendpoint")
}
