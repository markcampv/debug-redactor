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

    - For Linux: `debug-redactor-linux-amd64`
    - For macOS: `debug-redactor-darwin-amd64`
    - For Windows: `debug-redactor-windows-amd64.exe`

3. **Make the Binary Executable** (Linux and macOS):

   After downloading, you may need to make the binary executable. On Linux and macOS, you can do this with the following command:

   ```sh
   chmod +x debug-redactor-linux-amd64  # Adjust the filename for macOS or Linux as necessary



### Usage

Once Debug Redactor is installed, you can run it directly from the command line to redact sensitive information from your Consul debug logs. Here’s how to use Debug Redactor:

**On Linux and macOS:**

```sh
./debug-redactor-linux-amd64 --source /path/to/your/consul-debug.tar.gz
```

```sh
./debug-redactor-darwin-amd64 --source /path/to/your/consul-debug.tar.gz
```
**On Windows:**

```sh
.\debug-redactor-windows-amd64.exe --source C:\path\to\your\consul-debug.tar.gz
```

### Custom Redactions(Currently Under Development)
Future versions of Debug Redactor will allow for custom redaction rules, enabling users to specify additional patterns for sensitive information that should be redacted from debug bundles.



