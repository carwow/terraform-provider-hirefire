# HireFire Provider

The HireFire provider is used to configure resources supported by the
[HireFire API].

[HireFire API]: https://docs.hirefire.io/

## Example Usage

```hcl
provider "hirefire" {
  version = "~> 0.3"
  api_key = "${var.hirefire_api_key}"
}
```

## Argument Reference

The following arguments are supported:

- `api_key` - (required) This is the HireFire API key. This can also be
  specified with the `HIREFIRE_API_KEY` environment variable.
