#### Please place the text file here:

# Process Log

This document captures the daily activities, decisions,

and reflections for building and deploying our Spring Boot cloud
computing service.

Follow the template below to document your activities, decisions, and reflections for each day of the week.

## YYYY-MM-DD | Short Title

Intent:

Action:
Result:
Decision / Interpretation:
Next:

---

## 2026-05-27 | Repository Initialization & Core Validation Service

### Intent

To complete automated SSL/TLS domain validation, the Certificate Authority (CA) requires an HTTP challenge file to be served over a highly specific URL path. Modifying our main existing web production servers to route this temporary request can be invasive and complex. The intent was to build a zero-dependency, ultra-minimal HTTP microservice dedicated exclusively to serving this challenge.

### Action

* **Simplified Architecture:** Removed the external third-party router (`chi`) and configuration framework (`godotenv`) in favor of Go's built-in `net/http` package to reduce memory footprint and strip out maintenance bloat.
* **Single-File Portability:** Leveraged Go's native standard library `go:embed` directive to bake the CA validation text file directly into the executable binary at compile time.
* **Automation Scripts:** * Authored `run.sh` to handle rapid local building and sudo-privileged execution.
  * Authored a unified, interactive `compile.sh` cross-compilation wizard using standard Bash prompts to output tailored builds across multi-platform targets.

### Result

* The project structure was successfully reduced to just two source files (`main.go`, `api.go`) and the validation payload.
* Generated a self-contained web server that requires zero directory structure or companion text files on deployment targets.
* Successfully generated cross-compiled binaries targeting Linux, macOS, and Windows across both `amd64` and `arm64` computer architectures.

### Decision / Interpretation

Using the Go standard library rather than pulling down framework dependencies kept our artifact compact and highly secure. Baking the file asset directly into the binary memory space avoids runtime file-system disk I/O latency, making the verification mechanism virtually immune to classic file-not-found (`404`) path-misconfiguration issues during deployment.

### Next

* Finalize validation testing natively on an Apple Silicon macOS environment (`darwin/arm64`).
* Transfer the static `cgo`-disabled Linux binary (`main-linux-amd64`) directly to the production cloud infrastructure for live validation execution.

## 2026-07-30 | Improve architecture on Dynamic File Serving

### Intent

Compare the initial compile-time embedded model (`go:embed`) with file-system-based validation server (`/tmp/acme-challenges/` + Certbot hook scripts)  to determine the optimal deployment strategy for SSL/TLS verification workflow.

### Action

* The embed version focus on manual CA verification (ZeroSSL/Sectigo while new version mainly automated Let's Encrypt renewals via Certbot.

### Result

* **Dynamic File Model (`from`):** Offers superior long-term automation then orignal embed design, while remain support ofmanual CA verification with script

### Decision / Interpretation

Keep both the **embedded model (`go:embed`)** and  **one-off manual CA verifications** due to its  better for fully automated recurring renewals (Let's Encrypt) with cert bot

### Next

* Testing the work flow with Certbot
