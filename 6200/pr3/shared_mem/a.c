#include <pthread.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <sys/mman.h>
#include <unistd.h>

int main(int argc, char ** argv) {
    int size = 4096;
    int shm_fd;
    void * ptr;
    char * name = "nonsense";

    shm_fd = shm_open(name, O_RDONLY, 0666);

    ptr = mmap(0, size, PROT_READ, MAP_SHARED, shm_fd, 0);

    printf("%s", (char*)ptr);

    shm_unlink(name);

    return 0;
}