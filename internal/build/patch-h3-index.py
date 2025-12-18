#!/usr/bin/env python3
import json
import sys

def patch_h3_index(obj):
    modified = False
    if isinstance(obj, dict):
        for key, value in obj.items():
            if key == "h3_index_11":
                if "anyOf" in value:
                    for item in value["anyOf"]:
                        if item.get("type") == "integer":
                            if "format" not in item or item["format"] != "int64":
                                item["format"] = "int64"
                                modified = True
                                print(f"Patched h3_index_11 in object")
            
            if isinstance(value, (dict, list)):
                if patch_h3_index(value):
                    modified = True
    elif isinstance(obj, list):
        for item in obj:
            if patch_h3_index(item):
                modified = True
    return modified

def process_file(file_path):
    try:
        with open(file_path, 'r') as f:
            data = json.load(f)
    except Exception as e:
        print(f"Error reading {file_path}: {e}")
        sys.exit(1)

    if patch_h3_index(data):
        with open(file_path, 'w') as f:
            json.dump(data, f, indent=2)
        print(f"Updated {file_path}")
    else:
        print(f"No changes needed for {file_path}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: patch-h3-index.py <openapi-file>")
        sys.exit(1)
    
    process_file(sys.argv[1])
