# gvt-unused

[![Build Status](https://travis-ci.org/slomek/gvt-unused.svg?branch=master)](https://travis-ci.org/slomek/gvt-unused)

A small [gvt](https://github.com/FiloSottile/gvt)-based tool to list unused dependencies that can be removed from your vendor.

## Disclaimer

This application does not perform any action on your vendor files - it acts as a kind friend who has some advice which you may or may not follow. In the end it's your life, and nobody should be blamed by your choices other than yourself.

# Installation & Usage

In order to install `gvt-unused` type:

    $ go get -u github.com/slomek/gvt-unused/...

In order to list unused dependencies in a project using `gvt` as a vendoring tool, type:

    $ gvt-unused
    ‣ Listing dependencies from manifest file...
    ‣ Listing dependencies from source code...
    ↳ Looking in the sources
    ↳ Looking in the test sources
    ↳ Filtering non-standard imports...
    ...

If you are using a non-standard manifest JSON location (`vendor/manifest`), you can define a custom one:

    $ gvt-unused -manifest vendor/manifest-file-path

## Kudos

Kudos to [Filippo Valsorda](https://github.com/FiloSottile) for making [gvt](https://github.com/FiloSottile/gvt).
