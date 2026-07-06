## 2025-07-06 - [Git Argument Injection Protection]
**Vulnerability:** Git command argument injection (flag injection) when wrapping CLI tools.
**Learning:** Even when using separate arguments (not shell execution), positional arguments starting with a hyphen can be interpreted as flags (e.g., --upload-pack).
**Prevention:** Implement a helper to reject hyphen-prefixed arguments for all user-controlled inputs passed to external CLI commands. Use the -- separator where supported by the underlying tool.
