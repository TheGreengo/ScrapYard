#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <sys/mman.h>
#include <unistd.h>
#include "faelles.h"

int main(int argc, char ** argv) {
    int shm_fd;
    shared * ptr;
    char * name = "nonsense";

    shm_fd = shm_open(name, O_RDONLY, 0666);

    ptr = mmap(0, sizeof(shared), PROT_READ, MAP_SHARED, shm_fd, 0);

    printf("%s", (char*)ptr);
    printf(" - printed by process a\n");

	int res = munmap(name, sizeof(shared));
    close(shm_fd);

    return 0;
}