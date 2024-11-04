#include "dash.h"

static char **append(char **map, char *line) {
    char **new;
    int len;

    len = split_len(map);
    new = (char **)malloc((len + 2) * sizeof (char *));
    if (new == NULL) {
        exit(1);
    }
    memcpy(new, map, sizeof (char *) * len);
    new[len] = line;
    new[len + 1] = NULL;

    return new;
}

char **read_file(int fd) {
    char **res = NULL;
    char *line = NULL;

    while (true) {
        line = get_next_line(fd, O_RDONLY);
        if (line == NULL) {
            break;
        }
        res = append(res, line);
        if (res == NULL) {
            exit(1);
        }
    }
    return res;
}
