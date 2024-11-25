
NAME = marvin

all:
	go build -gcflags -m -o $(NAME) ./...

fclean:
	rm -f $(NAME)

run:
	$(MAKE) all
	time ./$(NAME) input/test.txt

.PHONY: all fclean run
