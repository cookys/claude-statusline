# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability, please report it by:

1. **Do NOT** open a public GitHub issue
2. Email the maintainer directly or use GitHub's private vulnerability reporting
3. Include detailed information about the vulnerability
4. Allow reasonable time for a fix before public disclosure

## Security Features

This project includes several security measures:

- **Signed Releases**: All release artifacts are signed with [Cosign](https://github.com/sigstore/cosign)
- **SLSA Provenance**: Releases include SLSA Level 3 provenance attestations
- **SBOM**: Software Bill of Materials provided in SPDX format
- **Checksums**: SHA256 checksums for all release artifacts

## Verifying Releases

```bash
# Verify checksum
sha256sum -c checksums.txt

# Verify Cosign signature
cosign verify-blob \
  --signature <file>.sig \
  --certificate <file>.pem \
  --certificate-identity-regexp 'https://github.com/kevinlincg/claude-statusline/.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  <file>
```

## Dependencies

This project uses only the Go standard library, minimizing supply chain risks.
