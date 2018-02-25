load("@k8s_deploy//:defaults.bzl", "k8s_deploy")

k8s_deploy(
    name = "cluster",
    template = ":deployment.yaml",
    images = {
        "gcr.io/stackb-151821/go-example-app:dev": "//go/example:server_image",
    },
)

k8s_deploy(
    name = "service",
    template = ":service.yaml",
)

# I ran the following BEFORE applying this:
# $ kubectl create clusterrolebinding my-cluster-admin-binding --clusterrole=cluster-admin --user=$(gcloud info --format="value(config.account)")
k8s_deploy(
    name = "ambassador-rbac",
    template = ":ambassador-rbac.yaml",
)

k8s_deploy(
    name = "ambassador-service",
    template = ":ambassador-service.yaml",
)

