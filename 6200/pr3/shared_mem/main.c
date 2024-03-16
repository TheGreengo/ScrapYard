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
    // I think we'll start by making a shared memory chunk
    // then we'll designate one spot for a char indicating 
    // whether it's been printed 'p' or if 'a' or 'b' is 
    // supposed to print that message
    char * name = "nonsense";
    const char * mess_a = "This should be printed by a\n";
    const char * mess_b = "This should be printed by b\n";
    int shm_fd;
    int size = 4096;
    void * ptr;

    shm_fd = shm_open(name, O_CREAT | O_EXCL |O_WRONLY, 0666);

    ftruncate(shm_fd, size);

    ptr = mmap(0, size, PROT_WRITE, MAP_SHARED, shm_fd, 0);

    sprintf(ptr, "%s", mess_a);
 
    ptr += strlen(mess_a);
    sprintf(ptr, "%s", mess_b);
    ptr += strlen(mess_b);

    sleep(2);

	// mmap cleanup
	int res = munmap(name, size);

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