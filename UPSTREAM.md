# Upstream base

The Czech translation (`examples/`, `examples.txt`, `templates/`) follows
https://github.com/mmcgrana/gobyexample. The example content is in sync with
upstream `master` at:

    28c81a7
    2026-09-07  Clarify mutability of structs

(The original translation from the `cz-translation` branch was based on
upstream `46fe6bc`, 2019-06-13; all upstream example changes between
`46fe6bc` and `28c81a7` were ported and the new topics translated.)

Example directories, file names and page URLs use the upstream (English) IDs,
so `examples/<id>` here corresponds to `examples/<id>` upstream; only the text
is translated. `examples.txt` lists each example as `<id>: <Czech title>`, in
upstream order.

Only the example content was synced. The site tooling and design (upstream
moved from Pygments to Chroma, restyled `site.css`, split the templates) were
not ported; `tools/generate.go` and `tools/measure.go` only got the changes the
new examples need.

The examples need Go 1.27 (`encoding/json/v2`, `errors.AsType`, `new(expr)`);
`go.mod` says `go 1.27.1`, so an older `go` downloads that toolchain
automatically.

## Syncing again

To see what changed upstream since the last sync:

    git remote add upstream https://github.com/mmcgrana/gobyexample.git
    git fetch upstream
    git diff 28c81a7 upstream/master -- examples.txt examples/

Port each change into the Czech files by hand (comments and printed text are
translated, and some examples translate identifiers too), then update the
commit above.
