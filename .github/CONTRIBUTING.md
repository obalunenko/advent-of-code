# Contributing

By participating to this project, you agree to
abide [code of conduct](https://github.com/obalunenko/advent-of-code/blob/master/.github/CODE_OF_CONDUCT.md).

Please note that I'll not accept any PR with new puzzle solutions - I don't like spoilers.

But if you have any idea how to improve the existing codebase - you are welcome to create a PR.

## Setup your machine

`advent-of-code` is written in [Go](https://golang.org/).

Prerequisites:

- [Go 1.17+](https://golang.org/doc/install)
- make

Clone `advent-of-code` anywhere:

```sh
git clone git@github.com:obalunenko/advent-of-code.git
```

`cd` into the directory and install the dependencies:

```sh
make install-tools
```

A good way of making sure everything is all right is running the test suite:

```sh
make test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

```sh
make build
```

When you are satisfied with the changes, we suggest you run:

```sh
make test && make test-regression
```

Before you commit the changes, we also suggest you run:

```sh
make format-project
```

## Create a commit

Commit messages should be well formatted, and to make that "standardized", we
are using Conventional Commits.

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Submit a pull request

Push your branch to your `advent-of-code` fork and open a pull request against the main branch.