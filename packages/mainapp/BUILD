load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_test")
go_binary(
    name = "mainapp",
    srcs = ["main.go"],
    importpath = "monorepo/packages/main",
    deps = [
        "//packages/mainapp/router:router"
    ],
)

