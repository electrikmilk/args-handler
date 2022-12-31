# args-parser

Simple arguments parser written in Go. Mainly here so I can easily use it in my other projects.

You can currently register options/flags, then check on if that flag was used and access it's value if applicable.

Flags follow the UNIX rules of having one dash for single letter versions of flags and double dashed versions of flags with whole words. (e.g. `-a` `--all`)

Argument values proceed the flag with a `=` sign separating (e.g. `-a=value` `--arg=value`)

Does not _yet_ support subcommands.
