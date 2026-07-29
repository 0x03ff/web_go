# web_go

A minimalist Go web server designed to pass ACME/SSL HTTP domain verification challenges with absolute minimum effort.

## Prerequisites

### Production Environment Installation

To compile `web_go` from source, ensure you have the Go toolchain installed:

```bash
sudo apt update
sudo apt install git golang-go -y
go version
```

### Deployment & Usage

Manual CA File Placement (ZeroSSL / Sectigo)

#### 1. Place Your Challenge File

Use the included helper script to easily add challenge files without dealing with manual directory creation or multi-line copy-paste shell errors:

Run the helper script in your terminal:

📦scripts
 ┣ 📜add-challenge-file.sh <--
 ┣ 📜compile.sh
 ┣ 📜renew-cert.sh
 ┗ 📜run.sh

```bash
 ./scripts/add-challenge-file.sh
```

1. Enter the target filename (e.g., 37B60801A77A091CF0AED6D1ECA6B65C.txt).
2. Paste the challenge payload into the editor prompt, copy the challenge content and paste on the editor.

The challenge file will automatically be created inside */tmp/acme-challenges/*, ready to be served.

#### 2. Configure Address and Port

Open `cmd/main.go` and adjust your network interface binding. Publicly exposing the application on the standard web port 80 will require root/sudo access:

```
	// Choose the running address here:
	// "127.0.0.1:8080"                  -> Local testing
	// "0.0.0.0:8080"                    -> Public on port 8080
	// "0.0.0.0:80"                      -> Public on standard HTTP port (Requires sudo)
	HTTP_ADDR = "127.0.0.1:8080"
```

Then use script to run server:

```bash
 ./scripts/run.sh
```

## Local Testing Instructions

Once your file is placed and `main.go` is updated, execute the local testing script to safely verify routing without writing permanent build artifacts to disk:

<pre id="tree-panel"><bold><span class="t-icon" name="icons">📦</span>scripts</bold><br/> ┣ <span class="t-icon" name="icons">📜</span>compile.sh<br/> ┗ <span class="t-icon" name="icons">📜</span>run.sh <-------</pre>

Open your web browser to verify the raw payload output:

```
http://127.0.0.1:8080/.well-known/pki-validation/37B60801A77A091CF0AED6D1ECA6B65C.txt( It depend)
```

---

### Automated Certbot Integration (Let's Encrypt)

If you are using Certbot for Let's Encrypt certificates, you can automate domain verification completely using standard HTTP hooks without manually editing files.

Run the issue script with your domain name:

```bash
sudo ./scripts/get-cert.sh example.com
```

This automatically passes the validation token to *web_go*, validates your domain, and cleans up the challenge file upon completion.

## Local Testing Instructions

Once your file is placed (or when testing the routing engine), execute the local development runner to safely verify routes without generating binary build artifacts on disk:

Start the server:

```bash
./scripts/run.sh
```

## Multi-Platform Compilation

When local verification passes, set your *HTTP_ADDR* constant to production preferences (*0.0.0.0:80*) and trigger the cross-compilation wizard:

```bash
./scripts/compile.sh
```

The wizard outputs stripped, static, binaries into the bin/ directory:

## Deployment

Transfer the single compiled asset to your target computer or remote cloud instance to solve domain challenges with zero external runtime dependencies.
