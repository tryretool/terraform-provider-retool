#!/usr/bin/env python3
"""Remove folder_id from required fields in resource response schemas."""
import json
from pathlib import Path


def safe_remove_folder_id(schema_path, description):
    """Safely try to remove folder_id from a schema's required array."""
    try:
        if "required" in schema_path and "folder_id" in schema_path["required"]:
            schema_path["required"].remove("folder_id")
            return description
    except (KeyError, TypeError):
        pass
    return None


def fix_folder_id_in_responses(spec):
    """Remove folder_id from required arrays in resource response schemas."""
    changes = []

    # Fix POST /resources response
    try:
        post_response_schema = spec["paths"]["/resources"]["post"]["responses"]["200"][
            "content"
        ]["application/json"]["schema"]["properties"]["data"]
        result = safe_remove_folder_id(post_response_schema, "POST /resources response")
        if result:
            changes.append(result)
    except KeyError:
        pass

    # Fix GET /resources/{resourceId} response
    try:
        get_response_schema = spec["paths"]["/resources/{resourceId}"]["get"][
            "responses"
        ]["200"]["content"]["application/json"]["schema"]["properties"]["data"]
        result = safe_remove_folder_id(
            get_response_schema, "GET /resources/{resourceId} response"
        )
        if result:
            changes.append(result)
    except KeyError:
        pass

    # Fix PATCH /resources/{resourceId} response
    try:
        patch_response_schema = spec["paths"]["/resources/{resourceId}"]["patch"][
            "responses"
        ]["200"]["content"]["application/json"]["schema"]["properties"]["data"]
        result = safe_remove_folder_id(
            patch_response_schema, "PATCH /resources/{resourceId} response"
        )
        if result:
            changes.append(result)
    except KeyError:
        pass

    # Fix GET /resources response (list)
    try:
        list_response_schema = spec["paths"]["/resources"]["get"]["responses"]["200"][
            "content"
        ]["application/json"]["schema"]["properties"]["data"]["items"]
        result = safe_remove_folder_id(
            list_response_schema, "GET /resources response (list)"
        )
        if result:
            changes.append(result)
    except KeyError:
        pass

    # Fix POST /resource_configurations response (nested resource object)
    try:
        resource_config_response = spec["paths"]["/resource_configurations"]["post"][
            "responses"
        ]["200"]["content"]["application/json"]["schema"]["properties"]["data"][
            "properties"
        ][
            "resource"
        ]
        result = safe_remove_folder_id(
            resource_config_response, "POST /resource_configurations response"
        )
        if result:
            changes.append(result)
    except KeyError:
        pass

    # Fix GET /resource_configurations/{configurationId} response (nested resource object)
    try:
        get_config_response = spec["paths"][
            "/resource_configurations/{configurationId}"
        ]["get"]["responses"]["200"]["content"]["application/json"]["schema"][
            "properties"
        ][
            "resource"
        ]
        result = safe_remove_folder_id(
            get_config_response,
            "GET /resource_configurations/{configurationId} response",
        )
        if result:
            changes.append(result)
    except KeyError:
        pass

    # Fix PATCH /resource_configurations/{configurationId} response (nested resource object with data wrapper)
    try:
        patch_config_response = spec["paths"][
            "/resource_configurations/{configurationId}"
        ]["patch"]["responses"]["200"]["content"]["application/json"]["schema"][
            "properties"
        ][
            "data"
        ][
            "properties"
        ][
            "resource"
        ]
        result = safe_remove_folder_id(
            patch_config_response,
            "PATCH /resource_configurations/{configurationId} response",
        )
        if result:
            changes.append(result)
    except KeyError:
        pass

    # Fix GET /resource_configurations response (list)
    try:
        list_config_response = spec["paths"]["/resource_configurations"]["get"][
            "responses"
        ]["200"]["content"]["application/json"]["schema"]["items"]["properties"][
            "resource"
        ]
        result = safe_remove_folder_id(
            list_config_response, "GET /resource_configurations response (list)"
        )
        if result:
            changes.append(result)
    except KeyError:
        pass

    return changes


def main():
    spec_file = Path("openAPISpec.json")
    spec = json.load(spec_file.open())

    changes = fix_folder_id_in_responses(spec)

    if changes:
        json.dump(spec, spec_file.open("w"), indent=2)
        print(f"✓ Removed folder_id from required in: {', '.join(changes)}")
    else:
        print("✓ No changes needed")


if __name__ == "__main__":
    main()
