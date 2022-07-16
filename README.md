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
cdk deploy --all
```