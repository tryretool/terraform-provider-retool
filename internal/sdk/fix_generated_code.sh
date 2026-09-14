#!/bin/bash
# Post-generation script to fix bugs in OpenAPI-generated Go code
#
# The field renames run before the marshaling fixes: a field the generator named
# `[]string` or `map[string]interface{}` won't match the marshaling patterns
# until it has a valid Go identifier.

echo "Fixing OpenAPI generator bugs..."

# 1. Fix invalid field names in anyOf structs (API 2.12.0+)
#    The generator sometimes creates invalid Go field names using keywords or type names
#    Fix: model__user_tasks_get_assigned_to_users_parameter.go
if [ -f ./api/model__user_tasks_get_assigned_to_users_parameter.go ]; then
    sed -i '' 's/\[\]string \*\[\]string/ArrayOfString *[]string/g' ./api/model__user_tasks_get_assigned_to_users_parameter.go
    sed -i '' 's/string \*string/String *string/g' ./api/model__user_tasks_get_assigned_to_users_parameter.go
    sed -i '' 's/dst\.\[\]string/dst.ArrayOfString/g' ./api/model__user_tasks_get_assigned_to_users_parameter.go
    sed -i '' 's/json\[\]string/jsonArrayOfString/g' ./api/model__user_tasks_get_assigned_to_users_parameter.go
    sed -i '' 's/dst\.string/dst.String/g' ./api/model__user_tasks_get_assigned_to_users_parameter.go
    sed -i '' 's/jsonstring/jsonString/g' ./api/model__user_tasks_get_assigned_to_users_parameter.go
    sed -i '' 's/src\.\[\]string/src.ArrayOfString/g' ./api/model__user_tasks_get_assigned_to_users_parameter.go
    sed -i '' 's/src\.string/src.String/g' ./api/model__user_tasks_get_assigned_to_users_parameter.go
fi

# 2. Fix invalid `map[string]interface{}` field names in anyOf option models.
#    The generator emits `map[string]interface{} *map[string]interface{}` as a
#    struct field (e.g. in *_options.go anyOf wrappers), which is not valid Go.
#    Rename it (and its references) to AdditionalProperties across every file
#    that contains the pattern.
for f in $(grep -rl 'map\[string\]interface{} \*map\[string\]interface{}' ./api --include='*.go'); do
    sed -i '' 's/map\[string\]interface{} \*map\[string\]interface{}/AdditionalProperties *map[string]interface{}/g' "$f"
    sed -i '' 's/dst\.map\[string\]interface{}/dst.AdditionalProperties/g' "$f"
    sed -i '' 's/jsonmap\[string\]interface{}/jsonAdditionalProperties/g' "$f"
    sed -i '' 's/src\.map\[string\]interface{}/src.AdditionalProperties/g' "$f"
done

# 3. Fix pointer-to-pointer marshaling bug: json.Marshal(&src.Field) -> json.Marshal(src.Field)
find ./api -name "*.go" -type f -exec sed -i '' 's/json\.Marshal(\&src\.\([A-Za-z0-9_]*\))/json.Marshal(src.\1)/g' {} \;

# 4. Fix value receiver issue: Change pointer receivers to value receivers for MarshalJSON in anyOf types
#    The OpenAPI generator creates pointer receiver MarshalJSON methods, but when the struct
#    is used as a value type in another struct, Go won't call the pointer receiver method.
find ./api -name "*.go" -type f -exec sed -i '' 's/func (src \*\([A-Za-z0-9_]*\)) MarshalJSON()/func (src \1) MarshalJSON()/g' {} \;

# 5. Drop decoder.DisallowUnknownFields() from model UnmarshalJSON methods.
#    Schemas the API declares as strict get this call, which turns any response
#    field we don't know about into a hard error. The provider has to read
#    instances newer than the spec it was built from, so an unrecognized field
#    must be ignored, not fatal. The required-property check above it stays, so
#    genuinely malformed responses are still rejected.
#
#    oneOf/anyOf variants are exempt — their strictness is what tells the variants
#    apart, and a variant with a custom UnmarshalJSON bypasses whatever settings
#    the union's dispatcher passed in, so it cannot come from anywhere else.
#    Both dispatcher styles depend on it:
#      - oneOf counts how many variants decode cleanly. Strip it and several match,
#        giving "data matches more than one schema in oneOf".
#      - anyOf returns on the first variant that decodes. Strip it and a payload
#        matches an earlier, narrower variant, silently dropping the fields that
#        variant doesn't declare.
#    The variant list is read back out of the generated code so it stays correct
#    as the spec changes.
variant_types=$(grep -rhoE '(newStrictDecoder\(data\)\.Decode\(|json\.Unmarshal\(data, )&dst\.[A-Za-z0-9_]+\)' ./api --include='*.go' \
    | sed -E 's/.*&dst\.([A-Za-z0-9_]+)\)$/\1/' | sort -u)

for f in $(grep -rl 'decoder\.DisallowUnknownFields()' ./api --include='*.go'); do
    receiver=$(grep -oE 'func \(o \*[A-Za-z0-9_]+\) UnmarshalJSON' "$f" \
        | sed -E 's/func \(o \*([A-Za-z0-9_]+)\) UnmarshalJSON/\1/' | head -1)
    if grep -qxF "$receiver" <<< "$variant_types"; then
        continue
    fi
    sed -i '' '/^[[:space:]]*decoder\.DisallowUnknownFields()$/d' "$f"
done

echo "Fixed marshaling bugs in generated code"

# Note: Permissions oneOf unmarshaling is not fixed automatically.
# See README.md for details on the known limitations.
