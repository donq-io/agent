---
aliases:
- /docs/agent/latest/flow/configuration-language/standard-library/env
title: env
---

# `env` Function

`env` gets the value of an environment variable from the system Grafana Agent
is running on. If the environment variable does not exist, `env` returns an
empty string.

## Examples

```
> env("HOME")
"/home/grafana-agent"

> env("DOES_NOT_EXIST")
""
```
