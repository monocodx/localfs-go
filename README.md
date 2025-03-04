# LocalFS

LocalFS is a portable, web-based local file server designed for file sharing between trusted devices on the same local network, not for public network.

## Quick Start

### Usage

To start the server, use the `localfs` command, then go to the web interface at `http://localhost:5000`.
```
Usage: localfs [options]
options
  -p, --port           server port to use (default 5000).
      --no-tmpfs       use '/var/tmp' for the temporary directory instead of
                       'tmpfs' to handle large file uploads (linux systems only).
  -h, --help           print this list and exit.
  -v, --version        print the version and exit.
```

### Build

Building from source code requires Go version 1.23 or above. Run the build script to generate the executable binary.

- `build.sh` for linux, macOS, android (termux)
- `build.bat` for windows 10/11 

### Development
Refer to [development guidelines](./DEVELOPMENT.md).

## Changelog
See [What's New](./CHANGELOG.md).

## Known Issues
- linux tmpfs memory limitation, use flag --no-tmpfs for large file uploads.


## LICENSE
LocalFS is licensed under the [GNU GPLv3](./LICENSE).
