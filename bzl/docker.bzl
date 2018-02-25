
def _execute(rtx, cmds):
    """Execute a command and fail if return code.
    Args:
      rtx: !repository_ctx
      cmds: !list<string>
    Returns: struct value from the rtx.execute method.
    """
    #print("Execute <%s>" % " ".join(cmds))
    result = rtx.execute(cmds)
    if result.return_code:
        fail(" ".join(cmds) + "failed: %s" %(result.stderr))
    return result

def _docker_distribution_repositories_impl(rtx):
    url = "https://github.com/docker/distribution/archive/%s.zip" % rtx.attr.commit
    rtx.download_and_extract(url, stripPrefix="distribution-"+rtx.attr.commit)

    result = _execute(rtx, ["cat", "vendor.conf"])
    deps = result.stdout.split("\n")
    lines = [
        'load("@io_bazel_rules_go//go:def.bzl", "go_repository")',
    ]

    lines.append("def docker_distribution_dependencies():")
    for line in deps:
        print("dep: " + line)
        tokens = line.split(" ")
        repo = tokens[0]
        commit = tokens[1]
        name = repo.replace("/", "_").replace(".", "_").replace('-', '_').lower()
        lines.append("    go_repository(")
        lines.append("        name = '%s'," % name)
        lines.append("        commit = '%s'," % commit)
        lines.append("        remote = 'https://%s.git'," % repo)
        lines.append("        importpath = '%s'," % repo)
        lines.append("    )")
        
    rtx.file("deps.bzl", "\n".join(lines))
        
    rtx.file("BUILD.bazel", "")

docker_distribution_repositories = repository_rule(
    implementation = _docker_distribution_repositories_impl,
    attrs = {
        "commit": attr.string(
            mandatory = True,
        ),
    },
)
    
