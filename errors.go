package main

import "errors"

var ErrInvalidCLIArgsLength = errors.New("invalid number of command line arguments")
var ErrBadFileFormat = errors.New("file does not have .bf extension")
var ErrFileDoesNotExist = errors.New("file does not exist")
