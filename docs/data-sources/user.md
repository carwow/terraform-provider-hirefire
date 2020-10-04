# hirefire_user

Use this data source to get the email addrerss of a Hirefire user.

## Example Usage

```hcl
data "hirefire_user" "my_user" {
  id = var.hirefire_user_id
}
```

## Argument Reference

The following arguments are supported:

- `id` - (required) The ID of the user assigned by HireFire.

## Attribute Reference

The following attributes are exported:

- `email` - The email address of the user.
- `notifications` - Whether the user has notifications enabled.
