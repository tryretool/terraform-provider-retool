package utils

import (
	"math"
	"net/http"
	"strconv"

	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// providerData is the data structure that is passed to the DataSource and Resource
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
		return "" // or some default value
	}
	rounded := math.Round(float64(*f))
	return strconv.Itoa(int(rounded))
}
