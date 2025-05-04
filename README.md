# gofck

Brainfuck interpreter in Go. 

Run Brainfuck files with this CLI interpreter.

## Brainfuck

[This fireship video explains what it is much better.](https://www.youtube.com/watch?v=hdHjjBS4cs8 "Brainf**k in 100 seconds.")

[Here's its Wikipedia for more information.](https://en.wikipedia.org/wiki/Brainfuck "Wikipedia - Brainfuck")

## Install

Pre-requisites: must have `go` installed in your machine.

1. Clone the repository.

2. `cd` to the repository.

3. Run `go build -o gofck`

4. Run any `.bf` or `.b` (brainfuck) files with `./gofck <your_file.bf>`

A couple examples have been included in the repository for demonstration:

```bash
./gofck example.bf
```

## gofck

The `,` brainfuck instruction is currently not supported.
