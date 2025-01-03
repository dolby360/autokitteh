#!/usr/bin/env python
"""Generate requirements from pyproject.toml"""

from argparse import ArgumentParser, FileType
import tomllib

parser = ArgumentParser(description=__doc__)
parser.add_argument(
    "input",
    help="input TOML",
    default="../pyproject.toml",
    type=FileType("rb"),
    nargs="?",
)
parser.add_argument(
    "output",
    help="output requirements file",
    default="requirements.txt",
    type=FileType("w"),
    nargs="?",
)
args = parser.parse_args()

header = f"""\
# Generated by gen-reqs.py, DO NOT EDIT
# This is an export of {args.input.name} for builds in readthedocs.
"""

config = tomllib.load(args.input)
print(header, file=args.output)
for dep in config["project"]["optional-dependencies"]["all"]:
    print(dep, file=args.output)
