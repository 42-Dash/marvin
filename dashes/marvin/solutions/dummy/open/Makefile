###############################################################################
######                            PROPERTIES                             ######
###############################################################################

CXX			= g++
RM			= rm -rf
MAKEFLAGS	= -j$(nproc) --no-print-directory

NAME		= marvin
SRCSDIR		= srcs
SRCS		= ${SRCSDIR}/main.cpp

OBJSDIR		= objs
OBJS		= $(addprefix $(OBJSDIR)/, $(SRCS:.cpp=.o))

###############################################################################
######                               RULES                               ######
###############################################################################

all		: $(NAME)

$(NAME)	: ${OBJS}
		$(CXX) -o $@ $^

${OBJSDIR}/%.o	: %.cpp
		@mkdir -p $(dir $@)
		${CXX} -c $< -o $@

clean	:
		$(RM) $(OBJSDIR)

fclean	:
		$(RM) $(OBJSDIR) $(NAME)

re		:
		$(MAKE) fclean
		$(MAKE) all

run		:
		$(MAKE) re
		./$(NAME) maps/planet.txt

.PHONY: all clean fclean re run
