# Upstream base

The Czech translation (`examples/`, `examples.txt`, `templates/`, from the
`cz-translation` branch) is based on https://github.com/mmcgrana/gobyexample at:

    46fe6bc
    2019-06-13  Merge pull request #247 from C-Ross/fixEnvarReference

The build tooling was later merged with upstream up to:

    65c2883a9f041ca1db718f6cfdcab1042950ef00
    2019-09-06  Merge pull request #274 from mmcgrana/consistent-import

so the topics below were derived from `65c2883`. Three topics were added
upstream between `46fe6bc` and `65c2883` and are also missing from the
translation: XML (`xml`), Directories (`directories`) and Temporary Files and
Directories (`temporary-files-and-directories`).

Example directories, file names and page URLs use the upstream (English) IDs,
so `examples/<id>` here corresponds to `examples/<id>` upstream; only the text
is translated. `examples.txt` lists each example as `<id>: <Czech title>`.

Compared against upstream `master` at `28c81a7` (checked 2026-10-24).

## Topics added upstream that are missing here

In upstream `examples.txt` order (date the example was added upstream):

| Topic | Upstream dir | Added |
|---|---|---|
| Range over Built-in Types | `range-over-built-in-types` | 2024-08 (renamed from *Range*) |
| Strings and Runes | `strings-and-runes` | 2022-02 |
| Enums | `enums` | 2024-05 |
| Struct Embedding | `struct-embedding` | 2022-05 |
| Generics | `generics` | 2022-03 |
| Range over Iterators | `range-over-iterators` | 2024-08 |
| Custom Errors | `custom-errors` | 2024-03 |
| Recover | `recover` | 2021-09 |
| Text Templates | `text-templates` | 2022-02 |
| SHA256 Hashes | `sha256-hashes` | 2022-03 (replaces *SHA1 Hashes*) |
| Embed Directive | `embed-directive` | 2022-05 |
| Testing and Benchmarking | `testing-and-benchmarking` | 2021-12 |
| Logging | `logging` | 2023-08 |
| HTTP Client | `http-client` | 2022-09 (renamed from *HTTP Clients*) |
| HTTP Server | `http-server` | 2022-09 (renamed from *HTTP Servers*) |
| TCP Server | `tcp-server` | 2025-11 |
| Context | `context` | 2020-01 |

## Topics here that upstream removed or renamed

- Range → Range over Built-in Types
- Collection Functions → removed (superseded by generics / `slices` package)
- SHA1 Hashes → SHA256 Hashes
- HTTP Clients → HTTP Client
- HTTP Servers → HTTP Server

## Existing topics significantly updated upstream

Changed lines in the `.go` sources since the fork point (largest first):
mutexes (111), errors (98), sorting-by-functions (66), random-numbers (49),
string-formatting (46), signals (45), spawning-processes (40), slices (35),
structs (32), atomic-counters (32), json (29), waitgroups (28), maps (27),
reading-files (26), arrays (26), epoch (24), sorting (21), recursion (21),
directories (21); ~35 more with smaller edits.

To see the upstream changes for one example:

    git remote add upstream https://github.com/mmcgrana/gobyexample.git
    git fetch upstream
    git diff 65c2883 upstream/master -- examples/mutexes
