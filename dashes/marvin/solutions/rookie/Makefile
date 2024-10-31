###############################################################################
######                            PROPERTIES                             ######
###############################################################################

CC			= cc
RM			= rm -rf
MAKEFLAGS	= -j$(nproc) --no-print-directory
INCS		= incs

###############################################################################
######                               LIBFT                               ######
###############################################################################

LIBDIR		= ./libft
LIBFT		= ${LIBDIR}/libft.a

###############################################################################
######                           ROOKIE LEAGUE                           ######
###############################################################################

NAME		= marvin
SRCSDIR		= srcs
SRCS		= \
			${SRCSDIR}/reader.c \
			${SRCSDIR}/dummy.c \
			${SRCSDIR}/utils.c \
			${SRCSDIR}/solver.c \
			${SRCSDIR}/main.c

OBJSDIR		= ${SRCSDIR}/objs
OBJS		= $(SRCS:${SRCSDIR}/%.c=${OBJSDIR}/%.o)

###############################################################################
######                              RULES                                ######
###############################################################################

all		: $(NAME)

$(NAME)	: ${OBJS}
		$(MAKE) --no-print-directory -C ${LIBDIR} all
		$(CC) -o $@ $^ -L. ${LIBFT}

${OBJSDIR}/%.o		: ${SRCSDIR}/%.c
		@mkdir -p $(dir $@)
		${CC} -c $< -o $@ -I ./$(INCS)

clean	:
		$(MAKE) --no-print-directory -C ${LIBDIR} clean
		$(RM) $(OBJSDIR)

fclean	:
		$(MAKE) --no-print-directory -C ${LIBDIR} fclean
		$(RM) $(OBJSDIR) $(NAME)

run		:
		$(MAKE) all
		@./$(NAME) maps/planet.txt

re		:
		$(MAKE) fclean
		$(MAKE) all

.PHONY: all clean fclean re run
