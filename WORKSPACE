load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

rules_go_version = "v0.23.6"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8663604808d2738dc615a2c3eb70eba54a9a982089dd09f6ffe5d0e75771bc4f",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/{0}/rules_go-{0}.tar.gz".format(rules_go_version),
    ],
)

gazelle_revision = "v0.21.1"

http_archive(
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/{0}/bazel-gazelle-{0}.tar.gz".format(gazelle_revision),
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(nogo = "@//:default_nogo")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

http_file(
    name = "dep",
    executable = True,
    sha256 = "1a7bdb0d6c31ecba8b3fd213a1170adf707657123e89dff234871af9e0498be2",
    urls = ["https://github.com/golang/dep/releases/download/v0.5.0/dep-darwin-amd64"],
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
    sum = "h1:SP2mbHeyu1pCO5i8Xv6e3xgoKKMEhm6DJNVnflRaim0=",
    version = "v1.4.4-0.20200519145626-92f53b0a566c",
)

go_mock_version = "6f3bc610237ab39f288738d338a2d3ef6bf68bba"

http_archive(
    name = "com_github_jmhodges_bazel_gomock",
    patch_args = ["-p1"],
    patches = [
        "//patches:bazel-gomock-gopath.patch",  # https://github.com/jmhodges/bazel_gomock/pull/44
    ],
    strip_prefix = "bazel_gomock-{}".format(go_mock_version),
    urls = ["https://github.com/jmhodges/bazel_gomock/archive/{}.zip".format(go_mock_version)],
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

go_repository(
    name = "org_golang_x_mod",
    build_extra_args = ["-exclude=vendor"],
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/mod",
    sum = "h1:RM4zey1++hCTbCVQfnWeKs9/IEsaBLA8vTkd0WVtmH4=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_google_go_jsonnet",
    importpath = "github.com/google/go-jsonnet",
    patch_args = ["-p1"],
    patches = [
        "//patches:jsonnet-cpp.patch",  # https://github.com/google/go-jsonnet/pull/452
    ],
    sum = "h1:Nb4EEOp+rdeGGyB1rQ5eisgSAqrTnhf9ip+X6lzZbY0=",
    version = "v0.16.0",
)

load("@com_github_google_go_jsonnet//bazel:repositories.bzl", "jsonnet_go_repositories")

jsonnet_go_repositories()
