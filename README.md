# Terraform Provider HireFire

This provider is used to configure resources supported by the [HireFire API].

[HireFire API]: https://docs.hirefire.io/


Install
---

Download from [releases] page and install in [Terraform's plugin directory].

The [install.sh] script will do this for you. Download the file, change `x.x.x`
to the latest version, and run it like:

    ./install.sh x.x.x

Then subscribe to Releases only notifications to be alerted of new releases.

[releases]: https://github.com/carwow/terraform-provider-hirefire/releases
[Terraform's plugin directory]: https://www.terraform.io/docs/configuration/providers.html#third-party-plugins
[install.sh]: https://github.com/carwow/terraform-provider-hirefire/blob/master/install.sh


Usage
---

    provider "hirefire" {
      version = "~> 0.1"
      api_key = "your-key" // or set environment variable HIREFIRE_API_KEY
    }

For resources documentation, see the code under [resources] directory, and
consult the [HireFire API] documenation for more details.

[resources]: https://github.com/carwow/terraform-provider-hirefire/tree/master/resources


Development
---

Requires [Go] 1.14.x.

    go build .

Then create a `main.tf` file and use Terraform as usual to experiment.

This project uses [Go Modules] for dependency management.

[Go]: https://golang.org/doc/install
[Go Modules]: https://github.com/golang/go/wiki/Modules


Testing
---

    go test -v ./...

In order to run Terraform acceptance tests which will create **real** resources
then create an `.env` file with:

    TF_ACC=on
    HIREFIRE_API_KEY=your-key
    # The following variables are optional.
    # For user data source:
    HIREFIRE_TEST_USER_ID=user-id
    HIREFIRE_TEST_USER_EMAIL=user-email
    # For membership resource:
    HIREFIRE_TEST_MEMBERSHIP_USER_ID=another-user-id

and run with:

    env $(cat .env) go test -v ./...

or to run tests for a single resource:

    env $(cat .env) go test -v ./resources/application


Release
---

Change `x.x.x` to the appropriate version, and build the binaries by running:

    rake release[x.x.x]

If you have GitHub's [hub] command, a GitHub release will be created for you.
Otherwise, create the GitHub release.

Then upload the binaries to the GitHub release. You **always** need to upload
the binaries, even if the release was created by hub.

[hub]: https://github.com/github/hub


Contributing
---

[Pull requests] are very welcome!

Please report bugs in a [new issue].

Everyone is expected to follow the [code of conduct].

[Pull requests]: https://github.com/carwow/terraform-provider-hirefire/pulls
[new issue]: https://github.com/carwow/terraform-provider-hirefire/issues/new
[code of conduct]: https://github.com/carwow/terraform-provider-hirefire/tree/master/CODE_OF_CONDUCT.md
