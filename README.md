# mtx

**mtx** is a wrapper around `mix test` and `mix test.watch`. That’s it.

It passes all your flags and args straight through, swaps in `mix test.watch` when you use `-w`, and updates your terminal tab with test results so you can see your TDD cycle at a glance.

I built this because switching between one-off test runs and watch mode shouldn’t suck. You tweak some args, run a test batch once, then want to go back to watching—except now you’re arrowing back to the start of the line just to replace `test` with `test.watch`. I got arthritis. Ain’t nobody got time for that.

## Features

* 🔁 **Watch mode**: Automatically re-runs tests with `mix test.watch` on file changes.
* 🔍 **Trace mode**: Adds `--trace` for detailed output.
* 🎯 **Pass-through flags**: Everything else goes straight to `mix test` or `mix test.watch`.
* 🪧 **Tab/window title updates**: Shows the latest test result in your terminal tab (works in iTerm2 and some others).

## Usage

```sh
mtx [flags] [mix test args...]
```

### Flags

* `-w`, `--watch` – Use `mix test.watch`
* `-t`, `--trace` – Adds `--trace`
* `-h`, `--help` – Show help

### Examples

```sh
# Regular test run
mtx

# Watch mode
mtx -w

# Trace output
mtx -t

# Run a specific directory with custom filters
mtx ./test --only wip

# Combine flags
mtx -w -t ./test --only wip
```

## Install

```sh
git clone <your-repo-url>
cd mtx
go build -o mtx
mv mtx /usr/local/bin/ # optional
```

## Requirements

* Go 1.18+
* Elixir with `mix`

## Notes

* Tab title updates are tested on iTerm2. VS Code and others may not play nice.
* Any extra flags or args are passed directly to `mix test`.

## License

[MIT](./LICENSE)
