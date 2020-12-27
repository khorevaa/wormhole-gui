# wormhole-gui

Wormhole-gui is a cross-platform graphical interface for magic-wormhole, providing easy and secure end-to-end encrypted sharing of files, folders and text between computers on a local network.
It uses the Go implementation of magic-wormhole, called [wormhole-william](https://github.com/psanford/wormhole-william), and compiles into a single binary.

<p align="center">
  <img src="internal/assets/screenshot.png" />
</p>

Built using the following Go modules:
- [fyne](https://github.com/fyne-io/fyne) (version 1.4.2)
- [wormhole-william](https://github.com/psanford/wormhole-william) (version 1.0.4 + [f69f6e8](https://github.com/psanford/wormhole-william/commit/f69f6e823d8cec6b3756b8ce63024c8cd3c3ebf2))
- [archiver](https://github.com/mholt/archiver) (version 3.5.0)

The initial version was built in less than one day to show how quick and easy it is to use [fyne](https://github.com/fyne-io/fyne) for developing applications.

## Requirements

Wormhole-gui compiles into a statically linked binary with no runtime dependencies.
Compiling requires a [Go](https://golang.org) compiler (1.13 or later) and the [prerequisites for Fyne](https://developer.fyne.io/started/).

## Downloads

Please visit the [release page](https://github.com/Jacalz/wormhole-gui/releases) for downloading the latest releases.
Versions for Linux (`x86-64` and `arm64`), FreeBSD, MacOS and Windows (`x86-64`) are available.

## Building

Systems with the compile-time requirements satisfied can build the project using `go build` in the project root:
```bash
go build 
```

The project can also be built and installed using GNU Make (installing is currently only supported on Linux and BSD):
```bash
make
sudo make install
```

## Contributing

Contributions are strongly appreciated. Everything from creating bug reports to contributing code will help the project a lot, so please feel free to help in any way, shape or form that you feel comfortable doing.

## License
- Wormhole-gui is licensed under `GNU GENERAL PUBLIC LICENSE Version 3`.
