package(default_visibility=["//visibility:public"])

filegroup(
    name = "files",
    srcs = [
        "BUILD",
    ],
)

py_library(
    name = "app",
    srcs = [
        "app.py",
    ],
)

py_binary(
    name = "main",
    srcs = [
        "main.py",
    ],
    deps = [
        ":app",
    ],
)

load("@io_bazel_rules_docker//python:image.bzl", "py_image")

py_image(
    name = "image",
    main = "main.py",
    srcs = [
        "main.py",
    ],
    deps = [
        ":app",
    ],
)
