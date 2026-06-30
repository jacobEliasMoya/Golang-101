# Golang 101: git-go

`git-go` is a beginner Golang CLI project that is meant to teach Go while building a useful developer tool.

The final executable will be named `ggo`. The first real feature is a shortcut for creating a fix commit, pushing it, and keeping the command easy to remember.

## Goal

The intended command will look like this:

```bash
ggo -F "Testing push ability"
```

That should eventually run this Git workflow:

```bash
git add .
git commit -m "(fix) Testing push ability."
git push
```

For now, the project should stay intentionally small:

- Accept only the `-F` flag.
- Read one quoted commit message.
- Format that message as a fix commit.
- Run Git commands in order.
- Stop immediately if any Git command fails.
- Teach Go basics clearly along the way.

Future flags like `-A`, `-R`, `-C`, or `-D` may be added later, but they are out of scope for the first version.

## Project Structure

```text
git-go/
  go.mod
  main.go
```

`go.mod` tells Go that this folder is a Go module. It records the module name, the Go version, and any dependencies added later.

This project should use only the Go standard library for the first version.

## Commands

Run during development:

```bash
cd git-go
go run main.go -F "Testing push ability"
```

Format Go code:

```bash
gofmt -w main.go
```

Build the executable:

```bash
go build -o ggo
```

On Windows, Go will create:

```text
ggo.exe
```

Run the built executable from the local folder:

```powershell
.\ggo.exe -F "Testing push ability"
```

Later, adding the executable folder to `PATH` would allow this from any Git repository:

```powershell
ggo -F "Testing push ability"
```

## Expected Input

Valid:

```bash
ggo -F "Testing push ability"
```

The quotes matter because they keep the multi-word message together as one terminal argument.

Invalid examples:

```bash
ggo
ggo -F
ggo -X "Testing push ability"
```

Expected error messages:

```text
Not enough arguments present. Use -F "your message"
Unknown flag. Use -F for a fix commit.
```

## Learning Roadmap

### Phase 1: Basic Go CLI

Build a `main.go` file that prints:

```text
git-go is running
```

Concepts introduced:

- `package main` marks this as an executable Go program.
- `import "fmt"` brings in Go's formatting and printing tools.
- `func main()` is the entry point where the program starts.
- `fmt.Println()` prints text to the terminal.
- `go run main.go` compiles and runs the file immediately for development.

Next phase: read command-line arguments.

### Phase 2: Read Command-Line Arguments

Support:

```bash
go run main.go -F "Testing push ability"
```

Concepts introduced:

- `os.Args` is a slice containing command-line arguments.
- A slice is similar to a JavaScript array, but it has Go's stricter typing rules.
- `args[0]` is the executable path, not the first user-provided value.
- User input begins at `args[1]`.
- `len(args)` checks how many values are present.
- Argument validation must happen before reading `args[1]` or `args[2]`.
- An early `return` stops the program before unsafe code runs.

Expected valid output:

```text
Fix message: Testing push ability
```

Next phase: format the commit message with a helper function.

### Phase 3: Helper Functions and Typed Return Values

Add a helper:

```go
func buildFixMessage(message string) string {
	return fmt.Sprintf("(fix) %s.", message)
}
```

Concepts introduced:

- Function parameters have names and types.
- Return values also have types.
- `return` sends a value back to the caller.
- Helpers live outside `main()` so `main()` stays readable.
- `fmt.Sprintf` builds a string instead of printing it immediately.
- `%s` is a placeholder for a string value.

Expected output:

```text
Generated commit message: (fix) Testing push ability.
```

Next phase: run `git add .`.

### Phase 4: Run `git add .`

Add `os/exec` and create a helper such as:

```go
func runGitAdd() error
```

Concepts introduced:

- `exec.Command("git", "add", ".")` runs a command with separate arguments.
- Separate arguments are safer than building shell strings with user input.
- `cmd.Stdout = os.Stdout` shows normal command output in the terminal.
- `cmd.Stderr = os.Stderr` shows command errors in the terminal.
- `git add .` is a side effect because it changes Git repository state.
- Functions that can fail should return an `error`.

Next phase: commit using the formatted message.

### Phase 5: Run `git commit -m`

Add a helper such as:

```go
func runGitCommit(message string) error
```

Concepts introduced:

- Variables can be passed into functions.
- Git command arguments should remain separate.
- Errors should be checked immediately.
- If commit fails, the program should stop before pushing.

Next phase: push only after a successful commit.

### Phase 6: Run `git push`

Add a helper such as:

```go
func runGitPush() error
```

The final flow should be:

```text
Validate input
Build commit message
Run git add .
If that succeeds, run git commit
If that succeeds, run git push
Print success
```

Concepts introduced:

- Sequential control flow keeps risky operations predictable.
- `git push` must never run if `git add` or `git commit` fails.
- Go commonly handles errors with:

```go
if err := runGitAdd(); err != nil {
	fmt.Println("Failed to stage files:", err)
	return
}
```

Next phase: build the CLI as `ggo`.

### Phase 7: Build the Executable

Build with:

```bash
go build -o ggo
```

Concepts introduced:

- `go build` creates a compiled program.
- On Windows, the executable is `ggo.exe`.
- A compiled executable can run without using `go run`.
- `PATH` is a system setting that lets terminals find commands globally.

## Desired First-Version Output

For:

```bash
ggo -F "Testing push ability"
```

The tool should eventually print something like:

```text
Creating fix commit...
Staging changes...
Creating commit: (fix) Testing push ability.
Pushing changes...
Done.
```

Git's own output should remain visible when useful.

## Future Improvements

- Add more commit types.
- Add better help output with `-h` or `--help`.
- Add tests for message formatting.
- Support messages without requiring quotes.
- Install `ggo` somewhere on `PATH`.
- Split the code into multiple files only when the program has enough logic to justify it.
