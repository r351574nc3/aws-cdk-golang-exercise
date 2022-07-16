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
        "//cmd:cli",
    ],
    outs = [
        "infra",
    ],
    cmd = "cp ./bazel-out/*/bin/cmd/*/cli \"$@\"",
    executable = 1,
    local = 1,
)
