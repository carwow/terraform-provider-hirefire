# Terraform Provider HireFire

This provider is used to configure resources supported by the [HireFire API].

[HireFire API]: https://docs.hirefire.io/


Requirements
---

- [Terraform] 0.12.x
- [Go] 1.13 (to build the provider plugin)

[Terraform]: https://www.terraform.io/downloads.html
[Go]: https://golang.org/doc/install


Usage
---

```
provider "hirefire" {
  version = "~> 0.1"
  api_key = "your-key" // or set environment variable HIREFIRE_API_KEY
}
```


Development
---

```
go build .
```

Then create a `main.tf` file and use Terraform as usual to experiment.

This project uses [Go Modules] for dependency management.

[Go Modules]: https://github.com/golang/go/wiki/Modules


Testing
---

```
go test -v ./...
```

In order to run Terraform acceptance tests which will create **real** resources
then create an `.env` file with:
```
TF_ACC=on
HIREFIRE_API_KEY=your-key
```

and run with:

```
env $(cat .env) go test -v ./...
env $(cat .env) go test -v ./resources/application # to run tests for a single resource
```


Release
---

Using GitHub's [hub] command, change `v0.X.X` to the appropriate version:

```
VERSION=vX.X.X # Set the version you intent to release
env GOOS=linux GOARCH=amd64 go build -o terraform-provider-hirefire_$VERSION_linux_amd64
hub release create $VERSION
```

Then upload the file(s) to the GitHub's release that was just created.
