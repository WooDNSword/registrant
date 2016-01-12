#!/usr/bin/env python

from cfg import read_cfg
from connection import message, recv_msg, send_msg
import os
import socket

if __name__ == '__main__':
	# os.path module shit deals with Python's idiotic file path mechanisms.
	cfg = read_cfg(
		os.path.dirname(
			os.path.realpath(__file__)
		) + '/res/json/cfg.json'
	)
	
	# TODO: Move connection stuff into connection module (duh).
	sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
	sock.connect((cfg['registrar']['address'], cfg['registrar']['port']))
	
	send_msg(sock, message('IDENT', 'registrant'))
	msg = recv_msg(sock)
