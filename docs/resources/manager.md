# hirefire_manager

Provides a Hirefire manager resource. This can be used to create, modify, and
delete managers from an application.

## Example Usage

```hcl
resource "hirefire_manager" "my_manager" {
  application_id = hirefire_application.my_application.id

  name    = "web"
  type    = "Manager::Web::Logplex::QueueTime"
  enabled = true
  minimum = 2
  maximum = 10

  aggregation            = "percentile"
  percentile             = 99
  minimum_queue_time     = 100
  maximum_queue_time     = 250
  downscale_quantity     = 1
  upscale_quantity       = 5
  downscale_sensitivity  = 2
  upscale_sensitivity    = 1
  downscale_timeout      = 1
  upscale_timeout        = 1
  upscale_on_initial_job = true
}
```

## Argument Reference

The following arguments are supported:

- `application_id` - (required) The ID of the application this manager belongs to.
- `name` - (required) The name of the manager. Must match the Procfile process name.
- `enabled` - (required) Whether to enable the manager.
- `minimum` - (required) The minimum amount of dynos to run.
- `maximum` - (required) The maximum amount of dynos to run.
- `notify` - (optional) Whether to send notifications. Default is `false`.
- `notify_quantity` - (optional) Notify when your dyno quantity exceeds this
  amount. Default is `0`.
- `notify_after` - (optional) Notify after the quantity has been exceeded for
  this amount of in minutes. Default is `0`.
- `type` - (required) The autoscaling strategy to use. Each strategy has a
  different set of additional fields documented below. Strategies:
  - [`Manager::Worker::HireFire::JobQueue`](#managerworkerhirefirejobqueue)
  - [`Manager::Worker::HireFire::JobLatency`](#managerworkerhirefirejoblatency)
  - [`Manager::Web::Logplex::Load`](#managerweblogplexload)
  - [`Manager::Web::Logplex::ResponseTime`](#managerweblogplexresponsetime)
  - [`Manager::Web::Logplex::ConnectTime`](#managerweblogplexconnecttime)
  - [`Manager::Web::Logplex::QueueTime`](#managerweblogplexqueuetime)
  - [`Manager::Web::Logplex::RPM`](#managerweblogplexrpm)
  - [`Manager::Web::NewRelic::V2::ResponseTime`](#managerwebnewrelicv2responsetime-and-v1)
  - [`Manager::Web::NewRelic::V2::RPM`](#managerwebnewrelicv2rpm-and-v1)
  - [`Manager::Web::NewRelic::V2::Apdex`](#managerwebnewrelicv2apdex-and-v1)
  - [`Manager::Web::NewRelic::V1::ResponseTime`](#managerwebnewrelicv2responsetime-and-v1) (legacy)
  - [`Manager::Web::NewRelic::V1::RPM`](#managerwebnewrelicv2rpm-and-v1) (legacy)
  - [`Manager::Web::NewRelic::V1::Apdex`](#managerwebnewrelicv2apdex-and-v1) (legacy)
  - [`Manager::Web::HireFire::ResponseTime`](#managerwebhirefireresponsetime) (legacy)

### Manager::Worker::HireFire::JobQueue
- `ratio` - Maintains a worker:job ratio by scaling based on the queue size.
- `decrementable` - Allows scaling down when there are still jobs in the queue. Default is `false`.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.
- `upscale_limit` - Limits the amount of dynos by which the manager scales up at a time. 0 means no limit.
- `downscale_limit` - Limits the amount of dynos by which the manager scales down at a time. 0 means no limit.

### Manager::Worker::HireFire::JobLatency
- `minimum_latency` - Scales down if your application's job latency goes below this amount in seconds.
- `maximum_latency` - Scales up if your application's job latency goes above this amount in seconds.
- `upscale_quantity` - The amount of dynos to scale up at a time.
- `downscale_quantity` - The amount of dynos to scale down at a time.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.
- `upscale_on_initial_job` - Ensures the availability of at least one Dyno when one or more jobs are enqueued, bypassing the initial maximum latency requirement. Default is `true`.

### Manager::Web::Logplex::Load
- `last_minutes` - The load average over the last n minutes. Possible values are `1`, `5`, and `15`.
- `minimum_load` - Scales down if your application's load goes below this amount (`100` equals load 1).
- `maximum_load` - Scales up if your application's load goes above this amount (`100` equals load 1).
- `upscale_quantity` - The amount of dynos to scale up at a time.
- `downscale_quantity` - The amount of dynos to scale down at a time.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.

### Manager::Web::Logplex::ResponseTime
- `aggregation` - Determines how the metric is aggregated. Possible values are `average` and `percentile`.
- `percentile` - The percentile to use when aggregation is set to `percentile`.
- `minimum_response_time` - Scales down if your application's response time goes below this amount in milliseconds.
- `maximum_response_time` - Scales up if your application's response time goes above this amount in milliseconds.
- `upscale_quantity` - The amount of dynos to scale up at a time.
- `downscale_quantity` - The amount of dynos to scale down at a time.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.

### Manager::Web::Logplex::ConnectTime
- `aggregation` - Determines how the metric is aggregated. Possible values are `average` and `percentile`.
- `percentile` - The percentile to use when aggregation is set to `percentile`.
- `minimum_connect_time` - Scales down if your application's connect time goes below this amount in milliseconds.
- `maximum_connect_time` - Scales up if your application's connect time goes above this amount in milliseconds.
- `upscale_quantity` - The amount of dynos to scale up at a time.
- `downscale_quantity` - The amount of dynos to scale down at a time.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.

### Manager::Web::Logplex::QueueTime
- `aggregation` - Determines how the metric is aggregated. Possible values are `average` and `percentile`.
- `percentile` - The percentile to use when aggregation is set to `percentile`.
- `minimum_queue_time` - Scales down if your application's queue time goes below this amount in milliseconds.
- `maximum_queue_time` - Scales up if your application's queue time goes above this amount in milliseconds.
- `upscale_quantity` - The amount of dynos to scale up at a time.
- `downscale_quantity` - The amount of dynos to scale down at a time.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.

### Manager::Web::Logplex::RPM
- `ratio` - Maintains a web:rpm ratio by scaling based on requests per minute.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.

### Manager::Web::NewRelic::V2::ResponseTime (and V1)
- `minimum_response_time` - Scales down if your application's response time goes below this amount in milliseconds.
- `maximum_response_time` - Scales up if your application's response time goes above this amount in milliseconds.
- `upscale_quantity` - The amount of dynos to scale up at a time.
- `downscale_quantity` - The amount of dynos to scale down at a time.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.
- `scale_up_on_503` - Scale up when your application returns HTTP/503 (Serivce Unavailable). Default is `false`.
- `new_relic_api_key` - New Relic api key.
- `new_relic_app_id` - New Relic application id.
- `new_relic_account_id` - V1 only. New Relic account id.

### Manager::Web::NewRelic::V2::RPM (and V1)
- `ratio` - Maintains a web:rpm ratio by scaling based on requests per minute.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.
- `new_relic_api_key` - New Relic api key.
- `new_relic_app_id` - New Relic application id.
- `new_relic_account_id` - V1 only. New Relic account id.

### Manager::Web::NewRelic::V2::Apdex (and V1)
- `minimum_apdex` - Scales up if your application's apdex goes below this amount.
- `maximum_apdex` - Scales down if your application's apdex goes above this amount.
- `upscale_quantity` - The amount of dynos to scale up at a time.
- `downscale_quantity` - The amount of dynos to scale down at a time.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.
- `scale_up_on_503` - Scale up when your application returns HTTP/503 (Serivce Unavailable). Default is `false`.
- `new_relic_api_key` - New Relic api key.
- `new_relic_app_id` - New Relic application id.
- `new_relic_account_id` - V1 only. New Relic account id.

### Manager::Web::HireFire::ResponseTime
- `url` - Specific URL to measure.
- `minimum_response_time` - Scales down if your application's response time goes below this amount in milliseconds.
- `maximum_response_time` - Scales up if your application's response time goes above this amount in milliseconds.
- `upscale_quantity` - The amount of dynos to scale up at a time.
- `downscale_quantity` - The amount of dynos to scale down at a time.
- `upscale_sensitivity` - The amount of threshold breaches to wait before scaling up.
- `downscale_sensitivity` - The amount of threshold breaches to wait before scaling down.
- `upscale_timeout` - The amount of minutes to wait before performing upscale operations.
- `downscale_timeout` - The amount of minutes to wait before performing downscale operations.
- `scale_up_on_503` - Scale up when your application returns HTTP/503 (Serivce Unavailable). Default is `false`.

## Attribute Reference

The following attributes are exported:

- `id` - The ID of the manager assigned by HireFire.

## Import

Managers be imported using the `id`, e.g.

```bash
terraform import hirefire_manager.my_manager 516a2b53-04c8-424e-8533-99d47ef1f9bf
```
