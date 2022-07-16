package stacks

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type KewlEc2StackProps struct {
	awscdk.StackProps
}

type KewlEc2Stack struct {
	sprops awscdk.StackProps
	scope constructs.Construct
	stack awscdk.Stack
	
	resources []any
}
