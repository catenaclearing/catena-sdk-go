import json
import sys

def check_h3_index(obj, path=""):
    if isinstance(obj, dict):
        for key, value in obj.items():
            current_path = f"{path}.{key}" if path else key
            if key == "h3_index_11":
                if "anyOf" in value:
                    for i, item in enumerate(value["anyOf"]):
                        if item.get("type") == "integer":
                            if item.get("format") != "int64":
                                print(f"Missing format: int64 at {current_path}.anyOf[{i}]")
            
            if isinstance(value, (dict, list)):
                check_h3_index(value, current_path)
    elif isinstance(obj, list):
        for i, item in enumerate(obj):
            check_h3_index(item, f"{path}[{i}]")

def process_file(file_path):
    with open(file_path, 'r') as f:
        data = json.load(f)
    check_h3_index(data)

if __name__ == "__main__":
    process_file(sys.argv[1])
