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

// Our shared memory should have:
// - a permission character stating who should be doing stuff
// - a message space
// - a mutex for controlling things
// - a condition variable for editing things
// - a condition variable for reading things

int main(int argc, char ** argv) {
    char * name = "nonsense";
    const char * mess_a = "The very first message";
    const char * mess_b = "The second message";
    int shm_fd;
    int size = 4096;
    shared * ptr;
    pthread_mutexattr_t m_attr;
    pthread_condattr_t c_attr;
    pthread_mutexattr_init(&m_attr);
    pthread_mutexcond_init(&c_attr);
    pthread_mutexattr_setpshared(&m_attr, PTHREAD_PROCESS_SHARED);
    pthread_condattr_setpshared(&c_attr, PTHREAD_PROCESS_SHARED);

    shm_fd = shm_open(name, O_CREAT | O_EXCL |O_WRONLY, 0666);

    ftruncate(shm_fd, size);

    ptr = mmap(0, size, PROT_WRITE, MAP_SHARED, shm_fd, 0);

    pthread_mutex_init(&ptr->mut, &m_attr);
    pthread_cond_init(&ptr->red, &c_attr);
    pthread_cond_init(&ptr->writ, &c_attr);

    /*
    * We'll first have a loop for if the permission is not 'a' or
    * 'b' when it is, we'll write the 'a' message, set the permission
    * to 'a'
    */

    sprintf(ptr, "%s", mess_a);
    ptr += strlen(mess_a);

    /*
    * Then wait for the permission to be 'p' again
    * then we'll write the 'b' message, then set the permission to
    * 'b'
    */

    sprintf(ptr, "%s", mess_b);
    ptr += strlen(mess_b);

    /*
    * Wait for it to be p again. When the bit is 'p', we 
    * then we'll unmap, close up, unlink, and destroy the mutexes
    */

	// mmap cleanup
	int res = munmap(name, size);
    close(shm_fd);
    shm_unlink(name);

    // pthread_mutex_destroy(&mut);
    // pthread_cond_destroy(&read);
    // pthread_cond_destroy(&writ);

    return 0;
}