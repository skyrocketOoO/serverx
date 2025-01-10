import os
import re

def check_enums_start_at_one(file_path):
    """
    Check if enums in the given Go file start at 1.
    """
    with open(file_path, "r") as f:
        lines = f.readlines()

    in_const_block = False
    for line in lines:
        line = line.strip()

        # Detect the start of a const block
        if line.startswith("const ("):
            in_const_block = True
            continue

        # Detect the end of a const block
        if in_const_block and line.endswith(")"):
            in_const_block = False

        # Look for enums using iota
        if in_const_block and "iota" in line:
            match = re.search(r"(\w+)\s*=\s*iota", line)
            if match:
                first_enum = match.group(1)
                if "= iota" in line and not "iota + 1" in line:
                    print(f"Error: Enum '{first_enum}' starts at 0 in {file_path}.")
                    return False
    return True

def check_new_mutex_usage(file_path):
    """
    Check if `new(sync.Mutex)` is used in the given Go file.
    """
    with open(file_path, "r") as f:
        lines = f.readlines()

    for line in lines:
        line = line.strip()

        # Check for usage of `new(sync.Mutex)`
        if "new(sync.Mutex)" in line:
            print(f"Warning: Unnecessary use of `new(sync.Mutex)` in {file_path}. Consider using the zero value.")
            return False
    return True

def check_no_panic(file_path):
    """
    Check if `panic` is used in the given Go file.
    """
    with open(file_path, "r") as f:
        lines = f.readlines()

    for line in lines:
        line = line.strip()

        if "panic(" in line:
            print(f"Warning: don't use panic")
            return False
    return True

def check_has_init_function(file_path):
    """
    Check if the Go file contains an `init()` function.
    """
    with open(file_path, "r") as f:
        lines = f.readlines()

    for line in lines:
        if line.strip().startswith("func init()"):
            print(f"Found init() function in {file_path}.")
            return False
    return True

def scan_directory(directory):
    """
    Recursively scan a directory for Go files and check enums.
    """
    for root, _, files in os.walk(directory):
      if 'docs' in root:
        continue
      if 'cmd' in root:
        for file in files:
          if file.endswith(".go"):
            file_path = os.path.join(root, file)
            if not check_enums_start_at_one(file_path):
              print(f"Issue found in {file_path}.")

            if not check_new_mutex_usage(file_path):
              print(f"Issue found in {file_path} regarding Mutex initialization.")

            if not check_no_panic(file_path):
              print(f"Issue found in {file_path} regarding panic usage.")
      else:
        for file in files:
          if file.endswith(".go"):
            file_path = os.path.join(root, file)
            if not check_has_init_function(file_path):
              print(f"Issue: init() function is missing in {file_path}.")
              
              
# Main entry point
if __name__ == "__main__":
    directory_to_scan = "./"  # Change this to the directory you want to scan
    scan_directory(directory_to_scan)
