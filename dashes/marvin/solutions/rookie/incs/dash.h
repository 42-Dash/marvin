#ifndef DASH_H
# define DASH_H

# include <stdbool.h>
# include <stdlib.h>
# include <limits.h>
# include <string.h>
# include <stdio.h>
# include <fcntl.h>

# include "../libft/libft.h"

// map macros
# define START_CHAR 'M'
# define FINISH_CHAR 'G'
# define NOT_NUMERIC "MG"

// directions macros
# define LEFT 'L'
# define RIGHT 'R'
# define UP 'U'
# define DOWN 'D'

# define ABS(x) ((x) < 0 ? -(x) : (x))
# define MIN(a, b) ((a) < (b) ? (a) : (b))

typedef enum e_status {
    OK,
    ERROR
}   t_status;

typedef struct s_point {
    int row, col;
}   t_point;

typedef struct s_map {
    // map atributtes
    char **map;
    int width;
    int height;

    // start and end
    int start;
    int end;
    int nodes;

    // graph
    int **graph;
    char **solutions;
    bool *visited;
    int *costs;
}   t_map;

// solvers
void print_dummy_solution(const t_map *map);
void print_djikstra_solution(t_map *map);

// solver utils
void init_map(t_map *map, char *filename);

// reader
char **read_file(int fd);

// utils
int split_len(char **map);
void debug_all(t_map *map);
void debug(t_map *map);

#endif // DASH_H
