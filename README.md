# Gator — RSS Feed Aggregator CLI

Gator is a lightweight command-line RSS feed aggregator written in Go. It lets you collect posts from across the internet, store them locally, and browse summaries right from your terminal.

---

## Features

* Add RSS feeds from anywhere on the web
* Store posts in a PostgreSQL database
* Follow and unfollow feeds added by other users
* View summarized posts with links to full articles
* Periodically aggregate and fetch new posts
* Simple multi-user support with login and registration

---

## Prerequisites

Before using Gator, make sure you have the following installed:

* **Go** (1.20 or newer recommended)
* **PostgreSQL** (running locally or accessible remotely)

You can verify installations with:

```bash
go version
psql --version
```

---

## Installation

Install the Gator CLI using `go install`:

```bash
go install github.com/yfaheid/gator@latest
```

Make sure your Go binary directory is in your `PATH`. Typically:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

## Configuration

Gator uses a config file to store database connection details and user preferences.

### 1. Create a config file

Create a file at:

```bash
~/.gatorconfig.json
```

### 2. Example config

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user": "your_username"
}
```

> `current_user` is managed automatically by `gator register` and `gator login`. You typically should not edit it manually.

### 3. Set up your database

Make sure your PostgreSQL database exists:

```sql
CREATE DATABASE gator;
```

Run any migrations required by your project (if applicable).

---

## Running Gator

Once installed and configured, you can run Gator commands directly:

```bash
gator <command>
```

---

## Common Commands

Here are a few useful commands to get started:

### Register a user

```bash
gator register alice
```

Creates a new user account and sets it as the current user.

---

### Login

```bash
gator login alice
```

Switches the active user.

---

### List users

```bash
gator users
```

Displays all registered users.

---

### Add a feed

```bash
gator addfeed "TechCrunch" https://techcrunch.com/feed/
```

Adds a new RSS feed to the system.

---

### Follow a feed

```bash
gator follow https://techcrunch.com/feed/
```

Start following a feed.

---

### Unfollow a feed

```bash
gator unfollow https://techcrunch.com/feed/
```

Stop following a feed.

---

### List feeds

```bash
gator feeds
```

View all available feeds.

---

### Browse posts

```bash
gator browse
```

Displays a list of recent posts with summaries and links.

---

### Aggregate feeds

```bash
gator agg 30s
```

Fetches and stores the latest posts from followed feeds every `30s`.

You can use any valid Go duration format, such as:

```bash
gator agg 10s
gator agg 1m
gator agg 5m
gator agg 1h
```

---

## Quick Command Reference

```bash
gator register <username>
gator login <username>
gator users

gator addfeed "<name>" <feed_url>
gator follow <feed_url>
gator unfollow <feed_url>

gator feeds
gator browse
gator agg <time_interval>
```
