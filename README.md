# flipper
flippin pancakes!

Take home coding challenge.<br/>
The builtin basic version uses os.Stdin and os.Stdout, but you can use whater `io.Reader` and `io.Writer` you choose to pass in.<br/>
The `example` directory has a `package main` you can run, or it's easy to import and make your own.
Included is `generate.py` (<a href="https://www.geeksforgeeks.org/print-all-combinations-of-given-length/">attribution</a>) which will generate permutations of `n` length and output directly in the format `flipper` accepts for the problem. Usage: `python3 generate.py n`<br/>
Also included is a simple bash wrapper `runPermutations.sh` wrapping the generator and piping to `example` if you've cloned the repo and built the example. Usage `./runPermutations.sh n`
