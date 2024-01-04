# Go HTTP Checker

## Build
```
git clone https://github.com/stobbsm/gohttpcheck
cd gohttpcheck
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static" -w'
```
Binary is called gohttpcheck. It is static, and can be included in any docker
container to facilitate health checks when they otherwise aren't available.

## Usage
```
gohttpcheck http://<HOST>:<PORT>/{ENDPOINT}
```
Where:
- HOST = either the IP or resolvable hostname to connect to
- PORT = the port number if non-standard
- ENDPOINT = _optional_, the path to check

All Messages are printed to Stderr.
Exit code 1 means it was unsuccessful.
Exit code 0 means it was successful.

If not successful, the error is printed on Stderr.
If successful, you will get a message of `status: <HTTP STATUS CODE>`, where
HTTP STATUS CODE is usually going to be 200.

## Caveats
The scope of this simple application is to check http endpoints only, not https,
used inside of containers.
It does not and will not support https. It may work for now, but I can't
garuntee it will work forever.

## TODO:
- Create actions workflow to create releases based on the main branch.
- The created binary is still to large, and since it really just need to make a
  TCP call and get the statuscode from the initial set of headers, I will
  eventually make this use a raw TCP call.
