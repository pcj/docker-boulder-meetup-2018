load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test", "go_binary")

DOTFILES = [
    "rule",
    "rules",
    "constructor",
    "dag",
]

filegroup(
    name = "images",
    srcs = glob([
        "*.jpg",
        "*.svg",
    ]) + [name + ".svg" for name in DOTFILES],
    visibility = ["//visibility:public"],
)

[genrule(
    name = name,
    srcs = ["%s.dot" % name],
    outs = ["%s.svg" % name],
    cmd = "dot $(SRCS) -Tsvg -o$@",
) for name in DOTFILES]
    
