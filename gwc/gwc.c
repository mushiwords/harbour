

#include <stdio.h>
#include <sys/wait.h>
//#include <stdatomic.h>
#include <stdint.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <errno.h>
#include <sys/signal.h>
#include <sys/stat.h>

int exec_ok(char *filename) {
    struct stat st;
    if (stat(filename, &st) < 0) {
        errno = ENOENT;
        return 0;
    }
    if (S_ISDIR(st.st_mode)) {
        errno = EISDIR;
        return 0;
    }
    if (!(st.st_mode & (S_IXUSR | S_IXGRP | S_IXOTH))) {
        errno = EACCES;
        return 0;
    }
    return 1;
}

char *exec_abspath(char *filename) {
    static char abspath[BUFSIZ];

    if (!filename) {
        errno = ENOENT;
        return NULL;
    }

    if (*filename == '.' || *filename == '/') {
        switch (exec_ok(filename)) {
            case  1: return realpath(filename, abspath);
            default: return NULL;
        }
    }

    char *path = getenv("PATH");
    if (!path) return NULL;

    char *v[(BUFSIZ/2)], **paths = v;
    char *p = strdup(path);
    path = p;

    while (*p) {
        while (*p && *p == ':') *p++ = 0;
        if (*p) *(paths++) = p;
        while (*p && *p != ':') p++;
    }
    *paths = NULL;

    int l, fnlen = strlen(filename);
    for (paths = v; *paths; paths++) {
        l = strlen(*paths) + fnlen + 2;
        if (l > sizeof(abspath)) continue;
        snprintf(abspath, l, "%s/%s", *paths, filename);
        if (exec_ok(abspath)) {
            free(path);
            return abspath;
        }
    }
    free(path);
    return NULL;
}


volatile uint32_t child_exit = 0;

/** compare and set **/
static inline uint32_t cas(volatile uint32_t *v,
        uint32_t old, uint32_t new)
{
    uint8_t ret;

    __asm__ volatile (
            "lock;"
            " cmpxchgl %3, %1; "
            " sete %0; "

            : "=a" (ret) : "m" (*v), "a" (old), "r" (new) : "cc", "memory");

    return ret;
}

void singal_child(int signo)
{
    int status;
    waitpid(-1, &status, 0);

    if (WIFSIGNALED(status)) {  /** 收到子进程异常退出 **/
        child_exit = 1;
    } else {                    /** 子进程主动退出 **/
        exit(0);
    }
}

int main(int argc, char **argv)
{
    if (argc == 1) {
        printf("%s <program> <args...>\n", argv[0]);
        return -1;
    }

    char *args[argc];
    int i;
    char *prog = exec_abspath(argv[1]);
    pid_t pid;

    if (prog == NULL) {
        printf("Invalid filepath: %s\n", argv[1]);
        return -1;
    }

    for ( i = 1; i < argc; i ++) {
        args[i-1] = argv[i];
    }

    args[argc-1] = NULL;

    daemon(1, 1);

    signal(SIGCHLD, singal_child);

start_again:
    if ((pid = fork()) < 0) { /** error **/
        printf("error: %s\n", strerror(errno));
        return -1;
    } else if (pid == 0) { /** child **/
        int ret = execvp(prog, args);
        printf("ret = %d, error = %s\n", ret, strerror(errno));
        return -1;
    }

    while (1)  {
        if (cas(&child_exit, 1, 0) == 1) {
            goto start_again;
        }
        sleep(1);
    }

    return 0;
}
