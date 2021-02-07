<?php

$msg = "你说谎， 你放屁，你这个傻子";
$SOCKET_FILE = "./spamcheck.sock";
$socket = socket_create(AF_UNIX, SOCK_STREAM, 0);
socket_connect($socket, $SOCKET_FILE);
socket_send($socket, $msg, strlen($msg), 0);
$response = socket_read($socket, 1024);
socket_close($socket);

var_dump($response);