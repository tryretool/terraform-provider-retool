# Development
Install Go, make sure `GOPATH` variable is set and `$GOPATH/bin` is in the `PATH`.

Follow the [instructions](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider#prepare-terraform-for-local-provider-install) to use local version of Retool provider.

## Updating Retool API client library
The repository contains code-generated API client library for Retool API - see more details in [internal/sdk/README.md](/internal/sdk/README.md)

# Testing
You can test the provider, resources and data sources against your local dev:
```
cd examples/provider-initialization
RETOOL_HOST=localhost:3000 RETOOL_SCHEME=http RETOOL_ACCESS_TOKEN=<the API token you created> terraform plan
```

To run unit tests, run the following command from the root folder:
```
make test-unit
```

You can run specific tests by setting `FILTER` env var:
```
FILTER=TestNewProvider make test-unit
``` 

To run acceptance tests, set `RETOOL_HOST`, `RETOOL_SCHEME` and `RETOOL_ACCESS_TOKEN` env vars, then run
```
make test-acc
```
Note that the acceptance tests will create, modify and delete resources in the Retool org.
You can run specific tests by setting `FILTER` env var as well.

If acceptance tests fail and leave some undeleted resources on the Retool instance, you can delete them using "sweepers":
```
make test-acc-sweep
```

If you want to see more detailed Terraform logs, add `TF_LOG=info` to your command.

## Recording and replaying HTTP requests for acceptance tests
It's hard to run acceptance tests on CI hitting a live Retool instance. Instead, we have a way to record and store HTTP responses during acceptance test run, then use canned responses to avoid hitting Retool instance.
Recorded responses are stored in `test/data/recordings/<test name>.yaml` files.

To record responses for an acceptance test TestAccFoo:
- Delete `test/data/recordings/TestAccFoo.yaml` if it already exists
- Set `RETOOL_HOST`, `RETOOL_SCHEME` and `RETOOL_ACCESS_TOKEN` env vars to appropriate values for the Retool instance you're going to use.
- Run `FILTER=TestAccFoo make test-acc-record` . This will run the acceptance test and store HTTP responses.

To run acceptance test using pre-recorded responses, do
```
FILTER=TestAccFoo make test-acc-replay
```

### Re-recording the whole suite (rate limits)
Recording the entire suite in a single `make test-acc-record` run almost always fails partway through with `429 Too Many Requests`. The Retool instance enforces a *rolling* API rate limit, and the suite sustains enough requests that the budget is exhausted after the first few packages — even though `test-acc-record` already runs serially (`-p 1 -parallel 1`).

To re-record everything reliably, record one test at a time with a pause between each so the rate limit can recover:
```
# 1. Clean up leftover resources from prior/aborted runs.
make test-acc-sweep
# 2. Start from a clean slate (cassettes store a redacted host, so partial/old
#    cassettes can't be re-recorded incrementally - you must remove them).
rm -rf test/data/recordings
# 3. Record each acceptance test in its own process, pausing in between.
make test-acc-record-each
```
`make test-acc-record-each` (a wrapper around `scripts/record_acc_tests.sh`) is **resumable**: it skips tests whose cassette already exists and removes the partial cassette of any test that fails, so if some still hit a `429` you can simply run it again to finish the rest. Tunable options:
- `FILTER` — only record tests matching a name regex, e.g. `make test-acc-record-each FILTER=TestAccPermissions`.
- `SLEEP_SECONDS` — seconds to wait between tests (default 30); raise it if your instance's limit is tight, e.g. `SLEEP_SECONDS=60 make test-acc-record-each`.

Two gotchas to keep in mind:
- **Users can't be deleted in Retool** (the provider's delete only deactivates them), so every successful re-record of the user tests consumes the test emails. Bump the `vN` suffix in `internal/provider/user/resource_test.go` before re-recording the user tests.
- The single-process `make test-acc-record` is still fine for recording an individual test (`FILTER=TestAccFoo make test-acc-record`); the per-test script is for re-recording many/all tests at once.

See `internal/acctest` for implementation details. The implementation was mostly copied from Auth0's Terraform provider: https://github.com/auth0/terraform-provider-auth0/tree/main/internal/acctest .
The most durable fix for the rate limiting is to increase or disable the API rate limit on your test instance. See https://registry.terraform.io/providers/tryretool/retool/latest/docs#rate-limiting for details.

# Documentation
Run `go generate` in the root of this repository to generate provider docs. See https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-documentation-generation for more details on how to add new examples.

# Publishing new version to Terraform Repository
Go to https://github.com/tryretool/terraform-provider-retool/releases, click "Draft a new release". Create new tag that follows `v<major version>.<minor version>.<patch version>` format. Most of the fixes should only increment the patch version. 
Minor version increments should be tied to Retool on-prem stable releases, since they are likely to rely on new API functionality not available in the previous versions. Major version increment should be reserved for major backward-incompatible changes.
Fill out release description and publish it. That'll kick off Github workflow defined [here](https://github.com/tryretool/terraform-provider-retool/blob/d8832910a0cb3cc719387b07caa788d573f6a24c/.github/workflows/release.yaml).
New version should show up on Terraform Repository in 10-15 minutes.
If signing keys need to be updated, sign in to `registry.terraform.io` using **Legacy auth via GitHub** (use the small link at the bottom of the auth page). Manage keys at https://registry.terraform.io/settings/gpg-keys. You must be an admin of the `tryretool` organization to update these keys. More information [here](https://developer.hashicorp.com/terraform/registry/providers/publishing#preparing-and-adding-a-signing-key)

# Update Retool CLI
Retool CLI has a `terraform` command that generates Terraform configuration from existing Retool org: https://github.com/tryretool/retool-cli/blob/master/src/commands/terraform.ts. 
If you added a new resource or updated an existing one, you should update Retool CLI as well, so that your changes are reflected in auto-generated configuration.

# Reviewing

Reviewing PRs to this repo can be a challenge, particularly when they impact the OpenAPI generated code for the SDK. 

As a general rule, code within `internal/provider/sdk` will be generated, and can largely be ignored during reviews. 

Code _outside_ of this should be given close attention. `internal/provider` code is what drives the terraform behavior, and should be reviewed for correctness and backward compatibility. Tests should accompany new features, and reviewers should make a habit of checking out PR branches and running them against realisitic example retool deployments.

Updates should always include documentation updates for new/changed behaviors.
