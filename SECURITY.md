# Security Policy

This is the security policy for Gamma Web Server.

## Supported Versions

### Stable

| Version     | Supported          | Date Released             |
| ----------- | ------------------ | ------------------------- |
| 1.5.0 (LTS) | :white_check_mark: | November 16th, 2025       |
| 1.4.1       | :x:                | January 13th, 2024        |
| 1.4 (LTS)   | :white_check_mark: | January 12th, 2024        |
| 1.3         | :x:                | June 29th, 2023           |
| 1.2         | :x:                | June 27th, 2023 (e673abe) |
| 1.1         | :x:                | December 24th, 2022       |
| 1.0         | :x:                | December 23rd, 2022       |

### Unstable

| Version      | Supported | Date Released      |
| ------------ | --------- | ------------------ |
| 1.4.0-beta.2 | :x:       | November 6th, 2023 |
| 1.4.0-beta.1 | :x:       | October 20th, 2023 |
| 1.3.0-rc1/2  | :x:       | June 28th, 2023    |

<!-- No unstable versions before v1.3.0-rc1 -->

## Version Support Policy

We are committed to maintaining the security and stability of Gamma Web Server. Our version support policy is as follows:

- **Latest Stable Release**: Always receives security updates and bug fixes until the next stable release.
- **LTS (Long-Term Support)**: Designated versions receive extended security support for **24 months** from release date.
- **Previous Stable Release**: Receives security updates for critical vulnerabilities for **6 months** after the next stable release.
- **Older Versions**: No longer supported once EOL is reached. Users are strongly encouraged to upgrade to the latest stable or LTS version.
- **Unstable/Beta/RC Versions**: Not supported for production use. Reach EOL immediately upon stable release. Security updates are not guaranteed.

### End-of-Life (EOL) Schedule

| Version     | Release Date      | EOL Date                | Reason                   |
| ----------- | ----------------- | ----------------------- | ------------------------ |
| 1.5.0 (LTS) | November 16, 2025 | November 16, 2027       | 24 months from release   |
| 1.4.1       | January 13, 2024  | November 16, 2025 (EOL) | Superseded by v1.5.0 LTS |
| 1.4 (LTS)   | January 12, 2024  | January 12, 2026        | 24 months from release   |
| 1.3         | June 29, 2023     | January 13, 2024 (EOL)  | Superseded by v1.4.1     |
| 1.2         | June 27, 2023     | December 29, 2023 (EOL) | Superseded by v1.3       |

**Note**: EOL dates are approximate and subject to change.

### Discontinuation Notice

When a version reaches end-of-life (EOL) and is no longer supported:

- It will be marked with :x: in the supported versions table above
- No further security patches or updates will be provided
- Users should migrate to a supported version as soon as possible
- Known vulnerabilities in discontinued versions will be documented but not patched

## Reporting a Vulnerability

We take the security of our software seriously and appreciate the efforts of the community in improving it. If you discover a security vulnerability, we kindly request that you follow our responsible disclosure process.

To report a security vulnerability, please follow these steps:

1. **Do not** publicly disclose the vulnerability or related details.
2. Send an email to recon at [recon@mail.recon.best](mailto:recon@mail.recon.best) with a detailed description of the vulnerability. Please include information such as the impact and potential exploit scenarios.
3. The email should be acknowledged within 2-3 days and recon may request additional information or clarification if needed.
4. We aim to respond to vulnerability reports and provide updates on the progress of addressing them in a timely manner.
5. If desired, you will be publicly acknowledged for your contribution once the vulnerability is resolved.

## Responsible Disclosure Policy

We kindly request that you adhere to responsible disclosure practices when reporting security vulnerabilities:

- Refrain from publicly disclosing the vulnerability before it has been resolved.
- Avoid actively exploiting the vulnerability or performing any malicious actions.
- Do not modify, access, or delete any data without explicit authorization.

We appreciate your dedication to protecting the security and integrity of our software and its users. As a token of our gratitude, we may consider providing rewards or recognition for responsibly disclosed security vulnerabilities, subject to our discretion.

Thank you for contributing to the security of Gamma Web Server.
