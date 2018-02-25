load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "registry",
    srcs = [
        "registry.go",
    ],
    deps = [
        "@com_github_gorilla_mux//:go_default_library",
    ],
    importpath = "github.com/pcj/docker-boulder-meetup-2018/go/registry",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "server",
    srcs = [
        "server.go",
    ],
    deps = [
        ":registry",
        "@com_github_gorilla_mux//:go_default_library",
    ],
    visibility = ["//visibility:public"],
)
