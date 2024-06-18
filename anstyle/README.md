# anstyle for Go

🦀 Rust anstyle crate ported to Go

<table align=center><td>

```go
style := anstyle.NewStyle().
    Bold().
    Underline().
    Foreground(anstyle.ANSIColorRed)
fmt.Printf("%vHello world!%v\n", style, style.RenderReset())
```

<tr><td>

![image](https://github.com/jcbhmr/go-anstyle/assets/61068799/4c76cf13-8389-4ae8-9b80-93de51c9200a)<br>
<sup>[See the example](examples/dump-style.go)</sup>

</table>

## Installation

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)

```sh
go get github.com/jcbhmr/go-anstyle/anstyle
```

## Usage

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)
![Terminal](https://img.shields.io/static/v1?style=for-the-badge&message=Terminal&color=4D4D4D&logo=Windows+Terminal&logoColor=FFFFFF&label=)

```go
style := anstyle.NewStyle().
    Bold().
    Underline().
    Foreground(anstyle.ANSIColorRed)
fmt.Printf("%vHello world!%v\n", style, style.RenderReset())
```

<!-- <sub>📝 This project follows the [Go language binding for Rust](https://jcbhmr.me/Rust-Go/) recommendations.</sub> -->

## Development

![Go](https://img.shields.io/static/v1?style=for-the-badge&message=Go&color=00ADD8&logo=Go&logoColor=FFFFFF&label=)

To get started after cloning this repository you can run the example:

```sh
go run ./examples/dump-style.go
```

This project tries to bring some of the Rust-isms from the anstyle crate to Go. That means things like `Option<T>` are represented as `*T` or `(T, bool)` and `Result<T, E>` is `(T, error)`. I don't really know if that's the best way to represent these things in Go, but it's what I'm doing for now. If you have better ideas or suggestions, please tell me! ❤️
