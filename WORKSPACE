load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
)
gazelle_revision = "39b19ce92c43c9915d69db5fa68b1772434df2e9"
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/archive/{}.zip".format(gazelle_revision)],
    strip_prefix = "bazel-gazelle-{}".format(gazelle_revision),
)
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains(go_version = "1.12.5")
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

go_repository(
    name = "com_github_golang_mock",
    build_extra_args = ["-exclude=vendor"],
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    commit = "d74b93584564161b2de771089ee697f07d8bd5b5",
    importpath = "github.com/golang/mock",
)

go_mock_revision = "e4a4d7dce76eedebd85d1bac72a9c02615c80eff"

http_archive(
    name = "com_github_jmhodges_bazel_gomock",
    strip_prefix = "bazel_gomock-{}".format(go_mock_revision),
    urls = ["https://github.com/jmhodges/bazel_gomock/archive/{}.zip".format(go_mock_revision)],
    sha256 = "44fa15a146e2be6eb545e432c5ba3d5cbe03967f3eb09b27607dc35863d13e06",
)
