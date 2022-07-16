package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	jsii "github.com/aws/jsii-runtime-go"
	"github.com/r351574nc3/aws-cdk-golang-excercise/stacks"
)

type AppStackProps struct {
	awscdk.StackProps
}

func main() {
	var props *stacks.KewlEc2StackProps
	var stack *stacks.KewlEc2Stack
	app := awscdk.NewApp(nil)

	props = props.New(env())
	stack = stack.New(app, "KewlEc2Stack", props)

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
