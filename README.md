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
