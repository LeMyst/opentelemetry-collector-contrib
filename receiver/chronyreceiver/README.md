# Chrony Receiver

<!-- status autogenerated section -->
| Status        |           |
| ------------- |-----------|
| Stability     | [alpha]: metrics   |
| Distributions | [contrib] |
| Issues        | ![Open issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aopen%20label%3Areceiver%2Fchrony%20&label=open&color=orange&logo=opentelemetry) ![Closed issues](https://img.shields.io/github/issues-search/open-telemetry/opentelemetry-collector-contrib?query=is%3Aissue%20is%3Aclosed%20label%3Areceiver%2Fchrony%20&label=closed&color=blue&logo=opentelemetry) |

[alpha]: https://github.com/open-telemetry/opentelemetry-collector#alpha
[contrib]: https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib
<!-- end autogenerated section -->

The [chrony] receiver is a pure go implementation of the command `chronyc tracking` to allow for
portability across systems and platforms. All of the data that would typically be captured by
the tracking command is made available in this receiver, see [documentation](./documentation.md) for
more details.

## Configuration

### Default

By default, the `chrony` receiver will default to the following configuration:

```yaml
chrony/defaults:
  address: unix:///var/run/chrony/chronyd.sock # The default port by chronyd to allow cmd access
  timeout: 10s # Allowing at least 10s for chronyd to respond before giving up

chrony:
  # This will result in the same configuration as above
```

### Customised

The following options can be customised:

- address (required) - the address on where to communicate to `chronyd`
  - The allowed formats are the following
    - udp://hostname:port
    - unix:///path/to/chrony.sock (Please note the triple slash)
    - unixgram:///path/to/chrony/sock
  - The network type `unix` will be converted to `unixgram` but both are permissible
- timeout (optional) - The total amount of time allowed to read and process the data from chronyd
  - Recommendation: This value should be set above 1s to allow `chronyd` time to respond
- collection_interval (optional) - how frequent this receiver should poll [chrony]
- `initial_delay` (default = `1s`): defines how long this receiver waits before starting.
- metrics (optional) - Which metrics should be exported, read the [documentation] for complete details

## Example

An example of the configuration is:

```yaml
receivers:
  chrony:
    address: unix:///var/run/chrony/chronyd.sock
    timeout: 10s
    collection_interval: 30s
    metrics:
      ntp.skew:
        enabled: true
      ntp.stratum:
        enabled: true
```

The complete list of metrics emitted by this receiver is found in the [documentation].

[documentation]: ./documentation.md
[chrony]: https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system_administrators_guide/ch-configuring_ntp_using_the_chrony_suite