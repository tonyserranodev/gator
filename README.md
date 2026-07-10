# Gator

Gator is a CLI RSS feed aggregator written in Go. It lets you register users, follow RSS feeds, and browse the latest posts from those feeds.

## Prerequisites

Before running Gator, make sure you have the following installed:

- [Go](https://go.dev/doc/install) (version 1.25.5 or later)
- [Postgres](https://www.postgresql.org/download/) (running locally or accessible via network)

## Installation

Install the `gator` CLI using `go install`:

```bash
go install github.com/tonyserranodev/gator@latest
```

Make sure `$GOPATH/bin` (usually `~/go/bin`) is in your `$PATH`.

## Database Setup

Start Postgres and create a database named `gator`:

   ```bash
   createdb gator
   ```

   ```

## Configuration

Gator looks for a config file at `~/.gatorconfig.json`. Create it with the following contents:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator",
  "current_user_name": ""
}
```

Replace the `db_url` values with your actual Postgres credentials.

## Usage

Run commands with:

```bash
gator <command> [args...]
```

### Example Commands

- **Register a new user:**

  ```bash
  gator register alice
  ```

- **Log in as a user:**

  ```bash
  gator login alice
  ```

- **Add an RSS feed:**

  ```bash
  gator addfeed "Go Blog" https://go.dev/blog/feed.atom
  ```

- **Follow a feed:**

  ```bash
  gator follow https://go.dev/blog/feed.atom
  ```

- **List followed feeds:**

  ```bash
  gator following
  ```

- **Aggregate feeds in the background:**

  ```bash
  gator agg 1m
  ```

- **Browse recent posts:**

  ```bash
  gator browse 10
  ```

- **List all available commands:**

  Run `gator` without arguments to see the command usage message.
