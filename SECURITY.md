# Security Policy

## Supported Versions

We support the following versions of the Catena SDK for Go:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0.0 | :x:                |

## Reporting a Vulnerability

We take the security of our software seriously. If you believe you have found a security vulnerability in the Catena SDK for Go, please report it to us as described below.

**Please do not report security vulnerabilities through public GitHub issues.**

### How to Report

Please email us at [security@catenaclearing.com](mailto:security@catenaclearing.com) with the following details:

- A description of the vulnerability.
- Steps to reproduce the issue.
- Any relevant code snippets or proof of concept.
- The version(s) of the SDK affected.

### Response Timeline

- We will acknowledge receipt of your report within 48 hours.
- We will provide an estimated timeline for fixing the vulnerability within 5 business days.
- We will notify you when a fix has been released.

### Disclosure Policy

We ask that you give us a reasonable amount of time to fix the vulnerability before disclosing it publicly. We will work with you to coordinate the public disclosure.

## Security Best Practices for Users

### Pinning Dependencies

Since this SDK is auto-generated from OpenAPI specifications, we strongly recommend pinning your dependency to a specific version in your `go.mod` file to ensure stability and security.

```go
require github.com/catenaclearing/catena-sdk-go v1.2.3
```

### Keeping Up to Date

We regularly update the SDK to reflect changes in the Catena API and to address any security issues. Please check for updates frequently and upgrade to the latest supported version.

### Generated Code

Please note that the code in the `gen/` directory is auto-generated. While we strive to ensure the security of the generator templates, the security of the generated code also depends on the security of the OpenAPI specifications provided by the Catena API.

## Third-Party Dependencies

This SDK may depend on third-party libraries. We monitor these dependencies for security vulnerabilities and update them as needed.
