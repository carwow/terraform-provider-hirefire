# Terraform Provider HireFire

This provider is used to configure resources supported by the [HireFire API].

[HireFire API]: https://docs.hirefire.io/


Install
---

Download from [releases] page and install in [Terraform's plugin directory].

For example, change `x.x.x` with latest version:

    cd ~/.terraform.d/plugins
    VERSION=x.x.x
    test -e terraform-provider-hirefire_v${VERSION} && exit || true
    wget -O terraform-provider-hirefire_v${VERSION} https://github.com/carwow/terraform-provider-hirefire/releases/download/v${VERSION}/terraform-provider-hirefire_v${VERSION}_linux_amd64
    chmod +x terraform-provider-hirefire_v${VERSION}

Then subscribe to Releases only notifications to be alerted of new releases.

[releases]: https://github.com/carwow/terraform-provider-hirefire/releases
[Terraform's plugin directory]: https://www.terraform.io/docs/configuration/providers.html#third-party-plugins


Usage
---

    provider "hirefire" {
      version = "~> 0.1"
      api_key = "your-key" // or set environment variable HIREFIRE_API_KEY
    }

*For resources documentation, see the code under [resources] directory for now.*

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

and run with:

    env $(cat .env) go test -v ./...

or to run tests for a single resource:

    env $(cat .env) go test -v ./resources/application


Release
---

Using GitHub's [hub] command, change `vX.X.X` to the appropriate version:

    VERSION=vX.X.X # Set the version you intent to release
    env GOOS=linux GOARCH=amd64 go build -o terraform-provider-hirefire_${VERSION}_${GOOS}_${GOARCH}
    hub release create $VERSION

Then upload the file(s) to the GitHub's release that was just created.

[hub]: https://github.com/github/hub


Contributing
---

[Pull requests] are very welcome!

Please report bugs in a [new issue].

Everyone is expected to follow the [code of conduct].

[Pull requests]: https://github.com/carwow/terraform-provider-hirefire/pulls
[new issue]: https://github.com/carwow/terraform-provider-hirefire/issues/new
[code of conduct]: https://github.com/carwow/terraform-provider-hirefire/tree/master/CODE_OF_CONDUCT.md
