# hirefire_organization

Provides a HireFire organization resource. This can be used to create, modify,
and delete organizations.

## Example Usage

```hcl
resource "hirefire_organization" "my_organization" {
  name      = "My Organization"
  time_zone = "UTC"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (required) The name of the organization.
- `time_zone` - (required) The time zone associated with the organization.
(Verify time zone name on list provided by HireFire UI.)

## Attribute Reference

The following attributes are exported:

- `id` - The ID of the organization assigned by HireFire.

## Import

Organizations can be imported using the `id`, e.g.

```bash
terraform import hirefire_organization.my_organization 516a2b53-04c8-424e-8533-99d47ef1f9bf
```
