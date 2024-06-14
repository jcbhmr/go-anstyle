# Rust anstyle for Go

🦀 Rust anstyle crate ported to Go

<table align=center><td>

</table>

<p align=center>
  <a href="https://docs.rs/anstyle">Docs</a>
  | <a href="https://github.com/rust-cli/anstyle">Original rust-cli/anstyle</a>
</p>

- **[anstyle](anstyle):** 🎨 ANSI color `struct` and `interface`
<!-- - **[anstream](anstream):** 🌊 Helpers for ANSI output streams
- **[anstylegit](anstylegit):** 🔸 Parse Git color config to `Style`
- **[anstylelossy](anstylelossy):** 🗑️ Lossy color conversions
- **[anstylels](anstylels):** 📁 Parse `LS_COLORS` into a `Style` config
- **[anstyleparse](anstyleparse):** 🔠 Parse ANSI style escapes
- **[anstyleroff](anstyleroff):** 📄 Convert ANSI to ROFF
- **[anstylesvg](anstylesvg):** 🖼️ Convert ANSI to SVG
- **[anstylewincon](anstylewincon):** 🟦 Style legacy Windows terminals
- **[colorchoiceclap](colorchoiceclap):** 👏 Automatic `--color` flag for Clap -->

## Installation

You're reading the root monorepo readme. 😉 Click the links above 👆 to check out subproject-specific installation instructions.

<sup>Spoiler: it's `go get github.com/jcbhmr/go-anstyle/<subproject>`</sup>

## Usage

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

Each of the Rust crates from the [rust-cli/anstyle](https://github.com/rust-cli/anstyle) repository are ported to Go in this monorepo as separate Go modules, each with their own `go.mod` file and their own version. Go monorepo versioning works if the Git tag starts with a `<subfolder>/` prefix like `anstyle/v0.1.0`.

- [ ] anstream
- [ ] anstyle-ansi-term
- [ ] anstyle-crossterm
- [ ] anstyle-git
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
- [ ] anstyle
- [ ] colorchoice-clap
- [ ] colorchoice
