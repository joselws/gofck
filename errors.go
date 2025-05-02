package main

import "errors"

var ErrInvalidCLIArgsLength = errors.New("invalid number of command line arguments")
var ErrBadFileFormat = errors.New("file does not have .bf extension")
var ErrFileDoesNotExist = errors.New("file does not exist")
var ErrInvalidPointer = errors.New("invalid pointer position")
var ErrInvalidBFSyntax = errors.New("non-matching numbers of brackets")
