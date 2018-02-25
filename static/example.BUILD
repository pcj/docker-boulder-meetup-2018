java_library(
    name = "app",
    srcs = [
        "App.java",
    ],
)

go_library(
    name = "app",
    srcs = [
        "app.go",
    ],
)

go_library(
    name = "app",
    srcs = [
        "app.go",
    ],
)

go_repository(
    name = "com_github_docker_distribution",
    importpath = "github.com/docker/distribution",
    commit = "6664ec703991875e14419ff4960921cce7678bab",
)

maven_jar(
    name = "com_google_guava_guava_21",
    artifact = "com.google.guava:guava:jar:21.0",
    sha1 = "3a3d111be1be1b745edfa7d91678a12d7ed38709",
)

container_pull(
    name = "nginx",
    registry = "index.docker.io",
    repository = "library/nginx",
    digest = "da7790e885efbd08f95102f62247217dd54b22be538804acff04643ae9b38826",
    tag = "latest",
)
