# anstyle for Go

🦀 Rust anstyle project ported to Go

<table align=center><td>

```go

```

</table>

<p align=center>
  <a href="https://docs.rs/anstyle">Docs</a>
  | <a href="https://github.com/rust-cli/anstyle">Original rust-cli/anstyle</a>
</p>

- **[anstyle](anstyle):** 🎨 ANSI color `struct` and `interface`
- **[anstylegit](anstylegit):** 🔸 Parse Git color config to `Style`
<!-- - **[anstream](anstream):** 🌊 Helpers for ANSI output streams
- **[anstylelossy](anstylelossy):** 🗑️ Lossy color conversions
- **[anstylels](anstylels):** 📁 Parse `LS_COLORS` into a `Style` config
- **[anstyleparse](anstyleparse):** 🔠 Parse ANSI style escapes
- **[anstyleroff](anstyleroff):** 📄 Convert ANSI to ROFF
- **[anstylesvg](anstylesvg):** 🖼️ Convert ANSI to SVG
- **[anstylewincon](anstylewincon):** 🟦 Style legacy Windows terminals
- **[colorchoiceclap](colorchoiceclap):** 👏 Automatic `--color` flag for Clap -->

## Installation

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)

You're reading the root monorepo readme. 😉 Click the links above 👆 to check out subproject-specific installation instructions.

<sup>Spoiler: it's `go get github.com/jcbhmr/go-anstyle/<subproject>`</sup>

## Usage

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)

```go
package main

import (
    "github.com/jcbhmr/go-anstyle/anstyle"
    "github.com/jcbhmr/go-anstyle/anstylels"
    "github.com/jcbhmr/go-anstyle/anstylesvg"
)

func main() {
    // TODO
}
```

**Alternatives:**

- [github.com/fatih/color](https://github.com/fatih/color) ⭐7.1k

## Development

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)

Each of the Rust crates from the [rust-cli/anstyle](https://github.com/rust-cli/anstyle) repository are ported to Go in this monorepo as separate Go modules, each with their own `go.mod` file and their own version. Go monorepo versioning works if the Git tag starts with a `<subfolder>/` prefix like `anstyle/v0.1.0`.

- [ ] anstream
- [ ] anstyle-ansi-term
- [ ] anstyle-crossterm
- [x] anstyle-git
- [ ] anstyle-lossy
- [ ] anstyle-ls
- [ ] anstyle-owo-colors
- [ ] anstyle-parse
- [ ] anstyle-query
- [ ] anstyle-roff
- [ ] anstyle-svg
- [ ] anstyle-syntect
- [ ] anstyle-termcolor
- [ ] anstyle-wincon
- [ ] anstyle-yansi
- [x] anstyle
- [ ] colorchoice-clap
- [ ] colorchoice
