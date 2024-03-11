#include <stdio.h>
#include <stdlib.h>
#include <signal.h>
#include <unistd.h>

void handle_int(int signo) {
    if (signo == SIGINT) {
        printf("User signal received");
        exit(-1);
    }
}

int main(int argc, char ** argv) {
    if (signal(SIGINT, handle_int) == SIG_ERR) {
        printf("You suck");
    }

    while (1) {}

    return 0;
}