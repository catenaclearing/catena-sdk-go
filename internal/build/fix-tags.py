#!/usr/bin/env python3
import json
import sys


def dedupe_list_preserve_order(items):
    seen = set()
    out = []
    for item in items:
        # JSON values here are expected to be hashable scalars (e.g., enum strings).
        # Fallback to repr for safety if a non-hashable sneaks in.
        try:
            key = item
            if key in seen:
                continue
            seen.add(key)
        except TypeError:
            key = repr(item)
            if key in seen:
                continue
            seen.add(key)
        out.append(item)
    return out


def dedupe_schema_lists(node):
    modified = False

    if isinstance(node, dict):
        for key, value in node.items():
            if key in ("enum", "required") and isinstance(value, list):
                deduped = dedupe_list_preserve_order(value)
                if deduped != value:
                    node[key] = deduped
                    modified = True
            if isinstance(value, (dict, list)):
                if dedupe_schema_lists(value):
                    modified = True
    elif isinstance(node, list):
        for item in node:
            if isinstance(item, (dict, list)):
                if dedupe_schema_lists(item):
                    modified = True

    return modified

def fix_tags(file_path):
    try:
        with open(file_path, 'r') as f:
            data = json.load(f)
    except Exception as e:
        print(f"Error reading {file_path}: {e}")
        sys.exit(1)

    modified = False
    if dedupe_schema_lists(data):
        modified = True

    if 'paths' in data:
        for path, methods in data['paths'].items():
            for method, operation in methods.items():
                if isinstance(operation, dict) and 'tags' in operation:
                    tags = operation['tags']
                    if isinstance(tags, list) and len(tags) > 1:
                        print(f"Fixing tags for {method.upper()} {path}: {tags} -> {[tags[0]]}")
                        operation['tags'] = [tags[0]]
                        modified = True

    if modified:
        with open(file_path, 'w') as f:
            json.dump(data, f, indent=2)
        print(f"Updated {file_path}")
    else:
        print(f"No changes needed for {file_path}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: fix-tags.py <openapi-file>")
        sys.exit(1)
    
    fix_tags(sys.argv[1])
