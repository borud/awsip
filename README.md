# awsip

`awsip` is a simple command line utility to list IP ranges for AWS.  

- You can limit the IP ranges by region, service and IP version.  
- Output to table, CSV, markdown and HTML.  
- Caches the data for 3 days before re-fetching it
- You can override `--max-cache-age` to set shorter refresh interval
- and you can specify `--no-cache` to bypass cache and fetch data directly

## Installing

If you have Go installed you can install awsip as a Go application:

```shell
go install github.com/borud/awsip/cmd/awsip@latest
```

## Running

`awsip` implements three subcommands.

- `regions` -- lists the regions available
- `services` -- lists the services available
- `range` -- lists IP ranges given regions and/or service constraints

The range command has the following command line options:

```text
Usage:
  awsip [OPTIONS] range [range-OPTIONS]

Application Options:
  -n, --no-cache                             do not use cache, download list directly
  -a, --max-cache-age=                       max age of cached file before re-download (default: 72h)
  -u, --url=                                 URL of AWS ranges JSON file (default: https://ip-ranges.amazonaws.com/ip-ranges.json)
  -v, --verbose                              verbose output

Help Options:
  -h, --help                                 Show this help message

[range command options]
      -s, --service=                         service, use AMAZON for superset of all services (default: AMAZON)
      -r, --region=                          region, use GLOBAL for superset of all regions (default: GLOBAL)
      -i, --ip=[4|6|0]                       IP version, use 0 for both IPv4 and IPv6 (default: 0)
      -f, --format=[table|html|markdown|csv] output format (default: table)
      -c, --no-color                         do not colorize output
```
