# FOSS Go 101

Slide deck: **Go 101 — A Practical Introduction for University Students**, built with the Go [`present`](https://pkg.go.dev/golang.org/x/tools/cmd/present) tool. Sources live under [`present/`](present/).

## Requirements

- [Go](https://go.dev/dl/) **1.25+** (see [`go.mod`](go.mod))

## Run locally

From the repository root:

```bash
go install golang.org/x/tools/cmd/present@latest
present -content ./present
```

Open the URL printed in the terminal (by default `http://127.0.0.1:3999`), then open [`go101.slide`](present/go101.slide).

To run editable code snippets against **play.golang.org** instead of a local WebSocket (useful if you do not need local execution):

```bash
present -content ./present -use_playground=true
```

## Classroom note

For student practice, check:

- [`present/project/todo/`](present/project/todo/) for the level-by-level mini project

## Publish on GitHub Pages

The workflow [`.github/workflows/pages.yml`](.github/workflows/pages.yml) exports the deck to static HTML and deploys it with GitHub Actions.

1. In the GitHub repo, go to **Settings → Pages**.
2. Under **Build and deployment**, set **Source** to **GitHub Actions**.
3. Push to **`main`** (or run the workflow manually via **Actions → Deploy slides to GitHub Pages → Run workflow**).

After a successful run, the site URL appears on the workflow run and on the Pages settings page. For a **project** site (`https://<user>.github.io/<repo>/`), the workflow rewrites asset paths so `/static` and `/play.js` resolve under the repo prefix.

## Layout

| Path | Purpose |
|------|---------|
| [`present/go101.slide`](present/go101.slide) | Main slide file |
| [`present/code/`](present/code/) | Example `.go` sources and related files referenced from the deck |

## Module path

[`go.mod`](go.mod) uses the placeholder module path `github.com/example/foss-go-101`. Replace it with your real module path if you publish this as its own module or import it elsewhere.
