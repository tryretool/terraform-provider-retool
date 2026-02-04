# Agent Guide: Updating the Retool API SDK

This guide documents the process for updating the Terraform provider's Go SDK to a new version of the Retool API. It's designed for both AI agents and human developers.

## Overview

The SDK is generated from an OpenAPI spec using [OpenAPI Generator](https://openapi-generator.tech/). However, the raw spec from Retool's API requires manual modifications to work correctly with the Go code generator. This guide walks through the entire update process.

## Prerequisites

- OpenAPI Generator CLI installed (see main README.md)
- Go development environment
- Access to a Retool instance for testing
- The new OpenAPI spec file (e.g., from https://api.retool.com/api/v2/spec)

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

The raw spec needs several modifications before it can generate working Go code. Create a Python script or manually apply these changes:

#### 2.1 Fix Multiple Tags (REQUIRED)

**Problem**: Endpoints with multiple tags cause the generator to create duplicate types, resulting in compilation errors.

**Solution**: Ensure every endpoint has only ONE tag. Keep the tag that matches the top-level path element.

Example:
```json
// BAD: Multiple tags
"tags": ["Users", "User Attributes"]

// GOOD: Single tag matching top-level path
"tags": ["Users"]
```

For `/users/{userId}/user_attributes/{attributeName}`, keep `["Users"]`.

**Automation**:
```python
def fix_tags(spec):
    for path, path_item in spec['paths'].items():
        path_parts = [p for p in path.split('/') if p and not p.startswith('{')]
        top_level_tag = path_parts[0].capitalize() if path_parts else None
        
        for method in ['get', 'post', 'put', 'patch', 'delete']:
            if method in path_item and 'tags' in path_item[method]:
                tags = path_item[method]['tags']
                if len(tags) > 1:
                    if top_level_tag and top_level_tag in tags:
                        path_item[method]['tags'] = [top_level_tag]
                    else:
                        path_item[method]['tags'] = [tags[0]]
```

#### 2.2 Fix Permissions Endpoints: oneOf → anyOf (REQUIRED)

**Problem**: Permissions endpoints return objects with identical structures but different enum values. Using `oneOf` causes unmarshaling errors ("data matches more than one schema").

**Solution**: Change `oneOf` to `anyOf` in these endpoints:
- `/permissions/listObjects` (POST response)
- `/permissions/grant` (POST response)  
- `/permissions/revoke` (POST response)

**Location in spec**: Look for response schemas where `oneOf` contains objects with identical fields but different `type` enum values (folder, app, resource, resource_configuration).

**Example**:
```json
// BEFORE
"items": {
  "oneOf": [
    { "type": "folder", "properties": {...} },
    { "type": "app", "properties": {...} }
  ]
}

// AFTER
"items": {
  "anyOf": [
    { "type": "folder", "properties": {...} },
    { "type": "app", "properties": {...} }
  ]
}
```

#### 2.3 Fix Bitbucket Configuration (KNOWN ISSUE)

**Problem**: The spec incorrectly marks `type` as required for Bitbucket configs. The Retool API rejects requests with this field and returns 400 errors.

**Solution**: Remove `"type"` from the `required` arrays in Bitbucket configuration schemas.

**Locations to check**:
- Look for schemas with `app_password` or `username` + Bitbucket-specific fields
- Both `AppPassword` and `Token` variants
- GET/POST/PUT endpoints for source control

**Example**:
```json
// BEFORE
"required": ["type", "username", "app_password"]

// AFTER  
"required": ["username", "app_password"]
```

#### 2.4 Fix Resource `folder_id` in Response Schemas (CRITICAL)

**Problem**: The spec incorrectly marks `folder_id` as required in resource response schemas, but the API doesn't always return it (depends on Retool version). This causes SDK unmarshaling errors.

**Solution**: Remove `"folder_id"` from the `required` arrays in all resource and resource_configuration response schemas. This makes the field optional while keeping it in the schema for forwards compatibility.

**Affected endpoints**:
- `/resources` - POST, GET (single), GET (list), PATCH responses
- `/resource_configurations` - POST, GET (single), GET (list), PATCH responses
- Note: Some responses have `data.folder_id`, others have `data.resource.folder_id`

**Automation**: Use the `fix_folder_id_response.py` script:
```python
# Script handles all variations automatically with error handling
python3 fix_folder_id_response.py
```

**Why this matters**:
- **Older Retool versions**: Don't return `folder_id` → SDK must accept response
- **Newer Retool versions**: Return `folder_id` → SDK parses it normally
- Making it optional ensures compatibility across all versions

#### 2.5 Search for Other Breaking Changes

Compare the new spec with the current one:
```bash
# Look for new required fields
diff <(jq '.paths[].post.requestBody.content."application/json".schema.required' internal/sdk/openAPISpec.json) \
     <(jq '.paths[].post.requestBody.content."application/json".schema.required' internal/sdk/2.13.spec.json)

# Check for structural changes in common endpoints
jq '.paths."/users/{userId}"' internal/sdk/2.13.spec.json
jq '.paths."/resources"' internal/sdk/2.13.spec.json
```

### Step 3: Generate the SDK

1. Replace the current spec:
   ```bash
   cd internal/sdk
   cp openAPISpec.json openAPISpec.json.backup  # Always keep a backup!
   cp 2.13.spec.json openAPISpec.json  # Use your transformed spec
   ```

2. Generate the code:
   ```bash
   go generate
   ```
   
   This runs: `openapi-generator-cli generate -i openAPISpec.json -g go -o ./api -c generatorConfig.yaml --minimal-update && ./fix_generated_code.sh`

3. **Check for syntax errors immediately**:
   ```bash
   go build ./api
   ```

### Step 4: Fix Generator Bugs

The OpenAPI Generator has known bugs that require post-generation fixes. These are automated in `fix_generated_code.sh`, but new issues may arise.

#### 4.1 Known Issues (Auto-Fixed)

1. **Pointer-to-pointer marshaling**: `json.Marshal(&src.Field)` → `json.Marshal(src.Field)`
2. **Value receiver for MarshalJSON**: Pointer receivers → value receivers
3. **Invalid field names** (API 2.12.0+): Generator creates invalid Go identifiers using keywords

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

**anyOf/oneOf wrapper types**: Newer API versions may wrap response data in discriminated unions.

Example from v2.12.0:
```go
// OLD (v2.9.0)
user.Data.Id
user.Data.Email

// NEW (v2.12.0)
var userData *api.UsersUserIdGet200ResponseDataAnyOf
if user.Data.UsersUserIdGet200ResponseDataAnyOf != nil {
    userData = user.Data.UsersUserIdGet200ResponseDataAnyOf
}
userData.Id
userData.Email
```

**Constructor changes**: Check if required parameters changed.

```bash
# Find all calls to constructors
grep -r "api\.New" internal/provider/
```

Example from v2.12.0:
```go
// OLD (v2.9.0)
api.NewBitbucketConfigAnyOf(username, password)

// NEW (v2.12.0) - now requires type
api.NewBitbucketConfigAnyOf("AppPassword", username, password)
```

**New required fields**: Check API changes in resource creation.

```bash
# Search for error messages containing "required property"
# Run acceptance tests and check failures
```

**Important**: If you see errors about `folder_id` being required:
- This is likely a spec bug, not a real API requirement
- Check if `folder_id` is in response schemas' required arrays
- Use `fix_folder_id_response.py` to remove it
- Add `folder_id: null` to request bodies in provider code for backwards compatibility

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

**IMPORTANT**: Acceptance tests use a recording/replay system for CI. Always use recordings:

```bash
# Step 1: Record tests against live API (one-time or when API changes)
RETOOL_HOST=test-instance.retool.dev \
RETOOL_SCHEME=https \
RETOOL_ACCESS_TOKEN=test_token \
FILTER=TestAccFoo \
make test-acc-record

# Step 2: Replay tests using recordings (for CI and regular development)
FILTER=TestAccFoo make test-acc-replay

# Full replay suite
make test-acc-replay
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

**Solution**: Change `oneOf` to `anyOf` in the spec (Step 2.2).

**Verification**: 
```bash
grep -n "oneOf" internal/sdk/openAPISpec.json | grep -A5 -B5 permissions
```

### Issue: Bitbucket config returns 400 errors

**Cause**: API rejects the `type` field despite spec marking it required.

**Solution**: Remove `type` from required arrays in Bitbucket schemas (Step 2.3).

**Verification**: Check that Bitbucket config constructors work in tests.

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
# Generate SDK
cd internal/sdk && go generate

# Test compilation
go build ./api

# Run unit tests
cd ../.. && make test-unit

# Run acceptance tests with credentials
RETOOL_HOST=host RETOOL_SCHEME=https RETOOL_ACCESS_TOKEN=token make test-acc

# Clean up test resources
make test-acc-sweep

# Check git status (see what changed)
git status

# Create a backup before major changes
cp internal/sdk/openAPISpec.json internal/sdk/openAPISpec.json.backup
```

## Automation Script Template

For future updates, you can create a transformation script:

```python
#!/usr/bin/env python3
"""Transform Retool OpenAPI spec for Go code generation."""
import json
from pathlib import Path

def fix_tags(spec):
    """Ensure every endpoint has only one tag."""
    # Implementation from Step 2.1
    pass

def fix_permissions_oneof_to_anyof(spec):
    """Change oneOf to anyOf in permissions endpoints."""
    # Implementation from Step 2.2
    pass

def fix_bitbucket_configs(spec):
    """Remove type from required in Bitbucket configs."""
    # Implementation from Step 2.3
    pass

def main():
    input_file = Path('2.XX.spec.json')
    output_file = Path('openAPISpec.json')
    
    spec = json.load(input_file.open())
    
    fix_tags(spec)
    fix_permissions_oneof_to_anyof(spec)
    fix_bitbucket_configs(spec)
    
    json.dump(spec, output_file.open('w'), indent=2)
    print(f"✓ Transformed {input_file} → {output_file}")

if __name__ == '__main__':
    main()
```

## Success Criteria

Before considering the update complete:

- [ ] **Unit tests pass 100%** ← REQUIRED, no exceptions
- [ ] **SDK compiles without errors** ← REQUIRED
- [ ] Smoke test successfully connects to API
- [ ] **ALL acceptance tests pass in replay mode** ← REQUIRED for CI
- [ ] Provider code updated for breaking changes
- [ ] `fix_generated_code.sh` updated for new generator bugs
- [ ] README.md updated with upgrade notes
- [ ] Test recordings updated for changed tests
- [ ] All changes committed with clear messages

**Critical**: Unit tests and replay tests must have 100% pass rate. Live API tests may fail due to rate limiting/conflicts - this is expected and why we use recordings.

## Historical Updates

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
2. **Compare specs before transforming** - Understand what changed
3. **Test incrementally** - Don't wait until the end to compile
4. **Backup before replacing** - Keep `openAPISpec.json.backup`
5. **Document as you go** - Update README.md with findings
6. **Check generator bugs** - Every major version may introduce new issues
7. **Run full test suite** - Acceptance tests catch subtle issues
8. **Read error messages carefully** - They often point to exact fixes needed

## Related Documentation

- [internal/sdk/README.md](./README.md) - SDK overview and known issues
- [CONTRIBUTING.md](../../CONTRIBUTING.md) - Testing and development workflow
- [OpenAPI Generator Docs](https://openapi-generator.tech/docs/generators/go/)
