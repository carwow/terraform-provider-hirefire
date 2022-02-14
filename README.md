# Terraform Provider HireFire

This provider is used to configure resources supported by the [HireFire API].

Documention at [Terraform registry].

[HireFire API]: https://docs.hirefire.io/
[Terraform registry]: https://registry.terraform.io/providers/carwow/hirefire/latest/docs


Development
---

    go build .

Then, create a `.terraformrc` file with:

    provider_installation {
      dev_overrides {
        "carwow/hirefire" = "<ABSOLUTE-PATH-TO-PROJECT>"
      }
    }

Finally, create a `main.tf` file (see `main.tf.example`) and call:

    terraform init
    HIREFIRE_API_KEY=your-key TF_CLI_CONFIG_FILE=.terraformrc terraform apply

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

Push a tag for the new version. CircleCI will do the rest.

[docs/index.md]: https://github.com/carwow/terraform-provider-hirefire/blob/main/docs/index.md


Contributing
---

[Pull requests] are very welcome!

Please report bugs in a [new issue].

Everyone is expected to follow the [code of conduct].

[Pull requests]: https://github.com/carwow/terraform-provider-hirefire/pulls
[new issue]: https://github.com/carwow/terraform-provider-hirefire/issues/new
[code of conduct]: https://github.com/carwow/terraform-provider-hirefire/tree/main/CODE_OF_CONDUCT.md
