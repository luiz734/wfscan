# wfscan

Find large files in your file system. Works on Windows and Linux.

This tool is not optimized for performance. If speed is important to you, Linux provides many more suitable tools out of the box.

## Installation

Clone the repo and compile it. All dependencies are managed by Go.

Alternatively, you can download the precompiled release.

## Examples (Windows)

- List files larger than 1GB:
```powershell
wfscan --path=C:\ --larger=1GB
```

- List files larger than 1GB but smaller than 2GB in the Downloads folder:
```powershell
wfscan --path=C:\Users\username\Downloads --larger=1GB --smaller=2GB
```
- List only file sizes, limiting output to 5 entries:
```powershell
wfscan --path=C:\User --entries=5 --format="s"
```

## Examples (Linux)
- List files smaller than 1MB in /home, skipping the first 10 results:
```bash
wfscan --path=/home --smaller=1MB --skip=10
```

- List all files larger than 200MB in /, displaying all error messages:
```bash
wfscan --path=/home --larger=200MB --errors
```
