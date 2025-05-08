# 🐈‍⬛ Migo: Your Friendly Migration Tool

Migo is a lightweight, fast, and easy-to-use migration tool designed to make your database schema changes a breeze. Think of Migo as your trusty cat companion—always quick, efficient, and here to help you focus on the important stuff: your code.

## Table of Contents

* [About Migo](#about-migo)
* [Features](#features)
* [Getting Started](#getting-started)
* [Usage](#usage)
* [Contributing](#contributing)
* [License](#license)

## About Migo

Migo was created to make database migrations simple and efficient. Whether you're adding new tables or making schema adjustments, Migo helps you manage changes with ease, so you don’t get bogged down with complicated tools.

### Key Features:

* Keeps all migrations neatly organized in a `migo/` folder (don’t worry, no mess here!).
* Tracks applied and pending migrations in a tidy SQLite database.
* You’ll never lose track of a migration again—Migo handles everything with timestamped filenames.
* Rollback your last migration easily, just like a cat retreating from a surprise!

### What Migo currently supports:

* Easy project setup with a clean folder structure.
* Create migration files with timestamped names.
* Automatically apply pending migrations.
* Rollback the most recent migration in a flash!

## Features

* **Organized Migrations:** Create migration files with timestamped names to keep everything tidy.
* **Apply Pending Migrations:** Automatically apply all pending migrations with a single command.
* **Rollback with Ease:** Roll back the last applied migration (don’t worry, it’s like taking a step back after a little mistake!).
* **SQLite Integration:** Track applied and pending migrations using SQLite (for a cozy and consistent experience).
* **What's Next:** Support for more databases and multiple rollbacks coming soon!

## Getting Started

It’s quick and easy to get started with Migo—just follow the steps below to install and set up your project!

### Clone the Repository

Start by cloning the repo and navigating into the project directory:

```bash
git clone https://github.com/wesleybertipaglia/migo.git
cd migo
```

### Choose Your Installation Method

Migo offers two installation options:

#### Option 1: Install Prebuilt Binary (The Easiest!)

No need to worry about Go or complex setups—just install the prebuilt binary and get to work!

**Prerequisites:**

* No Go required! Migo’s binary works out of the box.
* `make` should be available on your system (it’s preinstalled on most UNIX-like systems).

To install the prebuilt binary, run:

```bash
make install
```

> This installs Migo for your system—no fuss, no compilation!

#### Option 2: Build from Source (For the Code Lovers)

Want to compile Migo yourself? Follow these steps:

**Prerequisites:**

* Go 1.23 or higher installed.
* `make` available.

To build and install Migo from source:

```bash
make build
make install
```

> This will build Migo locally and install it to `/usr/local/bin` or any directory in your `PATH`.

## Usage

### Initialize a Migration Project

Creating your migration project is super easy. Just run:

```bash
migo init
```

This creates a neat folder structure for your migrations:

```
migo/
  ├── migrations/
  ├── logs/
  ├── state/
      └── migo.db
```

You can even specify a custom project directory:

```bash
migo init --project ./backend
```

### Create a New Migration

Adding a new migration? Migo makes it a snap:

```bash
migo add --name=create_users_table
```

This will create a migration file like:

```
migo/migrations/20250507123456_create_users_table.sql
```

With placeholders for your SQL:

```sql
-- Migration: create_users_table
-- Created at: 2025-05-07T12:34:56

-- UP

-- DOWN
```

### Apply Pending Migrations

Got some pending migrations? Apply them in one swift move:

```bash
migo update
```

### Roll Back the Last Migration

Made a mistake? No worries, just roll back the last migration:

```bash
migo rollback
```

### Makefile Commands

Handy makefile commands are included for building, installing, cleaning, and more:

```bash
make build      # Build the binary
make install    # Install the binary globally
make clean      # Clean up (goodbye, binary!)
make dev        # Run the project with `go run main.go`
make fmt        # Format the code
make lint       # Lint the code
```

## Contributing

Migo is open-source and we’d love to have your help! If you’ve got ideas, suggestions, or want to fix a bug, feel free to open an issue or submit a pull request. Together, we can make Migo even better!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
