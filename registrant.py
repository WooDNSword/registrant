#!/usr/bin/env python

import json

# TODO: Move read_cfg to a specialized module.
# TODO: Document read_cfg.
def read_cfg(file_path):
	with open(file_path) as f:
		return json.loads(f.read())

if __name__ == "__main__":
	cfg = read_cfg('res/json/cfg.json')
	# TODO: Initiate connection to WDNS Registrar.
