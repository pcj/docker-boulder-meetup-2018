load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")


go_library(
    name = "app1",
    srcs = [
        "app.go",
    ],
    importpath = "github.com/pcj/docker-boulder-meetup-2018/go/example",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "main",
    srcs = [
        "app.go",
    ],
    visibility = ["//visibility:public"],
)

load("@io_bazel_rules_docker//go:image.bzl", "go_image")

go_image(
    name = "image",
    srcs = [
        "app.go",
    ],
    visibility = ["//visibility:public"],
)


go_image(
    name = "server_image",
    srcs = [
        "server.go",
    ],
    visibility = ["//visibility:public"],
)
