# CoreWars Playground in Go

It's playground for testing solution of [CoreWars](https://www.codewars.com/)' *kata* in Go.

## How to play

1. Copy *kata* tests into [kata/test.go](./kata/tests.go)
2. Copy *kata* solution into [kata/solution.go](./kata/solutions.go)
3. Modify solution as you wish.
4. Run app:
   - in terminal run: `go run .`
   - in [VSC](https://code.visualstudio.com/) press: `F5` or `Ctrl+F5` (results are shown in "debug console")

## Limits

- Values of "solution" and "expected" are compared as texts.
- Keep in mind that your installed Go version can be different (usually newer) then in CodeWars' environment.
- Not all imports are allowed in CodeWars' environment.