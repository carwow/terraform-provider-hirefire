# hirefire_application

Provides a HireFire application resource. This can be used to create, modify,
and delete applications from an account.

## Example Usage

```hcl
resource "hirefire_application" "my_application" {
  account_id = hirefire_account.my_account.id

  name                  = heroku_app.my_application.name
  custom_domain         = heroku_domain.my_application.hostname
  logplex_drain_token   = heroku_drain.my_application.token
  ssl                   = true
  restart_crashed_dynos = true
}
```

## Argument Reference

The following arguments are supported:

- `account_id` - (required) The ID of the account this application belongs to.
- `name` - (required) The name of the application as registered with Heroku.
- `custom_domain` - (optional) The custom domain of the Heroku application, if
  any is being used.
- `logplex_drain_token` - (optional) The Heroku drain token, if using Logplex
  metrics.
- `ssl` - (optional) Use `https`. Default is `false`.
- `restart_crashed_dynos` - (optional) Restart dynos in a crashed state.
  Default is `false`.
- `new_issue_notifications` - (optional) Notify about new issues. Default is
  `false`.
- `resolved_issue_notifications` - (optional) Notify about resolved issues.
  Default is `false`.

## Attribute Reference

The following attributes are exported:

- `id` - The ID of the application assigned by HireFire.
- `token` - The token of the application assigned by HireFire.

## Import

Applications can be imported using the `id`, e.g.

```bash
terraform import hirefire_application.my_application 516a2b53-04c8-424e-8533-99d47ef1f9bf
```
