#!/usr/bin/env python

import json
import socket

# TODO: Move msg to a specialized module.
# TODO: Document msg.
def msg(msg_type, *msg_content):
	return {
		'type': msg_type,
		'content': msg_content
	}

# TODO: Move read_cfg to a specialized module.
# TODO: Document read_cfg.
def read_cfg(file_path):
	with open(file_path) as f:
		return json.loads(f.read())

# TODO: Move recv_msg to a specialized module.
# TODO: Document recv_msg.
def recv_msg(conn):
	# Retrieve JSON content length.
	json_len_buf = ''
	char_buf = conn.recv(1)
	while char_buf != '{':
		json_len_buf += char_buf
		char_buf = conn.recv(1)
	json_len = int(json_len_buf)
	
	# Retrieve raw JSON string.
	raw_msg = conn.recv(json_len)
	
	# Return decoded JSON object as dict.
	return json.loads(raw_msg)

# TODO: Move send_msg to a specialized module.
# TODO: Document send_msg.
def send_msg(conn, msg_dict):
	json_content = json.dumps(msg_dict)
	json_len = len(json_content)
	conn.sendall(str(json_len) + json_content)

if __name__ == '__main__':
	cfg = read_cfg('res/json/cfg.json')
	sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
	sock.connect((cfg['registrar']['address'], cfg['registrar']['port']))
	send_msg(sock, msg('IDENT', 'registrant'))
