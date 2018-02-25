package(default_visibility=["//visibility:public"])

load("@io_bazel_rules_docker//java:image.bzl", "java_image")

filegroup(
    name = "files",
    srcs = [
        "BUILD",
    ],
)

java_library(
    name = "app",
    srcs = [
        "App.java",
    ],
)

java_binary(
    name = "main",
    main_class = "example.App",
    runtime_deps = [
        ":app",
    ],
)

java_image(
    name = "image",
    main_class = "example.App",
    runtime_deps = [
        ":app",
    ],
)
