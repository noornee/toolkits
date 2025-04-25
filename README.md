# Toolkits

**Toolkits** is a collection of command-line utilities designed to streamline development workflows.

## Build

To build the CLI binary:

```bash
go build -o toolkits
```

After building, you can check available commands and flags with:

```bash
./toolkits --help
```

---

## Commands

<details>
<summary><strong>Redis</strong> – Store and retrieve data from Redis</summary>

The `redis` command in **Toolkits** allows you to interact with a Redis database directly from the command line. It includes two subcommands: `set` and `get`.

---

### Motivation

This command was created to make working with Redis less frustrating, especially when dealing with JSON files. Normally, using `redis-cli` requires escaping all the quotes in the JSON file, which is a bit stressful. i'll rather not open my browser to look for some website to help escape the quotes. (typing this out, i just got a new idea lmao. i just might add another tool that helps to escape json quotes-- that is if there's none that does that already)

While tools like **RedisInsight** exist and make things easier, I don’t always have access to a full development environment. Sometimes I'm on my phone and I just need a quick way to save a file’s contents into Redis without hassle. this command solves that problem for me.

---
### Usage

```bash
toolkits redis [subcommand] [flags]
```

If the `--env` flag is omitted, the CLI will attempt to load a `.env` file from the current directory.

---

### Subcommands

#### `set`

Stores the contents of a file into Redis under a specified key.

**Usage:**

```bash
toolkits redis set --key <redis-key> --file <path-to-file> [--env <path-to-env>]
```

**Example:**

```bash
toolkits redis set --key "templates" --file ../template.json --env sample.env
```

This command reads the contents of `../template.json` and stores it in Redis under the key `templates`, using credentials loaded from `sample.env`.

---

#### `get`

Retrieves the value stored in Redis under a specific key and prints it to the console.

**Usage:**

```bash
toolkits redis get --key <redis-key> [--env <path-to-env>]
```

**Example:**

```bash
toolkits redis get --key "templates" --env sample.env
```

This retrieves and prints the value stored in the Redis key `templates`.

---

### Environment File

Check the `sample.env` file to see the expected format for the environment file.

</details>

