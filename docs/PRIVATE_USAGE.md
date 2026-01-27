# Private Module Usage Guide

This SDK is a **private Go module**. Because it is not published to the public Go module proxy (proxy.golang.org), you must configure your Go environment to access it directly from GitHub.

## 1. Configure GOPRIVATE

Tell Go to skip the public proxy for this repository:

```bash
go env -w GOPRIVATE=github.com/Chalupa-Tech/*
```

## 2. Configure Authentication

Since the repository is private, Go needs credentials to fetch it.

### Option A: SSH (Recommended for Developers)

If you have your SSH keys set up with GitHub:

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

### Option B: Personal Access Token (PAT) (Recommended for CI/CD)

If you are using a Personal Access Token (e.g., in CI environments):

#### Required Scopes
- **Fine-grained PAT (Recommended)**:
  - **Repository access**: Select the specific repositories (e.g., `go-schwab-api-individual`, `tayvens-stock-report`).
  - **Permissions**: `Contents` -> `Read-only`.
- **Classic PAT**:
  - `repo` (Full control of private repositories).

```bash
git config --global url."https://<YOUR_USER>:<YOUR_PAT>@github.com/".insteadOf "https://github.com/"
```

## 3. Usage in consumers

In your `go.mod`:

```go
require github.com/Chalupa-Tech/go-schwab-api-individual v0.1.0
```

Run `go mod tidy` to fetch the module using your configured credentials.
