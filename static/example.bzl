def _markdown_library_impl(ctx):
    outs = []
    for f in ctx.files.srcs:
        outfile = ctx.action.declare_file(f.basename + '.html')
        outs.append(outfile)
    ctx.action.run(
        outputs = outs,
        inputs = ctx.files.srcs,
        executable = ctx.executable.compiler,
        args = [f.path for f in ctx.files.srcs],
    )
    return struct(files=outs)
# END

markdown_library = rule(
    implementation = _markdown_library_impl,
    attrs = {
        "srcs": attr.label_list(mandatory=True),
        "smartypants": attr.bool(),
        "compiler": attr.label("@com_github_russross_blackfriday//:compiler", executable=True),
    }
)
