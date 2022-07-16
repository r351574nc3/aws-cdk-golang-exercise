package stacks

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (p *KewlEc2StackProps) New(env *awscdk.Environment) *KewlEc2StackProps {
	if p == nil {
		p = new(KewlEc2StackProps)
		p.StackProps = awscdk.StackProps{
			Env: env,
		}
	}
	return p	
}


func (s *KewlEc2Stack) New(scope constructs.Construct, id string, props *KewlEc2StackProps) *KewlEc2Stack {
	if s == nil {
		s = new(KewlEc2Stack)
	}

	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	s.stack = awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("AppQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return s
}