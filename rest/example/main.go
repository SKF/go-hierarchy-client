package main

import (
	"context"
	"fmt"

	hierarchy "github.com/SKF/go-hierarchy-client/rest"
	"github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/auth/secretsmanagerauth"
	"github.com/SKF/go-utility/v2/stages"
	"github.com/aws/aws-sdk-go/aws/session"
)

const serviceName = "example-service"

func main() {
	ctx := context.Background()

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	client := hierarchy.NewClient(
		hierarchy.WithStage(stages.StageSandbox),
		client.WithDatadogTracing(serviceName),
		client.WithTokenProvider(&auth.SecretsManagerTokenProvider{
			Config: secretsmanagerauth.Config{
				WithDatadogTracing:       true,
				ServiceName:              serviceName,
				AWSSession:               sess,
				AWSSecretsManagerAccount: "633888256817",
				AWSSecretsManagerRegion:  "eu-west-1",
				SecretKey:                "user-credentials/" + serviceName,
				Stage:                    stages.StageSandbox,
			},
		}),
	)

	ancestors, err := client.GetAncestors(ctx, "a2b0645b-62e0-43c7-9d35-1f02fd245f67", 0, "company", "site")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ancestors)
}
