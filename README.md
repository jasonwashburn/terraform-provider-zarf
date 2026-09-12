# Terraform Provider Zarf

Terraform provider for [Zarf](https://zarf.dev), built with the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework).

## Status

This provider is in VERY early development and, at least for now, is primarily intended to be a personal learning project.

## Development Setup

[Mise](https://mise.jdx.dev/) manages the development tools used by this project, including Go, OpenTofu, Terraform, golangci-lint, HK, and Zizmor.

Install the configured tools and build the provider:

```shell
mise install
mise run build
```

The build task creates the provider binary in `build/` and generates `.tofurc.dev` with a machine-specific development override. Both files are intentionally ignored by Git.

Mise installs the HK Git hook automatically after tool installation. To install it manually, run `hk install --mise`.

When using the local provider with OpenTofu, run OpenTofu through Mise or activate Mise in your shell so `TF_CLI_CONFIG_FILE` is set:

```shell
mise exec -- tofu providers schema -json
```

## Tasks

| Task | Purpose |
| --- | --- |
| `mise run build` | Build the provider and configure the local OpenTofu override |
| `mise run configure` | Regenerate `.tofurc.dev` |
| `mise run build-check` | Compile all Go packages |
| `mise run fmt` | Format Go source files |
| `mise run lint` | Run golangci-lint |
| `mise run test` | Run the test suite |
| `mise run generate` | Generate provider documentation |
| `mise run testacc` | Run acceptance tests |

Documentation generation uses Terraform internally through `tfplugindocs`; OpenTofu remains the local development CLI.

Zizmor audits GitHub Actions workflows locally and in CI.

Run all repository checks manually with:

```shell
hk check --all
```
