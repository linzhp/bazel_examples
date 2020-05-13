load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

rules_go_version = "v0.23.0"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6a68e269802911fa419abb940c850734086869d7fe9bc8e12aaf60a09641c818",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/{0}/rules_go-{0}.tar.gz".format(rules_go_version),
    ],
)

gazelle_revision = "v0.21.0"

http_archive(
    name = "bazel_gazelle",
    sha256 = "bfd86b3cbe855d6c16c6fce60d76bd51f5c8dbc9cfcaef7a2bb5c1aafd0710e8",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/{0}/bazel-gazelle-{0}.tar.gz".format(gazelle_revision),
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

go_repository(
    name = "com_github_chai2010_webp",
    build_extra_args = ["-exclude=vendor"],
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "github.com/chai2010/webp",
    sum = "h1:4Ei0/BRroMF9FaXDG2e4OxwFcuW2vcXd+A6tyqTJUQQ=",
    version = "v1.1.0",
)

