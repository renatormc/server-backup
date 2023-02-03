#!/bin/python3
from pathlib import Path
import os 
import sys

exe_path = Path("./dist/server-backup")
link = Path("/usr/local/bin/server-backup")
try:
    os.symlink(exe_path, link)
except FileExistsError:
    res = input(f"file \"{link}\" already exists. Do you want to delete it (yes/no)?")
    if res.strip().lower() != "yes":
        sys.exit(0)
    link.unlink()
    os.symlink(exe_path, link)
    print("installed")