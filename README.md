# StatusAPI

## Build

```
make build
```

## Run 

```
./cmd/api/api
```

Flags:

* `--http-addr` - web address of the service, default: :6660
* `--log-level` - logLevel (info, warn, debug, error), default: info

## Development

Generate code via `swagger.yaml`
```
make apiv1
```