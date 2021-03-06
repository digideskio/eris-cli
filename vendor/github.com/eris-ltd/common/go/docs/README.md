# Documentation Generator

This folder contains the auto-magic documentation generator.

## Installation

Three steps to documentation auto-magic.

### Step 1 - Copy in files

The files in this directory which end in `_` should be copied into the `docs/` folder of your repository (and the underscores should be removed).

### Step 2 - Update the copied in files

In the `generator.go_` file the lines which currently read as:

```go
var Description = "Package Management Tooling"
var RenderDir = fmt.Sprintf("./docs/documentation/pm/%s/", version.VERSION)
```

Should be updated to match the appropriate short description and short-code for the repository in question.

The `main()` function also will need to be updated to reference and build the top level cobra command for the repository.

### Step 3 - There is no step 3

## Usage

Documentation is currently broken into three major areas. These three areas are built and linked into a package of documentation for a specific tool. The package of documentation for the tool is then versioned and added to the site. The following areas are covered by the docs generator:

* cli commands. These are auto-generated from a go template kept in the docs_generator.go of this repository utilizing a cobra.Command struct.
* `specifications`. These are any specs that your tool utilizes. By default, any `*.md` files added to `./docs/specs` will be rendered as specifications (although this source folder can be tuned). Specifications will be put into a specifications folder within the versioning tree and will be linked to via the other documents generated. Specifications can leverage any of the additional go template functions covered below.
* `examples`. These should be lightweight mini-tutorials which cover different things that can be done with your tool. By default, any `*.md` files added to `./docs/examples` will be rendered as examples (again, this folder can be tuned). Examples will be put into an examples folder within the versioning tree and will be linked to via the other documents generated. Examples can leverage any of the additional go template functions covered below.

### Build

To build the documentation for your tool, have your CI system call `./docs/build.sh` to build **only** the versioned documentation, or `./docs/build.sh master` to build the versioned documentation as well as `latest` for your tool.

### Setup

Set up your file tree as described above with specs and examples in separate folders from one another. Name your `.md` files as you want them titled. So `assets_specification.md` will be rendered with a title of `Assets Specification` and so on. This titling leverages go's `strings.Title()` function.

Each of the files may leverage special go template functions that have been added to the template for specs and examples.

These are as follows:

* `title`: Title case a string (wrapper for strings.Title)
* `replace`: Replace a string (wrapper for strings.Replace)
* `chomp`: Trim trailing whitespace (wrapper for strings.TrimSpace)
* `handle_file`: Given a string, render its `.md` filename
* `handle_link`: Given a filename, render it's link
* `insert_definition`: Given a file which is within the `./definitions` directory and the name of a struct, this will output the entire struct and add the appropriate code fences (triple back ticks + language modifier) (note, both variables must be quoted).
* `insert_bash_lines`: Given a file which is within the `./docs/tests` directory and a set of line numbers in the form `"10-16"` or simply `"16"` this function will embed the appropriate lines with the appropriate code fences (again, quote everything here).
* `insert_file`: Given a file name which is within the `./docs/tests` directory, renders the entire file with the appropriate code fences and also download (`curl`) command.