# web_go

A minimalist Go web server designed to pass ACME/SSL HTTP domain verification challenges with absolute minimum effort.

## Prerequisites

### Production Environment Installation

To compile and run this service, only need the Go toolchain installed:

```
sudo apt update
sudo apt install git golang-go -y
go version
```

## Development

### 1. Place Your Challenge File

Place the validation text file provided by your Certificate Authority (CA) directly inside the `cmd/` directory:

Modify the exist address and port if need, public access with port 80 require the sudo:

### 2. Configure Address and Port

Open `cmd/main.go` and adjust your network interface binding. Publicly exposing the application on the standard web port 80 will require root/sudo access:

```
	// Choose the running address here:
	// "127.0.0.1:8080"                  -> Local testing
	// "0.0.0.0:8080"                    -> Public on port 8080
	// "0.0.0.0:80"                      -> Public on standard HTTP port (Requires sudo)
	HTTP_ADDR = "0.0.0.0:8080"

```

### 3. Update the Embed Directives

Update the configuration constants and compiler directives in `cmd/main.go` to match your exact file name. This bakes the file contents straight into the compiled binary:

```

// The precise validation filename provided by Certificate Authority (CA) like ZeroSSL
TEXT_NAME = "example.txt" , update as file name:

//go:embed challenge.txt
var validationFileContents string

```

### 4. Update the URL routing path if need

*The URL routing path is managed dynamically in `cmd/api.go` relative to configuration settings:*

`routePath := "/.well-known/pki-validation/" + app.config.textName`
`routePath := "/.well-known/acme-challenge/" + app.config.textName`


---

## Local Testing Instructions

Once your file is placed and `main.go` is updated, execute the local testing script to safely verify routing without writing permanent build artifacts to disk:

<pre id="tree-panel"><bold><span class="t-icon" name="icons">📦</span>scripts</bold><br/> ┣ <span class="t-icon" name="icons">📜</span>compile.sh<br/> ┗ <span class="t-icon" name="icons">📜</span>run.sh <-------</pre>

Open your web browser or curl to verify the raw payload output:( It depend)

```
http://127.0.0.1:8080/.well-known/pki-validation/37B60801A77A091CF0AED6D1ECA6B65C.txt
http://127.0.0.1:8080/.well-known/acme-challenge/37B60801A77A091CF0AED6D1ECA6B65C.txt
```

---

## Multi-Platform Compilation

When local verification passes, change your `HTTP_ADDR` constant to production preferences (`0.0.0.0:80`) or keep fefault with port 80 request forward port 8080 and trigger the cross-compilation wizard:

<pre id="tree-panel"><bold><span class="t-icon" name="icons">📦</span>scripts</bold><br/> ┣ <span class="t-icon" name="icons">📜</span>compile.sh <-----<br/> ┗ <span class="t-icon" name="icons">📜</span>run.sh</pre>

The interactive wizard outputs target-specific assets cleanly into your `bin/` directory:

<pre id="tree-panel"><bold><span class="t-icon" name="icons">📦</span>bin</bold><br/> ┣ <span class="t-icon" name="icons">📜</span>.DS_Store<br/> ┣ <span class="t-icon" name="icons">📜</span>.gitkeep<br/> ┗ <span class="t-icon" name="icons">📜</span>main-linux-amd64 <-----</pre>


## Deployment

Because this architecture utilizes go:embed and compiles with CGO_ENABLED=0, the target binary is self-contained. It only need to transfer that single binary asset to other computer or remote cloud instance to solve domain challenge.
