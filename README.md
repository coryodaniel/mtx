I created this because it's such a pain in the ass to back up in the terminal and change the mix task when I want to watch vs run something one-off. This wraps `mix test` and `mix test.watch` and routes all your flags/args to the task and switches the task if `-w` is set. That's it, that's the fucking repo. Oh, and it updates the terminal tab with your test results so if you have a test suite running all day on a monitor it's easy to see the TDD builds as you are coding.

# mtx

**mtx** is a Go-based wrapper for Elixir's `mix test` command, providing convenient watch and trace options, and (optionally) updating your terminal tab/window title with test results.

## Features

- **Watch mode:** Automatically reruns tests on file changes (`mix test.watch`).
- **Trace mode:** Runs tests with detailed output (`--trace`).
- **Pass-through flags:** Any additional flags or arguments are passed directly to `mix test`.
- **Tab/Window title update:** (iTerm2 and some terminals) Updates the tab/window title with the latest test summary in watch mode.

## Usage

```sh
mtx [flags] [mix test args...]
```

### Flags

- `-w`, `--watch` &nbsp;&nbsp;&nbsp;&nbsp;Run tests in watch mode (`mix test.watch`)
- `-t`, `--trace` &nbsp;&nbsp;&nbsp;&nbsp;Run tests with trace (`mix test --trace`)
- `-h`, `--help` &nbsp;&nbsp;&nbsp;&nbsp;Show help

### Examples

```sh
# Run all tests
mtx

# Run tests in watch mode
mtx --watch

# Run tests with trace
mtx --trace

# Pass additional arguments to mix test
mtx ./test/ --only wip

# Combine options
mtx --watch --trace ./test/ --only wip
```

## Installation

1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd mtx
   ```

2. **Build the binary:**
   ```sh
   go build -o mtx
   ```

3. **(Optional) Move to a directory in your PATH:**
   ```sh
   mv mtx /usr/local/bin/
   ```

## Requirements

- Go 1.18+
- Elixir and `mix` installed

## Notes

- Tab/window title updates are supported in iTerm2 and some other terminals. VS Code's integrated terminal may not support this feature.
- All unrecognized flags and arguments are passed directly to `mix test` or `mix test.watch`.

## License

MIT

---

Let me know if you want to customize any part of this README!
