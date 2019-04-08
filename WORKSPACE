load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.2/rules_go-0.18.2.tar.gz"],
    sha256 = "31f959ecf3687f6e0bb9d01e1e7a7153367ecd82816c9c0ae149cd0e5a92bf8c",
)
gazelle_revision = "c8a9055d89a3678b14cfa300a0c1b860a52d0197"
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/archive/{}.zip".format(gazelle_revision)],
    strip_prefix = "bazel-gazelle-{}".format(gazelle_revision),
)
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

http_file(
    name = "dep",
    urls = ["https://github.com/golang/dep/releases/download/v0.5.0/dep-darwin-amd64"],
    sha256 = "1a7bdb0d6c31ecba8b3fd213a1170adf707657123e89dff234871af9e0498be2",
    executable = True,
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    commit = "25ecb14adfc7543176f7d85291ec7dba82c6f7e4",
    importpath = "github.com/mattn/go-sqlite3",
)