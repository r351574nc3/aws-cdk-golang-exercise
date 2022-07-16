load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/r351574nc3/aws-cdk-golang-excercise",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

genrule(
    name = "app",
    srcs = [
        "//:cli",
    ],
    outs = [
        "infra",
    ],
    cmd = "cp ./bazel-out/*/bin/*/cli \"$@\"",
    executable = 1,
    local = 1,
)

go_library(
    name = "go_default_library",
    srcs = ["app.go"],
    importpath = "github.com/r351574nc3/aws-cdk-golang-excercise",
    visibility = ["//visibility:private"],
    deps = [
        "//stacks:go_default_library",
        "@com_github_aws_aws_cdk_go_awscdk_v2//:go_default_library",
        "@com_github_aws_constructs_go_constructs_v10//:go_default_library",
        "@com_github_aws_jsii_runtime_go//:go_default_library",
    ],
)

go_binary(
    name = "cli",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["app_test.go"],
    embed = [":go_default_library"],
)
