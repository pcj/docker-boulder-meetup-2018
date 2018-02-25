load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test", "go_binary")

filegroup(
    name = "static",
    srcs = glob([
        "*.css",
        "*.js",
        "*.html",
        "favicon.ico",
        "example.*",
    ]),
    visibility = ["//visibility:public"],
)
