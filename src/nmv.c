#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <getopt.h>

#include "help.h"
#include "selector.h"
#include "move.h"

int main(int argc, char *argv[]) {
	// {{{1 Parse arguments
	// {{{2 Show help text if no arguments provided
	if (argc == 1) {
		print_help();
	}

	// {{{2 Variables to hold arguments
	int diff = 1;
	bool resize = false;
	Selector *from = NULL;
	Selector *to = NULL;

	// {{{2 Setup getopt
	struct option long_options[] = {
		{ "diff", 1, &diff, 0 },
		{ "resize", 0, 0, 0 },
		{ "help", 0, 0, 0 },
		{ 0, 0, 0, 0 },
	};

	// {{{2 Parse options
	char opt_c = 0;

	// Holds the number of options given so positional argument can 
	// be parsed
	int num_opts = 0;

	while (opt_c != (char)-1) {
		opt_c = (char) getopt_long(argc, argv, "+d:rh", long_options, 
				NULL);

		switch (opt_c) {
			// Help argument
			case 'h':
				print_help();
				break;

			// Difference argument
			case 'd':
				num_opts += 2;
				break;

			// Resize argument
			case 'r':
				resize = true;
				num_opts += 1;
				break;
		}
	}

	// {{{2 Parse positional arguments
	int num_pos_args = argc-(num_opts+1);

	if (num_pos_args < 2) {
		fprintf(stderr, "FROM and TO arguments must be provided\n");
		print_help();
	}

	for (int i = num_opts+1; i < argc; i++) {
		switch (i) {
			// FROM
			case 1:
				from = selector_init(argv[i]);
				break;

			// TO
			case 2:
				to = selector_init(argv[i]);
				break;
		}
	}

	// {{{1 Call move function
	// numeric_move(from, to, diff, resize);

	// {{{1 Cleanup
	selector_free(from);
	selector_free(to);

	return 0;
}
