RUNFILES /home/pcj/github/pcj/docker-boulder-meetup-2018/bazel-bin/k8s/cluster.apply.runfiles
The following image references were not found: [gcr.io/stackb-151821/go-example-app:dev]
error: no objects passed to apply

kubectl --cluster="gke_stackb-151821_us-central1-a_prod" apply -f CONFIGURATION_DIR

load("@k8s_deploy//:defaults.bzl", "k8s_deploy")

k8s_deploy(
    name = "cluster",
    template = ":deployment.yaml",
    images = {
        "gcr.io/stackb-151821/go-example-app:dev": "//go/example:image",
    },
)

died with crashloopbackoff

kubectl run curl --image=radial/busyboxplus:curl -i --tty
If you don't see a command prompt, try pressing enter.
[ root@curl-6896d87888-n6n9l:/ ]$ nslookup kube-dns
Server:    10.39.240.10
Address 1: 10.39.240.10 kube-dns.kube-system.svc.cluster.local

nslookup: can't resolve 'kube-dns'

[ root@curl-6896d87888-n6n9l:/ ]$ nslookup example-service
Server:    10.39.240.10
Address 1: 10.39.240.10 kube-dns.kube-system.svc.cluster.local

Name:      example-service
Address 1: 10.39.252.229 example-service.default.svc.cluster.local
