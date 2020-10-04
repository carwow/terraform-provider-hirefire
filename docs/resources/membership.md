# hirefire_membership

Provides a HireFire user membership resource. This can be used to create,
modify, and delete users memberships from an organization.

## Example Usage

```hcl
resource "hirefire_membership" "my_membership" {
  organization_id = hirefire_organization.my_organization.id

  user_id = var.hirefire_user_id
  owner   = true
}

resource "hirefire_organization" "my_organization" {
  name      = "My Organization"
  time_zone = "UTC"
}
```

## Argument Reference

The following arguments are supported:

- `organization_id` - (required) The ID of the organization this membership
  belongs to.
- `user_id` - (required) The ID of the user associated with this membership.
- `owner` - (optional) Whether the user is an owner of the organization.
  Default is `false`.

## Attribute Reference

The following attributes are exported:

- `id` - The ID of the membership assigned by HireFire.

## Import

Memberships can be imported using the `id`, e.g.

```bash
terraform import hirefire_membership.my_membership 516a2b53-04c8-424e-8533-99d47ef1f9bf
```
