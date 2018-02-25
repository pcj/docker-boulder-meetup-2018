load("@io_bazel_rules_go//go:def.bzl", "go_prefix")
load("@io_bazel_rules_docker//container:image.bzl", "container_image")

go_prefix("github.com/pcj/docker-boulder-meetup-2018")

load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar", "pkg_deb")

pkg_tar(
    name = "show_tar",
    srcs = [
        ":show",
        "@org_golang_x_tools//cmd/present",
        #"@org_golang_x_tools_cmd_present//:templates",
        #"@org_golang_x_tools_cmd_present//:static",
        ":slides",
        "//static",
        "//templates",
        "//images",
        "BUILD",
        "//java/example:files",
    ],
)

sh_binary(
    name = "show",
    srcs = [
        "show.sh",
    ],
    data = [
        "@org_golang_x_tools//cmd/present",
        #"@org_golang_x_tools_cmd_present//:templates",
        #"@org_golang_x_tools_cmd_present//:static",
        ":slides",
        "//static",
        "//templates",
        "//images",
        "BUILD",
        "//java/example:files",
    ],
    visibility = ["//visibility:public"],    
)

container_image(
    name = "container",
    base = "@busybox//image",
    files = [
        ":show",
        "@org_golang_x_tools//cmd/present",
        #"@org_golang_x_tools_cmd_present//:templates",
        #"@org_golang_x_tools_cmd_present//:static",
        ":slides",
        "//static",
        "//templates",
        "//images",
        "BUILD",
        "//java/example:files",
    ],
    cmd = "./show",
)

filegroup(
    name = "slides",
    srcs = [
        "bazel.slide",
        "docker.slide",
        "kubernetes.slide",
    ],
)

