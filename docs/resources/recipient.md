# hirefire_recipient

Provides a Hirefire notification recipient resource. This can be used to
create, modify, and delete notification recipients from an application.

## Example Usage

```hcl
resource "hirefire_recipient" "my_recipient" {
  application_id = hirefire_application.my_application.id

  email = "example@example.com"
}
```

## Argument Reference

The following arguments are supported:

- `application_id` - (required) The ID of the application this recipient belongs to.
- `email` - (required) An email address.

## Attribute Reference

The following attributes are exported:

- `id` - The ID of the recipient assigned by HireFire.

## Import

Recipients can be imported using the `id`, e.g.

```bash
terraform import hirefire_recipient.my_recipient 516a2b53-04c8-424e-8533-99d47ef1f9bf
```
