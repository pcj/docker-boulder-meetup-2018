<table><tr>
<td><img src="https://bazel.build/images/bazel-icon.svg" height="120"/></td>
<td><img src="https://raw.githubusercontent.com/docker-library/docs/471fa6e4cb58062ccbf91afc111980f9c7004981/swarm/logo.png" height="120"/></td>
<td><img src="https://d33wubrfki0l68.cloudfront.net/1567471e7c58dc9b7d9c65dcd54e60cbf5870daa/a2249/images/flower.png" height="120"/></td>
</tr><tr>
<td>Bazel</td>
<td>Docker</td>
<td>Kubernetes</td>
</tr></table>

# Building Fast, Reproducible Docker Images with Bazel

[Bazel](https://bazel.build) is the open-source version of Google's
internal multi-language build tool called "Blaze".  Bazel has some
attractive features that make builds fast and scalable.

One of bazel's "secret" (relatively unknown) superpowers is building
docker images, written by the same team that developed the Google
Container Registry <https://gcr.io>.

In this talk I'll introduce Bazel and show how you can use it in your
docker/kubernetes workflow to enable fast and reproducible
builds/deployments.

# Bio

Paul Johnston is a software developer at Windward Solutions in
Boulder, CO.  He is an active member of the Bazel open source
community and an advocate for less sucky DevOps experiences.
