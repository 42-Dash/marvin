###############################################################################
######                            PROPERTIES                             ######
###############################################################################

CC			= cc
RM			= rm -rf
CFLAGS		= -Wall -Wextra -Werror -MD -MP -g
MAKEFLAGS	= -j$(nproc) --no-print-directory
INCS		= incs

###############################################################################
######                               LIBFT                               ######
###############################################################################

LIBDIR		= ./libft
LIBFT		= ${LIBDIR}/libft.a

###############################################################################
######                           OPEN LEAGUE                             ######
###############################################################################

NAME	= marvin
SRCSDIR	= srcs
SRCS	= \
		${SRCSDIR}/reader.c \
		${SRCSDIR}/dummy.c \
		${SRCSDIR}/utils.c \
		${SRCSDIR}/solver.c \
		${SRCSDIR}/init.c \
		${SRCSDIR}/main.c

OBJSDIR	= ${SRCSDIR}/objs
OBJS	= $(SRCS:${SRCSDIR}/%.c=${OBJSDIR}/%.o)
DEPS	= $(SRCS:${SRCSDIR}/%.c=${OBJSDIR}/%.d)

###############################################################################
######                              RULES                                ######
###############################################################################

all		: $(NAME)


$(NAME)	: ${OBJS}
		$(MAKE) --no-print-directory -C ${LIBDIR} all
		$(CC) -o $@ $^ -L. ${LIBFT}

${OBJSDIR}/%.o		: ${SRCSDIR}/%.c
		@mkdir -p $(dir $@)
		${CC} ${CFLAGS} -c $< -o $@ -I ./$(INCS)

clean	:
		$(MAKE) --no-print-directory -C ${LIBDIR} clean
		$(RM) $(OBJSDIR)

fclean	:
		$(MAKE) --no-print-directory -C ${LIBDIR} fclean
		$(RM) $(OBJSDIR) $(NAME)

run		:
		@$(MAKE) all
		@./$(NAME) maps/planet.txt

re		:
		$(MAKE) fclean
		$(MAKE) all

-include $(DEPS)

.PHONY: all clean fclean re run
