# hirefire_account

Provides a HireFire Account resource. This can be used to create and delete
accounts from an organization.

-> This account must still be authorized in order to work with the Heroku
Platform API. See [HireFire documentation](https://docs.hirefire.io/api/account/create/).

## Example Usage

```hcl
resource "hirefire_account" "my_account" {
  organization_id = hirefire_organization.my_organization.id
}

resource "hirefire_organization" "my_organization" {
  name      = "My Organization"
  time_zone = "UTC"
}
```

## Argument Reference

The following arguments are supported:

- `organization_id` - (required) The ID of the organization this account
  belongs to.

## Attribute Reference

The following attributes are exported:

- `id` - The ID of the account assigned by HireFire.

## Import

Accounts can be imported using the `id`, e.g.

```bash
terraform import hirefire_account.my_account 516a2b53-04c8-424e-8533-99d47ef1f9bf
```
