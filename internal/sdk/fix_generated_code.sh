#!/bin/bash
# Post-generation script to fix bugs in OpenAPI-generated Go code

echo "Fixing OpenAPI generator bugs..."

# 1. Fix pointer-to-pointer marshaling bug: json.Marshal(&src.Field) -> json.Marshal(src.Field)
find ./api -name "*.go" -type f -exec sed -i '' 's/json\.Marshal(\&src\.\([A-Za-z0-9_]*\))/json.Marshal(src.\1)/g' {} \;

# 2. Fix value receiver issue: Change pointer receivers to value receivers for MarshalJSON in anyOf types
#    The OpenAPI generator creates pointer receiver MarshalJSON methods, but when the struct
#    is used as a value type in another struct, Go won't call the pointer receiver method.
find ./api -name "*.go" -type f -exec sed -i '' 's/func (src \*\([A-Za-z0-9_]*\)) MarshalJSON()/func (src \1) MarshalJSON()/g' {} \;

# 3. Fix invalid field names in anyOf structs (API 2.12.0+)
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

# 4. Fix: model__resources_post_request_options.go
if [ -f ./api/model__resources_post_request_options.go ]; then
    sed -i '' 's/map\[string\]interface{} \*map\[string\]interface{}/AdditionalProperties *map[string]interface{}/g' ./api/model__resources_post_request_options.go
    sed -i '' 's/dst\.map\[string\]interface{}/dst.AdditionalProperties/g' ./api/model__resources_post_request_options.go
    sed -i '' 's/jsonmap\[string\]interface{}/jsonAdditionalProperties/g' ./api/model__resources_post_request_options.go
    sed -i '' 's/src\.map\[string\]interface{}/src.AdditionalProperties/g' ./api/model__resources_post_request_options.go
fi

echo "Fixed marshaling bugs in generated code"

# Note: Permissions oneOf unmarshaling is not fixed automatically.
# See README.md for details on the known limitations.

