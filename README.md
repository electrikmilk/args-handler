# args-parser

<p>
        <a href="https://github.com/electrikmilk/args-parser/actions/workflows/go.yml"><img src="https://github.com/electrikmilk/args-parser/actions/workflows/go.yml/badge.svg?branch=main" alt="Build & Test"></a>
    <a href="https://pkg.go.dev/github.com/electrikmilk/args-parser?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/electrikmilk/args-parser"><img src="https://goreportcard.com/badge/github.com/electrikmilk/args-parser"/></a>
</p>

A simple arguments parser written in Go. Mainly here so I can easily use it in my other projects.

### Register arguments

You can currently register options/flags, then check on if that flag was used and access its value if applicable.

```go
args.Register(Argument{
        name: "arg",
        description: "My first argument",
})
```

### Auto-generated usage information

```go
args.PrintUsage()
```

### Usage

Flags follow the UNIX rules of having one dash for single-letter versions of flags and double-dashed versions of flags with whole words. (e.g. `-a` `--all`). It doesn't technically matter though since it just trims dashes from the beginning of the argument.

Argument values proceed the flag with a `=` sign separating (e.g. `-a=value` `--arg=value`).

Then either check if the flag is being used or get its value.

```go
args.Using("arg") // bool

args.Value("arg") // string
```

---

Does not _yet_ support subcommands.
