package stacks

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	jsii "github.com/aws/jsii-runtime-go"
)

const (
 	nixOsAmi string = "ami-0d72ab697beab5ea5"
	awsAmi string = "ami-0d9858aa3c6322f73"
	ubuntuAmi string = "ami-085284d24fe829cd0"

	// TODO: Definitely make this a template
	nixosUserData string = `{ pkgs, ... }:

{
	imports = [ <nixpkgs/nixos/modules/virtualisation/amazon-image.nix> ];
	ec2.hvm = true;
	environment.systemPackages = with pkgs; [
		curl
		emacs
		git 
		nodejs
		nodePackages.npm
	];
}
`
	awsAmiUserData string = `
`
)

func (p *KewlEc2StackProps) New(env *awscdk.Environment) *KewlEc2StackProps {
	if p == nil {
		p = new(KewlEc2StackProps)
		p.StackProps = awscdk.StackProps{ Env: env }
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
	var region string
	region = *props.StackProps.Env.Region
	s.stack = awscdk.NewStack(scope, &id, &sprops)

	vpc := ec2.NewVpc(s.stack, props.VpcName, &ec2.VpcProps{})
	sg := ec2.NewSecurityGroup(
		s.stack, 
		props.SecurityGroupName, 
		&ec2.SecurityGroupProps{
			Vpc: vpc,
			AllowAllOutbound: jsii.Bool(true),
			Description: jsii.String("Security Group for access to hackday-sandbox-sg."),
			SecurityGroupName: props.SecurityGroupName,
			DisableInlineRules: jsii.Bool(true),
		})
	sg.AddIngressRule(ec2.Peer_AnyIpv4(), ec2.Port_Tcp(jsii.Number(22)), jsii.String("allow ssh access from the world"), jsii.Bool(false))
	keypair := ec2.NewCfnKeyPair(
		s.stack,
		props.KeyName,
		&ec2.CfnKeyPairProps{
			KeyName: props.KeyName,
		})
	linuxAmi := ec2.MachineImage_GenericLinux(
		&map[string]*string{
			region: jsii.String(ubuntuAmi),
		},
		&ec2.GenericLinuxImageProps{
			UserData: ec2.UserData_Custom(jsii.String(awsAmiUserData)),
		})
	instance := ec2.NewInstance(
		s.stack, 
		props.InstanceName,
		&ec2.InstanceProps{
			Vpc: vpc,
			InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C5, ec2.InstanceSize_XLARGE9),
			MachineImage: linuxAmi,
			InstanceName: props.InstanceName,
			SecurityGroup: sg,
			KeyName: props.KeyName,
			VpcSubnets: &ec2.SubnetSelection{
				SubnetType: ec2.SubnetType_PUBLIC,
			},
			BlockDevices: &[]*ec2.BlockDevice{
				&ec2.BlockDevice{
					DeviceName: jsii.String("/dev/sda1"),
					Volume: ec2.BlockDeviceVolume_Ebs(jsii.Number(100), &ec2.EbsDeviceOptions{}),
				},
			},
		})
	s.resources = make([]any, 0)
	s.resources = append(
		s.resources,
		vpc,
		sg,
		instance,
		keypair,
	)

	return s
}