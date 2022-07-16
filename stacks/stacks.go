package stacks

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	jsii "github.com/aws/jsii-runtime-go"
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

	vpc := ec2.NewVpc(s.stack, jsii.String("hackday-sandbox-vpc"), &ec2.VpcProps{})
	sg := ec2.NewSecurityGroup(
		s.stack, 
		jsii.String("hackday-sandbox-sg"), 
		&ec2.SecurityGroupProps{
			Vpc: vpc,
			AllowAllOutbound: jsii.Bool(true),
			Description: jsii.String("Security Group for access to hackday-sandbox-sg."),
			SecurityGroupName: jsii.String("hackday-sandbox-sg"),
		})
	sg.addIngressRule(ec2.peer.anyIpv4(), ec2.port.tcp(jsii.Number(22)), jsii.String("allow ssh access from the world"))
	/*
	ingress := ec2.NewCfnSecurityGroupIngress(
		s.stack,
		jsii.String("hackday-sandbox-sg-ingress"),
		&ec2.CfnSecurityGroupIngressProps {
			IpProtocol: jsii.String("tcp"),
			CidrIp: jsii.String("0.0.0.0/0"),
			GroupId: jsii.String(*sg.SecurityGroupId()),
			FromPort: jsii.Number(-1),
			ToPort: jsii.Number(22),
		})
		*/
	s.resources = make([]any, 0)
	s.resources = append(
		s.resources,
		vpc,
		sg,
	)

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("AppQueue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return s
}