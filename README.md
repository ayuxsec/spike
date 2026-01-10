## Some context

This project is something I've been building in my spare time. If you simply want to try it out, head to the Releases section, download the latest version, and run it directly.

If you're on a different operating system or CPU Arch, install it using Go tool chain:

```console
go install github.com/ayuxsec/spike/cmd/spike@latest
```

You can email me at ayuxsec@proton.me if you need help with something or want to contribute.

## Flow Chart

see [pipeline.go](internal/pkg/scanner/core/pipeline.go)

```mermaid
flowchart TD
    A[__root_domain__] -->|wildcard? yes| B[subfinder]
    B --> C[httpx]

    A -->|wildcard? no| C

    A --> D[gau]

    C --> E[cachex]
    C --> F[katana]
    C --> G[nuclei-generic]

    F --> H[uro]
    D --> H
    H --> I[httpx]

    I --> J[nuclei-dast]
```

## config

config example

```yaml
tools:                                 # Tool configuration
    httpx:
        threads: 25                    # Number of concurrent httpx workers
        ports_to_scan:                 # Ports to probe for web services
            http: 80,8080,8000,8008,8888,3000,5000,9000,81,82,83,84,591,2082,2086,2095,10000
            https: 443,8443,9443,5001,3001,8001,8081,2083,2087,2096,10001,10443,10444
        screenshot: false              # Capture screenshots of discovered pages
        cmd_timeout_in_second: 900     # Max execution time

    subfinder:
        threads: 10                   # Concurrent subfinder workers
        enabled: true                 # Enable subdomain discovery
        cmd_timeout_in_second: 900    # Max execution time

    katana:
        enabled: true                 # Enable crawling
        threads: 10                   # Concurrent crawling threads
        crawl_depth: 3                # Maximum crawl depth
        max_crawl_time: 10m           # Total crawl time limit
        parallelism_threads: 10       # Parallel browser workers
        headless: false               # Run browser in headless mode
        no_sandbox: false             # Disable browser sandbox
        cmd_timeout_in_second: 900    # Max execution time

    gau:
        enabled: true                 # Enable URL collection
        threads: 10                   # Concurrent workers
        cmd_timeout_in_second: 900    # Max execution time

    nuclei:
        threads: 25                   # Concurrent scan threads
        template_settings:
            generic: true             # Enable generic template set
            dast: true                # Enable DAST templates
            headless: false           # Run headless templates
        template_paths:
            generic:
                include:              # Template directories to include
                    - http/
                    - cloud/
                    - javascript/
                    - dns/
                    - ssl/
                    - network/
                    - http/cves/2024
                    - http/cves/2023
                    - http/cves/2022
                    - http/cves/2021
                    - http/cves/2020
                exclude:              # Template directories to exclude
                    - http/cves/
                exclude_severity:     # Templates to ignore by severity
                    - info
            dast:
                include:              # DAST template directories
                    - dast/
                exclude:              # Template directories to exclude
                    - ""
                exclude_severity:     #  Templates to ignore by severity
                    - info
        cmd_timeout_in_second: 900    # Max execution time

    cachex:
        enabled: true                 # Enable cache poisoning testing
        threads: 10                   # Concurrent workers
        cmd_timeout_in_second: 900    # Max execution time

reporter:                            # Reporting configuration
    telegram:
        enabled: false               # Enable Telegram notifications
        bot_token: ""                # Telegram bot token
        chat_id: 0                   # Target chat ID
        request_timeout: 10          # Telegram API timeout
```
