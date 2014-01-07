
# GOLANG project template

This is my personal [GOLANG](http://golang.org/) project template.
It is built with these goals in mind:

* Everything should be local to the development folder as much as is possible without
  resorting to ugly hacks.

* 1-step build or test without any special setup after any `git clone` on a newly minted
  machine except for installing the `go` compiler itself.

* Do not rely on global `GOPATH` and yet still allows you to check your entire source
  folder in as if you would a normal go program inside one. This makes your repository
  plays well with other go coders who will surely be using the global `GOPATH` convention.

# Getting Started

Run this one-liner to dump the repository content into current directory.

```sh
curl -L https://github.com/chakrit/go-scratch/archive/master.tar.gz | tar -xzv --strip 1
```

Or a full setup for a new project:

```sh
mkdir your-new-shiny-project    # Make a new shiny folder for your new project.
cd your-new-shiny-project       # Get into the folder first, of course.
git init                        # Or not, doesn't matter.

# Download the scratch template
curl -L https://github.com/chakrit/go-scratch/archive/master.tar.gz | tar -xzv --strip 1

git add .
git commit -m "Initialize GOLANG project. (github.com/chakrit/go-scratch)"
```

# Makefile

Everything is done through the `Makefile` for convenience. A wrapper script `./go` is also
provided that invokes `go` with `GOPATH` sets to the local `.go` folder.

Makefile tasks defaults to `test`. The `all` task is simply an alias for the `build`
task. All common tasks you'd do with `go` is given the same name in the Makefile so
`go vet` is, through the Makefile, `make vet` or via the wrapper script `./go vet`.

# Dependencies

Dependencies are handled implicitly and automatically as you run tasks that involve import
paths thanks to the powerful `go get` command.

Specifically, `make deps` and `make test-deps` will download the dependencies into place
and subsequent `make test` or `make build` will automaticaly compiles the downloaded
dependencies as needed.

The initial `main.go` file provided with this project contains some dependencies as well
as tests (and test dependencies) to demonstrate this.

# Subpackages

This template supports multiple packages development inside a single folder out of the
box. You are, however, required to list all the subpackage paths inside the Makefile `PKG`
variable as automatically detecting them is tricky and error prone.

For example, if you have a `context` subpackage, edit the `PKG` line like so:

```make
PKG := . ./context
```

... or if you have nothing in the root folder and only subpackages `uno` `dos` and `tres`:

```make
PKG := ./uno ./dos ./tres
```

# Example

Here's a sample output of what happens when you simply cloned this repository and issue
`make`:

```sh
$ make
/Users/chakrit/Documents/go-scratch/go get -d -t .
/Users/chakrit/Documents/go-scratch/go test .
ok    _/Users/chakrit/Documents/go-scratch0.010s
```

# Silent run

Uses `make`'s `-s` flag to prevent `make` from echoing commands (useful for reducing
clutter to standard output.)

```sh
$ make -s
ok      _/Users/chakrit/Documents/go-scratch    0.011s
```

# LICENSE

[WTFPL](http://www.wtfpl.net/)

# TODO:

* Package the curl installation into a script.

