load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

rules_go_version = "v0.24.4"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "ac03931e56c3b229c145f1a8b2a2ad3e8d8f1af57e43ef28a26123362a1e3c7e",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/{0}/rules_go-{0}.tar.gz".format(rules_go_version),
    ],
)

gazelle_revision = "v0.22.0"

http_archive(
    name = "bazel_gazelle",
    sha256 = "d4113967ab451dd4d2d767c3ca5f927fec4b30f3b2c6f8135a2033b9c05a5687",
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
    sum = "h1:LDdKkqtYlom37fkvqs8rMPFKAMe8+SgjbwZ6ex1/A/Q=",
    version = "v1.11.0",
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

http_archive(
    name = "icu4c",
    strip_prefix = "icu",
    sha256 = "627d5d8478e6d96fc8c90fed4851239079a561a6a8b9e48b0892f24e82d31d6c",
    urls = [
        "https://github.com/unicode-org/icu/releases/download/release-64-2/icu4c-64_2-src.tgz",
    ],
    build_file = "//third_party/icu:BUILD.system",
)

