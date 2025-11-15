package resourceconfiguration

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// jsonSemanticEqualsPlanModifier is a plan modifier that suppresses diffs when JSON is semantically equal.
type jsonSemanticEqualsPlanModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m jsonSemanticEqualsPlanModifier) Description(_ context.Context) string {
	return "Suppresses diff when JSON values are semantically equal"
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m jsonSemanticEqualsPlanModifier) MarkdownDescription(_ context.Context) string {
	return "Suppresses diff when JSON values are semantically equal"
}

// PlanModifyString implements the plan modification logic.
func (m jsonSemanticEqualsPlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing if there is no state value.
	if req.StateValue.IsNull() {
		return
	}

	// Do nothing if there is no plan value.
	if req.PlanValue.IsUnknown() || req.PlanValue.IsNull() {
		return
	}

	// Do nothing if the values are already equal.
	if req.StateValue.Equal(req.PlanValue) {
		return
	}

	// Parse both JSON values.
	var stateJSON, planJSON interface{}
	if err := json.Unmarshal([]byte(req.StateValue.ValueString()), &stateJSON); err != nil {
		// If we can't parse the state value, let Terraform handle the diff.
		return
	}
	if err := json.Unmarshal([]byte(req.PlanValue.ValueString()), &planJSON); err != nil {
		// If we can't parse the plan value, let Terraform handle the diff.
		return
	}

	// Check if the plan JSON contains all the fields from the state JSON with the same values.
	// This allows the API to add additional fields without causing a diff.
	if jsonContains(stateJSON, planJSON) {
		// Values are semantically equal, use the state value to suppress the diff.
		resp.PlanValue = req.StateValue
	}
}

// jsonContains checks if container contains all fields from subset with matching values.
// This is a simplified check that works for our use case.
func jsonContains(container, subset interface{}) bool {
	switch subsetVal := subset.(type) {
	case map[string]interface{}:
		containerMap, ok := container.(map[string]interface{})
		if !ok {
			return false
		}
		// Check that all fields in subset exist in container with matching values.
		for key, subVal := range subsetVal {
			contVal, exists := containerMap[key]
			if !exists {
				// Subset has a field that container doesn't have.
				return false
			}
			if !jsonContains(contVal, subVal) {
				return false
			}
		}
		return true
	case []interface{}:
		containerSlice, ok := container.([]interface{})
		if !ok || len(containerSlice) != len(subsetVal) {
			return false
		}
		// Arrays must match exactly (same length and same values in same order).
		for i, subVal := range subsetVal {
			if !jsonContains(containerSlice[i], subVal) {
				return false
			}
		}
		return true
	default:
		// For primitive values, use direct comparison.
		return container == subset
	}
}

// JSONSemanticEquals returns a plan modifier that suppresses diffs when JSON is semantically equal.
func JSONSemanticEquals() planmodifier.String {
	return jsonSemanticEqualsPlanModifier{}
}
