# Enoughparse

Enoughparse is a command-line tool for extracting GPS coordinates from image files and generating Google Maps links.

## Building

To build the project, run:

```bash
go build -ldflags "-s -w -X github.com/gkwa/enoughparse/version.Version=$(git describe --tags --always --dirty) -X 'github.com/gkwa/enoughparse/version.Date=$(date -u +"%Y-%m-%dT%H:%M:%SZ")' -X 'github.com/gkwa/enoughparse/version.GoVersion=$(go version)' -X github.com/gkwa/enoughparse/version.ShortGitSHA=$(git rev-parse --short HEAD) -X github.com/gkwa/enoughparse/version.FullGitSHA=$(git rev-parse HEAD)" -o enoughparse
```

## Usage

### Basic usage:

```bash
./enoughparse hello /path/to/your/image.jpg
```

This will output a Google Maps link for the GPS coordinates in the image.

### Increasing verbosity:

You can increase the verbosity level by adding one or more `-v` flags:

```bash
./enoughparse hello /path/to/your/image.jpg -v
./enoughparse hello /path/to/your/image.jpg -vv
./enoughparse hello /path/to/your/image.jpg -vvv
```

Each additional `v` increases the verbosity level, showing more detailed logs.

### JSON output:

To get JSON-formatted logs, use the `--log-format json` flag:

```bash
./enoughparse hello /path/to/your/image.jpg --log-format json
```

### Viewing version information:

```bash
./enoughparse version
```

## Example Output

```
$ ./enoughparse hello /path/to/image.jpg -vvv
12:26PM TRC core/gps.go:10 > Parsing GPS coordinates logger=root v=2
12:26PM DBG core/exif.go:11 > Initializing exiftool logger=root v=1
12:26PM DBG core/exif.go:19 > Extracting metadata imagePath=/path/to/image.jpg logger=root v=1
12:26PM TRC core/gps.go:17 > Parsing latitude latStr="47 deg 36' 42.94\" N" logger=root v=2
12:26PM TRC core/gps.go:24 > Parsing longitude logger=root lonStr="122 deg 18' 46.86\" W" v=2
12:26PM TRC core/gps.go:31 > Checking longitude reference logger=root lonRef=West v=2
12:26PM DBG core/gps.go:41 > Extracted GPS coordinates gpsInfo="Latitude: 47.611928, Longitude: -122.313017" logger=root v=1
12:26PM DBG cmd/hello.go:23 > Extracted GPS coordinates gpsInfo="Latitude: 47.611928, Longitude: -122.313017" logger=root v=1
https://www.google.com/maps?q=47.611928,-122.313017
```

This example shows the tool extracting GPS coordinates from an image and generating a Google Maps link, with detailed logging enabled.
```
