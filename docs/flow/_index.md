---
aliases:
- /docs/agent/latest/flow/
title: Grafana Agent Flow
weight: 900
---

# Grafana Agent Flow (Experimental)

Grafana Agent Flow is a _component-based_ experimental revision of Grafana
Agent with a focus on ease-of-use, debuggability, and ability to adapt to the
needs of power users.

> **EXPERIMENTAL**: Grafana Agent Flow is an [experimental][] feature.
> Experimental features are subject to frequent breaking changes and are
> subject for removal if the experiment doesn't work out.
>
> You should only use Grafana Agent Flow if you are okay with bleeding edge
> functionality and want to provide feedback to the developers. It is not
> recommended to use Grafana Agent Flow in production.

[experimental]: {{< relref "../operation-guide#stability" >}}

## Features

* Write declarative configurations with a Terraform-inspired configuration
  language
* Declare components to configure parts of a pipeline
* Use expressions to bind components together to build a programmable pipeline

## Example

```river
// Discover Kubernetes pods to collect metrics from.
discovery.kubernetes "pods" {
  role = "pod"
}

// Scrape metrics from Kubernetes pods and send to a prometheus.remote_write
// component.
prometheus.scrape "default" {
  targets    = discovery.kubernetes.pods.targets
  forward_to = [prometheus.remote_write.default.receiver]
}

// Get an API key from disk.
local.file "apikey" {
  filename  = "/var/data/my-api-key.txt"
  is_secret = true
}

// Collect and send metrics to a Prometheus remote_write endpoint.
prometheus.remote_write "default" {
  remote_write {
    url = "http://localhost:9009/api/prom/push"

    basic_auth {
      username = "MY_USERNAME"
      password = local.file.apikey.content
    }
  }
}
```

## Next steps

* Learn about the [core concepts][] of Grafana Agent Flow.
* Follow our [tutorials][] to get started with Grafana Agent Flow.
* Learn how to use Grafana Agent Flow's [configuration language][].
* Check out our [reference documentation][] to find specific information you
  might be looking for.

[core concepts]: {{< relref "./concepts/" >}}
[tutorials]: {{< relref "./tutorials/ ">}}
[configuration language]: {{< relref "./config-language/" >}}
[reference documentation]: {{< relref "./reference" >}}

## Provide feedback

Feedback about Grafana Agent Flow and its configuration language can be
provided in our dedicated [GitHub discussion for feedback][feedback].

[feedback]: https://github.com/grafana/agent/discussions/1969
