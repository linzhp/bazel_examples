load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

rules_go_version = "745de507d7f7e8837eba01dd5ff9d4ef7edf31c7"

http_archive(
    name = "io_bazel_rules_go",
    strip_prefix = "rules_go-{}".format(rules_go_version),
    urls = [
        "https://github.com/bazelbuild/rules_go/archive/{}.zip".format(rules_go_version),
    ],
)

gazelle_revision = "7e93b118efcdfe63f43057a88711762e0b5d7267"

http_archive(
    name = "bazel_gazelle",
    strip_prefix = "bazel-gazelle-{}".format(gazelle_revision),
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/archive/{}.zip".format(gazelle_revision),
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains(nogo="@//:default_nogo")

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
    importpath = "github.com/mattn/go-sqlite3",
    sum = "h1:gXHsfypPkaMZrKbD5209QV9jbUTJKjyR5WD3HYQSd+U=",
    version = "v2.0.3+incompatible",
)

go_repository(
    name = "com_github_golang_mock",
    build_extra_args = ["-exclude=vendor"],
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "github.com/golang/mock",
    sum = "h1:Rd1kQnQu0Hq3qvJppYSG0HtP+f5LPPUiDswTLiEegLg=",
    version = "v1.4.0",
)

go_mock_version = "1.2"

http_archive(
    name = "com_github_jmhodges_bazel_gomock",
    sha256 = "850394d47c5331898728c02be8eb60e7ac580f09559fcb87c790f6db12dd3338",
    strip_prefix = "bazel_gomock-{}".format(go_mock_version),
    urls = ["https://github.com/jmhodges/bazel_gomock/archive/v{}.zip".format(go_mock_version)],
)

go_repository(
    name = "com_github_vektra_mockery",
    importpath = "github.com/vektra/mockery",
    sum = "h1:Xim2mBRFdXzXmKRO8DJg/FJtn/8Fj9NOEpO6+WuMPmk=",
    version = "v0.0.0-20181123154057-e78b021dcbb5",
)
