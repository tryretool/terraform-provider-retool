#!/usr/bin/env python3
"""Transform a raw Retool OpenAPI spec into the form expected by the Go generator.

Usage:
    python3 transform_spec.py <raw_spec.json> [output.json]

Defaults output to openAPISpec.json (the file consumed by `go generate`).

The transformations applied here are required for OpenAPI Generator to produce
compiling, working Go code against the Retool API. See AGENTS.md / README.md for
the full background. They are intentionally written semantically (rather than by
JSON path) so they keep working as the upstream spec is reordered between
versions.

Transforms:
  1. Single tag per operation. Multiple tags make the generator emit duplicate
     types and break compilation. Keep the tag that matches the top-level path
     segment when present, otherwise the first tag.
  2. Permissions oneOf -> anyOf. /permissions/listObjects, /permissions/grant and
     /permissions/revoke return arrays whose variants are structurally identical
     (only the `type` enum differs). oneOf makes the generated unmarshaler fail
     with "data matches more than one schema"; anyOf works.
  3. Free-form resource `options` request bodies. The provider treats a
     resource's / resource configuration's `options` as an opaque JSON blob, but
     the spec models it as a large `anyOf` union. The generated Go client
     ambiguously matches "thin" option objects (e.g. bearer-token auth) to the
     wrong union member and silently drops fields such as `base_url`. Replace the
     `options` request schema with a free-form object so it passes through
     untouched. (Responses are left typed.)
  4. Optional response fields. `folder_id`, `seat_type` and `default_value` are
     marked required on several response schemas but the API omits them on
     older / not-yet-migrated instances, breaking unmarshaling. Remove them from
     every `required` array (they stay in `properties`). None of these are
     required in request bodies, so removing them globally is safe.
"""
import json
import sys
from pathlib import Path

METHODS = ("get", "post", "put", "patch", "delete")
PERMISSIONS_ENDPOINTS = (
    "/permissions/listObjects",
    "/permissions/grant",
    "/permissions/revoke",
)


def fix_tags(spec):
    changes = 0
    for path, path_item in spec.get("paths", {}).items():
        segments = [p for p in path.split("/") if p and not p.startswith("{")]
        top_level = segments[0].capitalize() if segments else None
        for method in METHODS:
            op = path_item.get(method)
            if not op or "tags" not in op:
                continue
            tags = op["tags"]
            if len(tags) > 1:
                op["tags"] = [top_level] if top_level in tags else [tags[0]]
                changes += 1
    return changes


def fix_permissions_oneof_to_anyof(spec):
    changes = 0
    for endpoint in PERMISSIONS_ENDPOINTS:
        try:
            items = spec["paths"][endpoint]["post"]["responses"]["200"]["content"][
                "application/json"
            ]["schema"]["properties"]["data"]["items"]
        except (KeyError, TypeError):
            continue
        if "oneOf" in items:
            items["anyOf"] = items.pop("oneOf")
            changes += 1
    return changes


def fix_freeform_resource_options(spec):
    """Replace the `options` request schema of resources / resource configurations
    with a free-form object (the provider treats options as opaque JSON)."""
    changes = 0
    for path, item in spec.get("paths", {}).items():
        if not path.startswith(("/resources", "/resource_configurations")):
            continue
        for method in METHODS:
            op = item.get(method)
            if not isinstance(op, dict):
                continue
            content = op.get("requestBody", {}).get("content", {})
            for media in content.values():
                schema = media.get("schema")
                if not isinstance(schema, dict):
                    continue
                props = schema.get("properties")
                if not isinstance(props, dict):
                    continue
                opt = props.get("options")
                if isinstance(opt, dict) and ("anyOf" in opt or "oneOf" in opt):
                    new_opt = {"type": "object", "additionalProperties": True}
                    if "description" in opt:
                        new_opt["description"] = opt["description"]
                    props["options"] = new_opt
                    changes += 1
    return changes


def _walk(obj):
    if isinstance(obj, dict):
        yield obj
        for value in obj.values():
            yield from _walk(value)
    elif isinstance(obj, list):
        for value in obj:
            yield from _walk(value)


# Fields that the spec marks required on responses but the API may omit on
# instances that have not migrated to the relevant feature yet. None of these
# are required in request bodies.
OPTIONAL_RESPONSE_FIELDS = ("folder_id", "seat_type", "default_value")


def fix_optional_response_fields(spec):
    changes = 0
    for node in _walk(spec):
        if isinstance(node, dict) and isinstance(node.get("required"), list):
            for field in OPTIONAL_RESPONSE_FIELDS:
                if field in node["required"]:
                    node["required"].remove(field)
                    changes += 1
    return changes


def main():
    if len(sys.argv) < 2:
        print(__doc__)
        sys.exit(1)
    raw_path = Path(sys.argv[1])
    out_path = Path(sys.argv[2]) if len(sys.argv) > 2 else Path("openAPISpec.json")

    spec = json.loads(raw_path.read_text())

    summary = {
        "single-tag operations": fix_tags(spec),
        "permissions oneOf->anyOf": fix_permissions_oneof_to_anyof(spec),
        "free-form resource options requests": fix_freeform_resource_options(spec),
        "optional response fields relaxed": fix_optional_response_fields(spec),
    }

    out_path.write_text(json.dumps(spec, indent=2) + "\n")
    print(f"Transformed {raw_path} -> {out_path}")
    for label, count in summary.items():
        print(f"  {label}: {count}")


if __name__ == "__main__":
    main()
