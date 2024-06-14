/*
ANSI Text Styling

A portmanteau of “ansi style”

anstyle provides core types describing ANSI styling escape codes for interoperability between crates.

Example use cases:

- An argument parser allowing callers to define the colors used in the help-output without putting the text formatting crate in the public API
- A style description parser that can work with any text formatting crate

Priorities:

API stability
- Low compile-time and binary-size overhead
- const friendly API for callers to statically define their stylesheet

For integration with text styling crate, see:

- anstyle-ansi-term
- anstyle-crossterm
- anstyle-owo-colors
- anstyle-termcolor
- anstyle-yansi

User-styling parsers:

- anstyle-git: Parse Git style descriptions
- anstyle-ls: Parse LS_COLORS style descriptions

Convert to other formats.

- anstream: A simple cross platform library for writing colored text to a terminal
- anstyle-roff: For converting to ROFF
- anstyle-syntect: For working with syntax highlighting

Utilities.

- anstyle-lossy: Convert between anstyle::Color types
- anstyle-parse: Parsing ANSI Style Escapes
- anstyle-wincon: Styling legacy Microsoft terminals

# Examples

The core type is Style:

	let style = anstyle::Style::new().bold();
*/
package anstyle
