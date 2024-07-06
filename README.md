# inLama

Using LLM via CLI has never been easier. Intoducing inLama, a CLI tool to use LLM with ease. Pipe your text to inLama and get the output in seconds.

## Description

Leverage the power of LLMs in you terminal. `inLama` allows you you setup a pipeline with any or your existing data generation commands and effortlessly generate intelligent responses.

Do you wish to generate a summary of your logs? `inLama` has got you covered.

```bash
# Generate a summary of your logs
$ cat logs.txt | inLama
```

Do you wish to setup a continuous monitoring system for your servers? `inLama` has got you covered.

```bash
# Monitor your servers
$ tail -f /var/log/syslog | inLama -f
```

## Requirements

- `inLama` currently relies of `ollama` style APIs and a active `ollama` server. The CLI tool provides you a flexibility to configure the server URL (defaults to `http://localhost:11434`).
- `go` version 1.16 or higher.

## Installation

To install `inLama` from source, you can run the following commands:

```bash

# Clone the repository

git clone https://github.com/NishantJoshi00/inlama.git

cd inlama

# Build the binary

make build

# Install the binary

make install

# # If the installation fails due to permission issues, you can run the following command
# sudo make install

```

## Usage

To use `inLama`, you can run the following command:

```bash
$ inlama -h
```

If you wish to setup shell completion for `inLama`, you can run the following command:

- For `bash`:
  ```bash
  source <(inlama --completion bash)
  ```
- For `zsh`:
  ```zsh
  source <(inlama --completion zsh)
  ```
- For `fish`:
  ```fish
  inlama --completion fish | source
  ```
