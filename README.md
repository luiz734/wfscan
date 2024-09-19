# wfscan
Find large files in you file system. Works on Windows and Linux.

This tool is not performatic by any means. If speed is important to you, linux has a lot of more suitable tools out of the box.

## Installation
Clone the repo and compile it. All dependencies are handled by go.

## Examples windows
- List files larger than 1GB
```powershell
wfscan --path=C:\ --larger=1GB
```

- List files larger than 1GB, smaller than 2GB on folder Downloads
```powershell
wfscan --path=C:\Users\username\Downloads --larger=1GB --smaller=2GB
```

- List only file size, limiting output to 5 entries
```powershell
wfscan --path=C:\User --entries=5 --format="s"
```

## Examples linux
- List files smaller than 1MB in `/home`, skipping the first 10 results
```bash
wfscan --path=/home --smaller=1MB --skip=10
```

- List all files larger than 200MB in `/`, showing all error messages
```bash
wfscan --path=/home --larger=200MB --skip=10 --errors
```

