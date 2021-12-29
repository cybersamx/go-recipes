# Goenv

[Goenv](https://github.com/syndbg/goenv) is a Go version manager for switching between different versions of Go in your machine seamlessly.

## Setup

You can install goenv directly using the [official installation instructions](https://github.com/syndbg/goenv/blob/master/INSTALL.md).

If you are on the Mac, you may install goenv via Homebrew by running the following command:

```bash
$ brew install goenv
```

To update goenv, run the following:

```javascript
$ brew update           # Update Homebrew itself
$ brew upgrade goenv    # Upgrade goenv
```

## Usage

You can have multiple versions of Go installed on your machine. And you can specify one version of Go per user (global version) or a version of Go per project (local version).

To get a specific version of Go downloaded to your machine, run the following command:

```bash
$ goenv install --list  # Get a list of downloadable versions of Go
$ goenv install 1.13.4  # Download and install Go version 1.13.4
```

## Unsupported Go Version

There are times when the Go version you want to install isn't found in the list from `goenv install --list`. This means that you can't run `goenv install` to install that particular version. Usually this happens when a new version of Go has not been added by the maintainers of goenv.

In this case, you have to manually install the new version Go and then configure goenv to use the new version.

You can use [this script](./goenv.sh) to install, shim, and bind an unsupported Go framework to goenv - currently the script only support MacOS. Or follow the instructions below.

1. Go to the official [Go download page](https://golang.org/dl/) to download the latest. Say it's go version 1.14.3. **Please download the tarball for the Go binaries/libraries. Do not download the pkg (MacOS) or msi (Windows) file**. You want to unarchive the file into a directory and move it to the right path for goenv to recognize it.

1. Expand the file and copy it to the goenv directory.

   ```bash
   $ # Should expand to `go` directory
   $ tar xvfz go1.14.3.darwin-amd64.tar.gz
   $ # Be sure to rename the directory to its canonical version
   $ mv go ~/.goenv/versions/1.14.3
   ```

1. Verify that goenv recognizes the new Go version.

   ```bash
   $ goenv versions
     system
     1.10.3
     1.11.4
   * 1.13.4 (set by /Users/schow/.goenv/version)
     1.14.3
   ```    

1. Switch the global version of Go to `1.14.3`.

   ```bash
   $ goenv global 1.14.3
   $ goenv vesions
     system
     1.10.3
     1.11.4
     1.13.4
   * 1.14.3 (set by /Users/schow/.goenv/version)
   $ go version
   go version go1.14.3 darwin/amd64
   ```

## Reference

* [Ooenv](https://github.com/syndbg/goenv)
