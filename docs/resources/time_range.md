# hirefire_time_range

Provides a Hirefire time range resource. This can be used to create, modify,
and delete time ranges from a manager.

## Example Usage

```hcl
resource "hirefire_time_range" "my_time_range" {
  manager_id = hirefire_manager.my_manager.id

  email = "example@example.com"
  from_minute  = 415 // 06:55
  until_minute = 435 // 07:15
  minimum      = 5
  maximum      = 10
  monday       = true
  tuesday      = true
  wednesday    = true
  thursday     = true
  friday       = true
}
```

## Argument Reference

The following arguments are supported:

- `manager_id` - (required) The ID of the manager this time range belongs to.
- `from_minute` - (required) The minute of the day that this time range should
  start applying (min 0, max 1339).
- `until_minute` - (required) The minute of the day that this time range should
  start applying (min 0, max 1339).
- `minimum` - (required) The minimum amount of dynos to run.
- `maximum` - (required) The maximum amount of dynos to run.
- `position` - (optional) The position of the time range (lower numbers take
  precedence over higher numbers with overlapping days/hours. Default is `0`.
- `monday` - (optional) Apply this time range on monday. Default is `false`.
- `tuesday` - (optional) Apply this time range on tuesday. Default is `false`.
- `wednesday` - (optional) Apply this time range on wednesday. Default is `false`.
- `thursday` - (optional) Apply this time range on thursday. Default is `false`.
- `friday` - (optional) Apply this time range on friday. Default is `false`.
- `saturday` - (optional) Apply this time range on saturday. Default is `false`.
- `sunday` - (optional) Apply this time range on sunday. Default is `false`.

## Attribute Reference

The following attributes are exported:

- `id` - The ID of the time range assigned by HireFire.

## Import

Time ranges can be imported using the `id`, e.g.

```bash
terraform import hirefire_time_range.my_time_range 516a2b53-04c8-424e-8533-99d47ef1f9bf
```
