# Timer

A small desktop app showing the current Unix timestamp and human-readable time, with a Unix ↔ Date converter.

## Build

Requires [Go](https://go.dev/) and a C compiler (GCC via [WinLibs](https://winlibs.com/) recommended on Windows).

```powershell
$env:CGO_ENABLED=1; $env:CC="C:\path\to\gcc.exe"; go build -ldflags "-H windowsgui" -o timer.exe .
```

Replace `C:\path\to\gcc.exe` with the actual path to your GCC binary, for example:

```
C:\Users\<you>\AppData\Local\Microsoft\WinGet\Packages\BrechtSanders.WinLibs.POSIX.UCRT_Microsoft.Winget.Source_8wekyb3d8bbwe\mingw64\bin\gcc.exe
```

The `-H windowsgui` flag suppresses the terminal window when launching the app.

## Run tests

```powershell
$env:CGO_ENABLED=1; $env:CC="C:\path\to\gcc.exe"; go test ./... -v
```
