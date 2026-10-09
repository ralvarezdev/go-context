# go-context

Context helper functions for Go projects.

**Note:** This repository is archived and read-only.

## Features

- **`CancelContextOnError`** — wraps a `func(context.Context) error` so that the context is cancelled when it returns an error or panics (the panic is recovered and returned as an error).

## License

GPL-3.0.
