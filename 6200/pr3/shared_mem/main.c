#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <sys/mman.h>
#include <unistd.h>

#define SHAR_MEM "/nonsense/shared"
#define BUF_SIZE 4096

int main(int argc, char ** argv) {
    // I think we'll start by making a shared memory chunk
    // then we'll designate one spot for a char indicating 
    // whether it's been printed 'p' or if 'a' or 'b' is 
    // supposed to print that message
    const char ** mess_a = "This should be printed by a";
    const char ** mess_b = "This should be printed by b";
    int shm_fd;
    void * ptr;

    shm_fd = shm_open(SHAR_MEM, O_RDONLY, 0666);

    ptr = mmap(0, BUF_SIZE, PROT_READ, MAP_SHARED, shm_fd, 0);

    printf("%s", (char*)ptr);
    
    shm_unlink(SHAR_MEM);
    // pthread_mutex_t mut;
    // pthread_cond_t read;
    // pthread_cond_t writ;

    // pthread_mutex_init(&mut, NULL);
    // pthread_cond_init(&read, NULL);
    // pthread_cond_init(&writ, NULL);



    // pthread_mutex_destroy(&mut);
    // pthread_cond_destroy(&read);
    // pthread_cond_destroy(&writ);

    return 0;
}