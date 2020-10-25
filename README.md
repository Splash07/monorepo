# Example of managing a monorepo with Bazel

## *Install gazelle*
>go get github.com/bazelbuild/bazel-gazelle/cmd/gazelle
Gazelle is a Bazel build file generator for Bazel projects

## *Init go.mod*
> go mod init monorepo

## *Add WORKSPACE file, which defines bazel imports and context for the BUILD file*

## *Add BUILD file*
> BUILD file is equivalent to Makefile

## *Add codes following a monorepo structure*
#gazelle:prefix part tells gazelle that we are operating in a golang module and import path starts with the same prefix as golang module

## *Run gazelle to generate the BUILD.bazel files for all of the packages*
> bazel run //:gazelle

## *Associate go.mod with bazel*
> gazelle update-repos --from_file=go.mod -to_macro=go_third_party.bzl%go_deps

A go_third_party.bzl will be generated to have "bazel-ish" representation of go.mod. That way, bazel will be aware of dependencies mentioned in go.mod

## *Bazel build all*
> bazel build //...

## *To build mainapp service in this repo*
> bazel build //packages/mainapp

## *To run mainapp service in this repo*
> bazel run //packages/mainapp

## *To generate dependency graph*
> bazel query 'allpaths(packages/...,//packages/shared/handlers/health:health)' --output graph | dot -Tpng > dep.png


# TO-DO: Building docker images with bazel

# Notes:
- Build files have two kinds of syntaxes: load() and macro
  - Load imports other bazel module
  - Macros are functions of the imported modules using load and can be called in build files
- After bazel compiles, there should be every one BUILD field in each of the subfolders of packages
