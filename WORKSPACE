PRESENT_TEMPLATES_BUILD_CONTENT="""
filegroup(
    name = "templates",
    srcs = glob(["templates/**"]),
    visibility = ["//visibility:public"],
)
filegroup(
    name = "static",
    srcs = glob(["static/**"]),
    visibility = ["//visibility:public"],
)
"""

RULES_GO_COMMIT="3e6e9bbf8f9ed35f5956fba80dde9283e121a66c"

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/archive/%s.zip" % RULES_GO_COMMIT, # PR #1178
    strip_prefix = "rules_go-" + RULES_GO_COMMIT,
    sha256 = "0710cf2945ebea31b519d868cc1610002a8ca99595d30b2c886bf5271132efa2",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")

# Override grpc for 1.9.0
go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    tag = "v1.9.0",
)

go_rules_dependencies()

go_register_toolchains()

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    commit = "9b61fcc4c548d69663d915801fc4b42a43b6cd9c",
)

new_http_archive(
    name = "org_golang_x_tools_cmd_present",
    url = "https://github.com/golang/tools/archive/2ae76fd1560b622911f444c1e66b70a857e1de67.zip", 
    strip_prefix = "tools-2ae76fd1560b622911f444c1e66b70a857e1de67/cmd/present",
    sha256 = "b50c19119577b75bcb871e13b3a59abafe2054a363ff7c100972b9e57b66a7db",
    build_file_content = PRESENT_TEMPLATES_BUILD_CONTENT,
)

git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    commit = "198367210c55fba5dded22274adde1a289801dc4",
)

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
    container_repositories = "repositories",
)

# This is NOT needed when going through the language lang_image
# "repositories" function(s).
container_repositories()

container_pull(
    name = "busybox",
    registry = "index.docker.io",
    repository = "library/busybox",
    # 'tag' is also supported, but digest is encouraged for reproducibility.
    #digest = "sha256:b50c19119577b75bcb871e13b3a59abafe2054a363ff7c100972b9e57b66a7db",
)

load(
    "@io_bazel_rules_docker//java:image.bzl",
    _java_image_repos = "repositories",
)

_java_image_repos()


load(
    "@io_bazel_rules_docker//python:image.bzl",
    _python_image_repos = "repositories",
)

_python_image_repos()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()



go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    commit = "c0091a029979286890368b4c7b301261e448e242",
)

go_repository(
    name = "com_github_gorilla_context",
    importpath = "github.com/gorilla/context",
    commit = "08b5f424b9271eedf6f9f0ce86cb9396ed337a42",
)


git_repository(
    name = "io_bazel_rules_k8s",
    commit = "fb4e8367e90c1bae45a3483a1100bc7b2b8fceae",
    remote = "https://github.com/bazelbuild/rules_k8s.git",
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories")

k8s_repositories()

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_defaults")

k8s_defaults(
  # This becomes the name of the @repository and the rule
  # you will import in your BUILD files.
  name = "k8s_deploy",
  kind = "deployment",
  # This is the name of the cluster as it appears in:
  #   kubectl config current-context
  cluster = "gke_stackb-151821_us-central1-a_prod",
)

DOCKER_DISTRIBUTION_COMMIT = "6664ec703991875e14419ff4960921cce7678bab"

go_repository(
    name = "com_github_docker_distribution",
    importpath = "github.com/docker/distribution",
    commit = DOCKER_DISTRIBUTION_COMMIT,
)

load(
    "//bzl:docker.bzl",
    "docker_distribution_repositories",
)

docker_distribution_repositories(
    name = "docker_distribution",
    commit = DOCKER_DISTRIBUTION_COMMIT,
)

load(
    "@docker_distribution//:deps.bzl",
    "docker_distribution_dependencies",
)

docker_distribution_dependencies()
