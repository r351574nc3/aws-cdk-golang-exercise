# aws-cdk-golang-exercise
Me messing around with aws cdk and golang


## Building 

```shell
bazelisk build //:app
```

## Update repos/dependencies

```shell
bazelisk run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

## CDK Operations

### Deployment

```shell
cdk synth
cdk deploy --profile brave-lp --capabilities CAPABILITY_NAMED_IAM  KewlEc2Stack 
```

AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_DEFAULT_REGION