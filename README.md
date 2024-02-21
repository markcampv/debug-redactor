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

- Go version 1.15 or higher.
- Access to Consul debug logs that you wish to redact.

### Installation

Clone the repository and build the tool:

```sh
git clone https://github.com/yourusername/debug-redactor.git
cd debug-redactor
go build -o debug-redactor

