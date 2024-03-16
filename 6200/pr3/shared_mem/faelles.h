#include <pthread.h>

typedef struct{
    char perm;
    char message[128];
    pthread_mutex_t mut;
    pthread_cond_t red;
    pthread_cond_t writ;
} shared;