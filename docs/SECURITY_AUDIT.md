# Security Threat Model & Audit Report

* Zero persistence of user tokens.
* Immediate token revocation after single in-memory evaluation.
* Strict Content Security Policy (CSP) forbidding unsafe-inline execution.
