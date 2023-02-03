#!/bin/python3
from pathlib import Path
import subprocess
import sys

current_dir = Path(".")
exe_path = Path("./dist/server-backup")
subprocess.check_call(['go1.19', "build", "-o", str(exe_path)])

if len(sys.argv) > 1:
    subprocess.check_call([str(exe_path)] + sys.argv[1:])

