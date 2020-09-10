# Superclog

Superclog is a tool to help generate useful release notes using conventional
commit syntax.

## How to use:

Superclog allows you to pass the following options

| Option     |      Parameters                                                   |
|----------  |-------------------------------------------------------------------|
| -path      | Path to git repository                                            |
| -from      | Git Hash from e.g. "d670460b4b4aece5915caf5c68d12f560a9fe3e4"     |
| -to        | Git hash to e.g. e.g. "d670460b4b4aece5915caf5c68d12f560a9fe3e4"  |
| -tmpl      | Template from built in template list e.g. "ExternalRelease"       |
| -tmpl-file | Template file path e.g. "./my-template.tmpl"                      |

The prebuilt binary comes with a bunch of built in templates. These are available in
`github.com/lwaddicor/superclog/templates` but can be referenced using short names.

### Built in templates
 - ExternalRelease
 - InternalQARelease

## Building and using


The standard CLI can be build using:

```
go build ./cmd/superclog
```

### Regnerate included templates

A standard go generate will rebuild the templates:

```
go generate
```