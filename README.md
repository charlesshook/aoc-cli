# aoc-cli

This project was built to make it easy to grab the AOC prompt and input file. It was also built to make it easy to setup your projects for the AOC. I wanted an easy way to switch the language I was using during AOC that would set up some base files that includes tests.

## Features
- Download Advent of Code prompt into a markdown file.
- Download the input file for any given day.
- Structures your Advent of Code solutions.
- System to setup projects for different languages.

## Current Languages Supported
- [x] [Go](https://go.dev/) <br>

If there is a language that is not supported that you would like to exist, please create a issue!

## Installation
### Compile from source
Install a recent version of [Go](https://go.dev/) <br>
Then to build:
```
go build main.go
```

### Insall with go get
Install a recent version of [Go](https://go.dev/) <br>
Then:
```
go get github.com/charlesshook/aoc-cli@latest
```

## Session cookie
Every user of Advent of Code gets different puzzle input files. That is why the session cookie is needed. This is to ensure you get the input file that belongs to you. There are many ways to get your session cookie. The best way is to Google how to get it based on watch browser you are using because the method may vary.

## Puzzle Project Structure
This cli will setup a specific folder structure for Advent of Code. Currently there is no way to change this structure. This is what it looks like:
```
2015/
|___1/
|___2/
|___3/
|   |___go/
|       |   main.go
|       |   test.go
|
2016
|___1/
|___2/
|   |___python/
|       |   main.py
|       |   test.py
|       go/
|       |   main.go
|       |   test.go
```

## Usage
```
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  aoc-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  input       Gets input file.
  prompt      Gets AOC prompt.
  setup       Sets up a project for a given day of AOC.

Flags:
  -h, --help     help for aoc-cli
  -t, --toggle   Help message for toggle

Use "aoc-cli [command] --help" for more information about a command.
```

```
Sets up a project for a given day of AOC.

Usage:
  aoc-cli setup [flags]

Flags:
  -d, --day int           Day to create the project for. (default 27)
  -h, --help              help for setup
  -l, --language string   The language to create the project for (default "go")
  -y, --year int          Year to create the project for. (default 2024)
```