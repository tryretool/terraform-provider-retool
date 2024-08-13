// Package utils provides utility functions for the provider and resources.
package utils

import (
	"context"
	"math"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// providerData is the data structure that is passed to the DataSource and Resource.
type ProviderData struct {
	Client            *api.APIClient
	RootFolderIdCache *map[string]string
}

// If httpResponse is not nil, add the status code to the properties map.
func AddHttpStatusCode(props map[string]any, httpResponse *http.Response) map[string]any {
	if httpResponse != nil {
		props["httpResponse.status_code"] = httpResponse.StatusCode
	}
	return props
}

// Some of the ids are returned as numbers, and our SDK client represents them as float32, but we need them as strings.
func Float32PtrToIntString(f *float32) string {
	if f == nil {
		return "" // Or some default value.
	}
	rounded := math.Round(float64(*f))
	return strconv.Itoa(int(rounded))
}

func GetStringListValue(ctx context.Context, list types.List, diags *diag.Diagnostics) []string {
	listValue := make([]types.String, 0, len(list.Elements()))
	d := list.ElementsAs(ctx, &listValue, false)
	diags.Append(d...)
	result := make([]string, 0, len(list.Elements()))
	for _, s := range listValue {
		result = append(result, s.ValueString())
	}
	return result
}

// Returns true if the object value is null or unknown.
func IsEmptyObject(value types.Object) bool {
	return value.IsNull() || value.IsUnknown()
}

func IsEmptyMap(value types.Map) bool {
	return value.IsNull() || value.IsUnknown() || len(value.Elements()) == 0
}

func IsEmptyList(value types.List) bool {
	return value.IsNull() || value.IsUnknown() || len(value.Elements()) == 0
}
