# anstyle-git for Go

🦀 Rust anstyle-git crate ported to Go

<table align=center><td>

```go

```

</table>

## Installation

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)

```sh
go get github.com/jcbhmr/go-anstyle/anstylegit
```

## Usage

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)

```go
style := anstylegit.Parse("red blue")
log.Println(style.GetFgColor())
//=> red
log.Println(style.GetBgColor())
//=> blue
```

> #### `color.*`
>
> If you want to be more specific about which commands are colored and how, Git provides verb-specific coloring settings. Each of these can be set to `true`, `false`, or `always`:
>
>     color.branch
>     color.diff
>     color.interactive
>     color.status
>
> In addition, each of these has subsettings you can use to set specific colors for parts of the output, if you want to override each color. For example, to set the meta information in your diff output to blue foreground, black background, and bold text, you can run:
>
> ```
> $ git config --global color.diff.meta "blue black bold"
> ```
>
> You can set the color to any of the following values: `normal`, `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, or `white`. If you want an attribute like bold in the previous example, you can choose from `bold`, `dim`, `ul` (underline), `blink`, and `reverse` (swap foreground and background).
