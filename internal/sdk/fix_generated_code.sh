#!/bin/bash
# Post-generation script to fix bugs in OpenAPI-generated Go code

echo "Fixing OpenAPI generator bugs..."

# 1. Fix pointer-to-pointer marshaling bug: json.Marshal(&src.Field) -> json.Marshal(src.Field)
find ./api -name "*.go" -type f -exec sed -i '' 's/json\.Marshal(\&src\.\([A-Za-z0-9_]*\))/json.Marshal(src.\1)/g' {} \;

# 2. Fix value receiver issue: Change pointer receivers to value receivers for MarshalJSON in anyOf types
#    The OpenAPI generator creates pointer receiver MarshalJSON methods, but when the struct
#    is used as a value type in another struct, Go won't call the pointer receiver method.
find ./api -name "*.go" -type f -exec sed -i '' 's/func (src \*\([A-Za-z0-9_]*\)) MarshalJSON()/func (src \1) MarshalJSON()/g' {} \;

echo "Fixed marshaling bugs in generated code"

# Note: Permissions oneOf unmarshaling is not fixed automatically.
# See README.md for details on the known limitations.

