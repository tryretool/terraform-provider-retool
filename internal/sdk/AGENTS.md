# Agent Guide: Updating the Retool API SDK

This guide documents the process for updating the Terraform provider's Go SDK to a new version of the Retool API. It's designed for both AI agents and human developers.

## Overview

The SDK is generated from an OpenAPI spec using [OpenAPI Generator](https://openapi-generator.tech/). The raw spec from Retool's API needs a handful of transformations before it generates working Go code; those are applied by **`internal/sdk/transform_spec.py`**. This guide walks through the entire update process.

## Prerequisites

- OpenAPI Generator CLI **7.6.0** (the version pinned in `generate.go`). Install per the main README, or download the jar from Maven Central and invoke it with `java -jar` (needs a JDK).
- Go development environment (see `go.mod` for the required version).
- Terraform CLI on `PATH` (for `go generate` docs + acceptance tests). The test framework's auto-install is currently broken (expired HashiCorp signing key), so install Terraform yourself and, for acceptance tests, set `TF_ACC_TERRAFORM_PATH=$(which terraform)`.
- Access to a Retool instance for testing.
- The new OpenAPI spec file (e.g., from https://api.retool.com/api/v2/spec).

> The previous spec is tracked in git (`internal/sdk/openAPISpec.json`), so git history is your backup — no need to keep `.backup` files.

## Update Process

### Step 1: Obtain and Prepare the New Spec

1. Download the new spec from Retool's API endpoint or get it from the Retool team
2. Place it in `internal/sdk/` with a descriptive name (e.g., `2.13.spec.json`)
3. Format it for readability:
   ```bash
   # In VSCode: "Format Document" command
   # Or use jq: cat 2.13.spec.json | jq '.' > 2.13.spec.formatted.json
   ```

### Step 2: Apply Required Transformations

All required spec transformations are consolidated in **`internal/sdk/transform_spec.py`**. It reads a raw spec and writes the transformed `openAPISpec.json` that `go generate` consumes:

```bash
cd internal/sdk
python3 transform_spec.py <raw_new_spec.json> openAPISpec.json
```

The transforms it applies (each written semantically so it survives upstream reordering):

1. **Single tag per operation.** Multiple tags make the generator emit duplicate types and break compilation. Keeps the tag matching the top-level path segment (e.g. `/users/{userId}/user_attributes` → `Users`). As of 4.0 the upstream spec already does this, so it's usually a no-op.
2. **Permissions `oneOf` → `anyOf`.** `/permissions/listObjects`, `/permissions/grant`, `/permissions/revoke` responses return structurally identical variants; `oneOf` causes "data matches more than one schema" unmarshal errors, `anyOf` doesn't.
3. **Free-form resource `options` request bodies.** The provider treats `options` as an opaque JSON blob, but the spec models it as a large `anyOf` union. The generated client ambiguously matches "thin" option objects (e.g. bearer-token auth) to the wrong member and drops fields like `base_url`. The transform rewrites the `options` *request* schema for `/resources` and `/resource_configurations` to `{type: object}` (`Options` becomes `map[string]interface{}`). Responses are left typed.
4. **Relax response-only required fields** (`folder_id`, `seat_type`, `default_value`). These are marked required on responses but omitted by not-yet-migrated instances, causing unmarshal errors. Removed from every `required` array (none are required in request bodies, so this is safe).

> **Bitbucket note:** older versions had to *remove* `type` from the Bitbucket config (the API rejected it). 4.0 reversed this — `type` is now **required** (`AppPassword`/`Token`), so there is no Bitbucket transform anymore and the provider sends `type` (see Step 5). If you see Bitbucket 400s, check which way the current API wants it.

**Finding NEW transformations for a version bump:** diff the new raw spec against the current `openAPISpec.json` (both are the same generator output style if you transform first). The most reliable approach is a small Python script that walks both specs and reports: paths added/removed, methods added/removed, and `required`/`properties`/`enum` diffs on the endpoints the provider actually calls. Pay special attention to:
- new `oneOf`/`anyOf` unions (may need fix #2-style handling),
- new fields marked `required` on responses (may need fix #4),
- constructor signature changes (required fields added/removed),
- response shape changes (e.g. an `anyOf` wrapper becoming a flat object).

If you discover a new transformation, add it to `transform_spec.py` (and note it in the summary it prints).

### Step 3: Generate the SDK

1. Produce the transformed spec (Step 2) into `internal/sdk/openAPISpec.json`.

2. Generate the code:
   ```bash
   cd internal/sdk
   go generate
   ```
   This runs: `openapi-generator-cli generate -i openAPISpec.json -g go -o ./api -c generatorConfig.yaml --minimal-update && ./fix_generated_code.sh`. If `openapi-generator-cli` isn't on `PATH`, run the generator jar directly with `java -jar <generator-7.6.0.jar> generate ...` followed by `./fix_generated_code.sh`.

3. **Remove stale generated files.** `--minimal-update` does not delete files, so renamed types (e.g. when a tag is pluralized like `Access Request` → `Access Requests`, or when a shared type splits in two) leave orphaned `.go`/`.md` files that cause "redeclared" compile errors. Delete anything in `api/` that is no longer in the generator manifest, preserving the two manually-maintained permission alias files:
   ```bash
   cd api
   python3 - <<'PY'
   import os
   manifest = {os.path.normpath(l.strip()) for l in open('.openapi-generator/FILES') if l.strip()}
   keep = {'model__permissions_grant_post_200_response_data_inner.go',
           'model__permissions_revoke_post_200_response_data_inner.go'}
   for root, _, files in os.walk('.'):
       for fn in files:
           rel = os.path.normpath(os.path.join(root, fn)).removeprefix('./')
           if rel.startswith('.openapi-generator') or rel in ('.openapi-generator-ignore', '.gitignore'):
               continue
           if rel.endswith(('.go', '.md', '.yaml', '.sh', '.yml')) and rel not in manifest and fn not in keep:
               os.remove(rel); print('removed', rel)
   PY
   ```

4. **Check for syntax errors immediately**:
   ```bash
   go build ./api
   ```

### Step 4: Fix Generator Bugs

The OpenAPI Generator has known bugs that require post-generation fixes. These are automated in `fix_generated_code.sh`, but new issues may arise.

#### 4.1 Known Issues (Auto-Fixed)

1. **Pointer-to-pointer marshaling**: `json.Marshal(&src.Field)` → `json.Marshal(src.Field)`
2. **Value receiver for MarshalJSON**: Pointer receivers → value receivers
3. **Invalid field names**: the generator emits `map[string]interface{} *map[string]interface{}` as a struct field in some `anyOf` option wrappers. `fix_generated_code.sh` now rewrites this (and its references) to `AdditionalProperties` across **every** affected file (it greps for the pattern rather than hard-coding filenames), so newly-split option models are handled automatically.

#### 4.2 New Invalid Field Names (Manual Check Required)

If `go build ./api` fails with syntax errors like:
- `unexpected keyword map, expected field name`
- `unexpected [, expected field name`  
- `unexpected keyword interface`

**Diagnosis**: The generator created field names using Go keywords or types.

**Common culprits**:
```go
// BAD - Generator output
type MyStruct struct {
    []string *[]string              // Invalid!
    string *string                  // Invalid!
    map[string]interface{} *map[string]interface{}  // Invalid!
}

// GOOD - Fixed
type MyStruct struct {
    ArrayOfString *[]string
    String *string
    AdditionalProperties *map[string]interface{}
}
```

**Fix process**:
1. Identify the file (error message shows filename)
2. Find all occurrences of the invalid field name
3. Replace with a valid identifier (ArrayOfString, String, AdditionalProperties, etc.)
4. Update in: struct definition, UnmarshalJSON, MarshalJSON methods

**Add to fix_generated_code.sh**:
```bash
# Fix model__problematic_file.go
if [ -f ./api/model__problematic_file.go ]; then
    sed -i '' 's/\[\]string \*\[\]string/ArrayOfString *[]string/g' ./api/model__problematic_file.go
    sed -i '' 's/dst\.\[\]string/dst.ArrayOfString/g' ./api/model__problematic_file.go
    # ... add all replacements
fi
```

### Step 5: Update Provider Code

The provider code in `internal/provider/` uses the generated SDK. API changes require updates here.

#### 5.1 Check for Compilation Errors

```bash
cd ../..  # Back to repo root
make test-unit 2>&1 | grep "undefined\|not enough arguments\|cannot use"
```

#### 5.2 Common Breaking Changes

**anyOf/oneOf wrapper types**: response data shapes can change between discriminated unions and flat objects.

```go
// 2.12.0: GET /users/{userId} data was an anyOf wrapper
var userData *api.UsersUserIdGet200ResponseDataAnyOf
if user.Data.UsersUserIdGet200ResponseDataAnyOf != nil {
    userData = user.Data.UsersUserIdGet200ResponseDataAnyOf
}

// 4.0: it's a flat object again
userData := &user.Data
```

**Constructor changes**: Check if required parameters changed.

```bash
# Find all calls to constructors
grep -r "api\.New" internal/provider/
```

The Bitbucket constructor has flipped between versions as the spec's `type`
requirement changed:
```go
// 2.12.0 (type removed from spec)
api.NewBitbucketConfigAnyOf(username, password)

// 4.0 (type required again)
api.NewBitbucketConfigAnyOf("AppPassword", username, password)
```
`NewAWSCodeCommitConfig` similarly changed in 4.0 to `(url, region)` only, with
the credential fields becoming optional setters.

**New required fields**: Check API changes in resource creation. Run acceptance
tests and look for `no value given for required property X` (an unmarshal error
on a response field the API omits) — relax it via `transform_spec.py`'s
`OPTIONAL_RESPONSE_FIELDS` list (this is how `folder_id`, `seat_type`, and
`default_value` are handled). `retoolresource` still sends `folder_id: null` in
its create request for backwards compatibility.

#### 5.3 Files to Check

Priority order:
1. `internal/provider/sourcecontrol/` - Often affected by config changes
2. `internal/provider/user/` - User structure changes frequently
3. `internal/provider/retoolresource/` - Resource API changes
4. Test files: `*_test.go` - Update constructors in tests

### Step 6: Test the SDK

#### 6.1 Unit Tests

```bash
make test-unit
```

**CRITICAL**: Must pass 100% before proceeding. No compilation errors allowed. This is non-negotiable.

#### 6.2 SDK Smoke Test

Test basic SDK functionality against a live Retool instance:

```bash
cd internal/sdk/client
RETOOL_HOST=your-instance.retool.dev \
RETOOL_SCHEME=https \
RETOOL_ACCESS_TOKEN=your_token \
go run main.go
```

Expected: Should successfully fetch and display folders.

#### 6.3 Acceptance Tests

**IMPORTANT**: Acceptance tests use a recording/replay system for CI. Always use recordings.

Because the test framework can no longer auto-install Terraform (expired signing key), set `TF_ACC_TERRAFORM_PATH` to a local Terraform binary for any live run:

```bash
# Record a SINGLE test (lightest weight - only its cassette is written, others untouched):
rm -f test/data/recordings/TestAccFoo.yaml
RETOOL_HOST=test-instance.retool.dev RETOOL_SCHEME=https RETOOL_ACCESS_TOKEN=test_token \
  TF_ACC_TERRAFORM_PATH="$(which terraform)" \
  FILTER='^TestAccFoo$' make test-acc-record

# Re-record the WHOLE suite (one test at a time, resumable, rate-limit friendly):
make test-acc-sweep
rm -rf test/data/recordings
RETOOL_HOST=test-instance.retool.dev RETOOL_SCHEME=https RETOOL_ACCESS_TOKEN=test_token \
  TF_ACC_TERRAFORM_PATH="$(which terraform)" \
  make test-acc-record-each      # tune with SLEEP_SECONDS / scope with FILTER

# Replay (CI + regular dev; offline, no live API):
make test-acc-replay
FILTER=TestAccFoo make test-acc-replay
```

**CRITICAL REQUIREMENTS**:
- ✅ **All replay tests must pass** - this is non-negotiable (13/13 test suites)
- ✅ Unit tests must pass 100%
- ✅ No "minor test data mismatches" allowed - all tests must be green

**Recording Tips**:
1. **Users can't be deleted** in Retool - increment version numbers in test emails (v3 → v4)
2. **Record after cleanup**: Run `make test-acc-sweep` before recording
3. **One test at a time**: Use `FILTER=TestName` to record specific tests
4. **Check recordings**: Ensure `.yaml` files are created in `test/data/recordings/`
5. **Re-recording the whole suite**: Don't run a single `make test-acc-record` over everything - the instance's rolling rate limit gets exhausted partway through and the rest fail with 429s (even though the target runs serially). Use `make test-acc-record-each` (wraps `scripts/record_acc_tests.sh`), which records each test in its own process with a pause between them. It is resumable - re-run it after a 429 to finish the rest. Start from `rm -rf test/data/recordings` since cassettes store a redacted host and can't be re-recorded incrementally. Tune with `SLEEP_SECONDS` and scope with `FILTER`.

**Expected results for LIVE API** (not CI):
- ⚠️ **429 (Rate Limiting)**: Normal - use recordings for CI
- ⚠️ **409 (Conflict)**: Resources exist - use recordings or increment versions
- ❌ **422/400/500**: Investigate and fix - these are real errors

**Unexpected failures**: Must be fixed
- **422 (Unprocessable)**: Check request validation, spec mismatch
- **400 (Bad Request)**: Check required fields, request structure
- **404 (Not Found)**: API endpoint changes
- **500 (Server Error)**: Possible spec mismatch

### Step 7: Document Changes

Update `internal/sdk/README.md` with a new section:

```markdown
## Notes from updating to X.Y.Z

Updated from A.B.C to X.Y.Z. Key changes:

### API Changes
- List breaking changes
- New required fields
- Structure changes (anyOf wrappers, etc.)

### Generator Fixes  
- New automatic fixes added
- Files affected

### Testing Results
- Unit tests: Status
- Acceptance tests: Status and known issues

### Breaking Changes
If upgrading from A.B.C, note:
1. ...
```

## Troubleshooting

### Issue: "data matches more than one schema in oneOf"

**Cause**: Unmarshaling error in discriminated unions with identical structures.

**Solution**: Change `oneOf` to `anyOf` in the spec (Step 2, transform #2 in `transform_spec.py`).

**Verification**: 
```bash
grep -n "oneOf" internal/sdk/openAPISpec.json | grep -A5 -B5 permissions
```

### Issue: Bitbucket config returns 400 errors

**Cause**: mismatch on the Bitbucket `type` field. This has flipped between
versions: pre-4.0 the API rejected `type` (it had to be omitted); 4.0 requires
it (`AppPassword`/`Token`).

**Solution**: make the provider/spec match what the current API wants. For 4.0,
`type` is required, so the provider passes `"AppPassword"` to
`NewBitbucketConfigAnyOf` and there is no Bitbucket spec transform. Verify with
`TestAccSourceControlTest`.

### Issue: `redeclared in this block` after generating

**Cause**: `--minimal-update` left stale files when a type was renamed/split
(common when tags get pluralized).

**Solution**: run the stale-file cleanup from Step 3.3 (delete files not in
`.openapi-generator/FILES`).

### Issue: a request silently drops fields (e.g. `base_url`) for some auth types

**Cause**: the strongly-typed `anyOf` for resource `options` ambiguously matches
"thin" objects to the wrong union member.

**Solution**: ensure the `options` request schema is free-form (Step 2, transform
#3). Confirm `ResourcesPostRequest.Options` is `map[string]interface{}`.

### Issue: "no value given for required property X"

**Cause**: New API version added required field not present in provider code.

**Solution**: 
1. Check the spec to confirm it's truly required
2. Update provider code to provide the field
3. If field shouldn't be required, fix the spec

**Finding the issue**:
```bash
# Search for the property in the spec
grep -n "required_field_name" internal/sdk/openAPISpec.json
```

### Issue: Go compilation fails with "unexpected keyword"

**Cause**: Generator created invalid Go identifier using keyword/type name.

**Solution**: Apply field name fixes (Step 4.2) and add to `fix_generated_code.sh`.

### Issue: Rate limiting in tests (429 errors)

**Not an SDK issue** - This is expected when running many tests rapidly.

**Solutions**:
- Use test recording/replay (see CONTRIBUTING.md)
- Run tests against a dedicated test instance
- Disable rate limiting on test instance (if possible)
- Add delays between tests

### Issue: Resource conflicts in tests (409 errors)

**Cause**: Previous test runs left resources in the system.

**Solution**:
```bash
make test-acc-sweep
```

## Quick Reference: Common Commands

```bash
# Transform a raw spec into the one the generator consumes
cd internal/sdk && python3 transform_spec.py <raw_spec.json> openAPISpec.json

# Generate SDK (then delete stale files per Step 3.3)
go generate

# Test compilation
go build ./api

# Run unit tests
cd ../.. && make test-unit

# Replay acceptance tests (offline)
make test-acc-replay

# Clean up test resources on the instance
RETOOL_HOST=host RETOOL_SCHEME=https RETOOL_ACCESS_TOKEN=token make test-acc-sweep
```

## Transformation Script

The transformations live in **`internal/sdk/transform_spec.py`** — extend that
file rather than writing a new script. Add a new `fix_*` function, call it from
`main()`, and include it in the printed summary so the counts are visible when
you regenerate. The legacy `fix_folder_id_response.py` is superseded by
`transform_spec.py`'s `fix_optional_response_fields` and can be ignored.

## Success Criteria

Before considering the update complete:

- [ ] **Unit tests pass 100%** ← REQUIRED, no exceptions
- [ ] **SDK compiles without errors** ← REQUIRED
- [ ] Stale generated files removed (`go build ./api` is clean)
- [ ] Smoke test successfully connects to API
- [ ] **ALL acceptance tests pass in replay mode** ← REQUIRED for CI
- [ ] Provider code updated for breaking changes
- [ ] `transform_spec.py` updated for new spec transformations
- [ ] `fix_generated_code.sh` updated for new generator bugs
- [ ] README.md updated with upgrade notes
- [ ] Test recordings updated for changed tests
- [ ] Provider docs regenerated (`make docs`)
- [ ] All changes committed with clear messages

**Critical**: Unit tests and replay tests must have 100% pass rate. Live API tests may fail due to rate limiting/conflicts - this is expected and why we use recordings.

## Historical Updates

- **3.334 → 4.0.0** (Jun 2026):
  - Consolidated all spec transformations into `transform_spec.py` (single-tag,
    permissions `oneOf`→`anyOf`, free-form resource `options` requests, and
    relaxing `folder_id`/`seat_type`/`default_value`).
  - Bitbucket `type` is required again → dropped the old Bitbucket transform;
    `NewBitbucketConfigAnyOf` takes `("AppPassword", username, appPassword)`.
  - `GET /users/{userId}` data went from an `anyOf` wrapper back to a flat object.
  - `NewAWSCodeCommitConfig` reduced to `(url, region)`; credential fields optional.
  - Generalized the invalid-field-name fix in `fix_generated_code.sh` to all files.
  - Fixed bearer-token auth on REST API resources (#61) via free-form `options`.
  - New optional provider attributes: user `seat_type`, configuration variable
    `default_value`, SSO `oidc end_session_url`, source control CodeCommit
    `assume_role` / `auth_with_default_credential_provider_chain`.
  - Added `make test-acc-record-each` for rate-limit-friendly re-recording.
  - ✅ Unit + full replay suite green.
- **2.9.0 → 2.12.0** (Feb 2026): 
  - Fixed Bitbucket config (removed `type` field entirely from spec)
  - **Fixed resource `folder_id` issue**: Created `fix_folder_id_response.py` to remove from required arrays in all response schemas
  - Added automatic fixes for invalid Go field names (anyOf with keywords/types)
  - Fixed pointer-to-pointer marshaling bug in anyOf types
  - Updated user data handling (anyOf wrapper structure)
  - Updated test emails (v3 → v4) - users can't be deleted in Retool
  - Added `folder_id: null` to resource creation requests for backwards compatibility
  - ✅ **All unit tests passing (100%)**
  - ✅ **All acceptance tests passing in replay mode (13/13)**
- **2.4.0 → 2.9.0** (Previous): Initial transformation automation with tags fix, permissions anyOf

## Tips for AI Agents

1. **Always read README.md first** - Contains critical context and known issues
2. **Compare specs before transforming** - Understand what changed; focus on endpoints the provider actually calls
3. **Test incrementally** - Don't wait until the end to compile
4. **Extend `transform_spec.py`** - Don't hand-edit `openAPISpec.json`; git history is your backup
5. **Delete stale generated files** - `--minimal-update` leaves orphans that fail to compile
6. **Document as you go** - Update README.md with findings
7. **Check generator bugs** - Every major version may introduce new issues
8. **Run the full replay suite** - Acceptance tests catch subtle issues
9. **Read error messages carefully** - They often point to exact fixes needed

## Related Documentation

- [internal/sdk/README.md](./README.md) - SDK overview and known issues
- [CONTRIBUTING.md](../../CONTRIBUTING.md) - Testing and development workflow
- [OpenAPI Generator Docs](https://openapi-generator.tech/docs/generators/go/)
