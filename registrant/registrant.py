#!/usr/bin/env python
# TODO: Document the module!

import cfg
import conn
import os
import socket


if __name__ == '__main__':
	# TODO: Move cfg loading logic into a separate module/repo (exists in both
	# registrar and registrant).
	# os.path module shit deals with Python's idiotic file path mechanisms.
	cfg_data = cfg.read_cfg(
		os.path.dirname(
			os.path.realpath(__file__)
		) + '/res/json/cfg.json'
	)
	
	sock = conn.initiate(cfg_data)
	
	# TODO: Move session setup to session module.
	conn.send_msg(sock, conn.message('IDENT', 'registrant'))
	msg = conn.recv_msg(sock)
