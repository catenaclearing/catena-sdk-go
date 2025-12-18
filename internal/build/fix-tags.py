#!/usr/bin/env python3
import json
import sys
import os

def fix_tags(file_path):
    try:
        with open(file_path, 'r') as f:
            data = json.load(f)
    except Exception as e:
        print(f"Error reading {file_path}: {e}")
        sys.exit(1)

    modified = False
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
