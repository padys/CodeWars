# The Codewars Playground in Go

It's a playground for comfortable testing the solution of the [Codewars](https://www.codewars.com/)' *kata* in Go.
The playground works in offline (don't need internet connection).

## How to play

1. Copy the *kata* tests into [kata/test.go](./kata/tests.go) (only `var _ = Describe[...]`).
2. Copy the *kata* solution into [kata/solution.go](./kata/solution.go) (whole content).
3. Modify a solution as you wish.
4. Run app:
   - in terminal run: `go run .`
   - in [VSC](https://code.visualstudio.com/) press: `F5` or `Ctrl+F5` (results are shown in "debug console")

## Limitations 

- Values of "solution" and "expected" are compared as texts.
- Keep in mind that your installed Go version can be different (usually newer) then in the Codewars' environment.
- Not all imports are allowed in the Codewars' environment.
