# args-parser

<p>
    <a href="https://pkg.go.dev/github.com/electrikmilk/args-parser?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/electrikmilk/args-parser"><img src="https://goreportcard.com/badge/github.com/electrikmilk/args-parser"/></a>
</p>

Simple arguments parser written in Go. Mainly here so I can easily use it in my other projects.

You can currently register options/flags, then check on if that flag was used and access it's value if applicable.

Flags follow the UNIX rules of having one dash for single letter versions of flags and double dashed versions of flags with whole words. (e.g. `-a` `--all`)

Argument values proceed the flag with a `=` sign separating (e.g. `-a=value` `--arg=value`)

Does not _yet_ support subcommands.
