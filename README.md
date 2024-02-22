# Debug Redactor

## Overview

Debug Redactor is a easy-to-use tool designed to enhance the security of your Consul debug logs. It automatically redacts sensitive information, such as IP addresses, ensuring that your debugging processes adhere to privacy standards and regulations. Whether you're sharing logs with support teams, storing them for analysis, or auditing system behavior, Debug Redactor helps keep your sensitive data safe.

## Features

- **Automated Redaction**: Automatically detects and redacts IP addresses and tokens from Consul debug logs.
- **Custom Redaction Rules**: Easily extendable to include custom redaction patterns based on your specific needs.
- **User-Friendly**: Simple command-line interface for straightforward operation.
- **Efficient Processing**: Designed to handle large log files with minimal performance impact.
- **Secure Handling**: Ensures that the original logs are not modified, creating redacted copies for safe sharing and storage.

## Getting Started

### Prerequisites

- Access to Consul debug bundle that you wish to redact.

### Installation

You can download pre-built binaries of Debug Redactor directly from our GitHub [Releases](https://github.com/markcampv/debug-redactor/releases) page. This allows you to install the tool without needing to compile it from the source.

1. **Go to the Releases Page**: Navigate to [Releases](https://github.com/markcampv/debug-redactor/releases) in the Debug Redactor repository.

2. **Download the Binary**: Download the appropriate binary for your operating system and architecture. We provide binaries for Windows, macOS, and Linux.

    - For Linux: `debug-redactor-linux-amd64.zip`
    - For macOS: `debug-redactor-darwin-amd64.zip`
    - For Windows: `debug-redactor-windows-amd64.zip`

3. **Make the Binary Executable** (Linux and macOS):

   After downloading, you may need to make the binary executable. On Linux and macOS, you can do this with the following command:

```sh
   chmod +x debug-redactor 
```


### Usage

Debug Redactor is designed to enhance the privacy of your Consul debug bundles by redacting sensitive information. Once installed, you can use it by specifying the source file with the `--source` flag. The command below illustrates how to run Debug Redactor:

```sh
./debug-redactor --source /path/to/your/consul-debug.tar.gz
```
### Example Output
Here's an example of running Debug Redactor on a Consul debug bundle from a [hcdiag](https://github.com/hashicorp/hcdiag/) bundle:
```shell
./debug-redactor --source ~/hcdiag-2024-01-29T093000Z/ConsulDebug17709286985/ConsulDebug.tar.gz

File processed successfully. Output file: ~/hcdiag-2024-01-29T093000Z/ConsulDebug17709286985/ConsulDebug.tar-redacted.gz
```

This indicates that the tool has successfully processed the input file and created a new, redacted version named ConsulDebug-redacted.tar.gz in the same directory as the source file. The redacted file omits sensitive information, making it safer for sharing or analysis.

### Redacted IPs in logs:

The following examples demonstrate how IP addresses are redacted within the logs, ensuring enhanced privacy and security:

```shell
[DEBUG] agent.server.memberlist.lan: memberlist: Initiating push/pull sync with: node-1 [REDACTED]:8301
[DEBUG] agent.server.memberlist.lan: memberlist: Initiating push/pull sync with: node-2 [REDACTED]:8301
[DEBUG] agent.server.memberlist.lan: memberlist: Initiating push/pull sync with: node-3 [REDACTED]:8301
```

### Custom Redactions(Currently Under Development)
Future versions of Debug Redactor will allow for custom redaction rules, enabling users to specify additional patterns for sensitive information that should be redacted from debug bundles.



