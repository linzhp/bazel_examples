load("@io_bazel_rules_go//go:def.bzl", "go_rule")

def _heatpipe_impl(ctx):
    ctx.actions.write(ctx.outputs.types_out, "types1")
    ctx.actions.write(ctx.outputs.fx_out, "fx1")

    return [
        DefaultInfo(files = depset([ctx.outputs.types_out, ctx.outputs.fx_out])),
    ]

_heatpipe_rule = go_rule(
    _heatpipe_impl,
    attrs = {
        "topic": attr.string(
            doc = "The topic for the schema",
            mandatory = True,
        ),
        "types_out": attr.output(
            doc = "The __types.go file to be generated",
            mandatory = True,
        ),
        "fx_out": attr.output(
            doc = "The __fx.go file to be generated",
            mandatory = True,
        ),
    },
)

def heatpipe(name, topic, types_out = "types.go", fx_out = "fx.go"):
    _heatpipe_rule(
        name = name,
        topic = topic,
        types_out = types_out,
        fx_out = fx_out,
    )