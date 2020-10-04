# Terraform Provider HireFire

This provider is used to configure resources supported by the [HireFire API].

Documention at [Terraform registry].

[HireFire API]: https://docs.hirefire.io/
[Terraform registry]: https://registry.terraform.io/providers/carwow/hirefire/latest/docs


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

1. Update version in [docs/index.md] if necessary and commit.
2. Push a tag for the new version. CircleCI will do the rest.

[docs/index.md]: https://github.com/carwow/terraform-provider-hirefire/blob/master/docs/index.md


Contributing
---

[Pull requests] are very welcome!

Please report bugs in a [new issue].

Everyone is expected to follow the [code of conduct].

[Pull requests]: https://github.com/carwow/terraform-provider-hirefire/pulls
[new issue]: https://github.com/carwow/terraform-provider-hirefire/issues/new
[code of conduct]: https://github.com/carwow/terraform-provider-hirefire/tree/master/CODE_OF_CONDUCT.md
