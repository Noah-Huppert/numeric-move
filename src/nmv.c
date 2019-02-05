#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <getopt.h>

#include "help.h"

int main(int argc, char *argv[]) {
	// {{{1 Parse arguments
	// {{{2 Show help text if no arguments provided
	if (argc == 1) {
		print_help();
	}

	// {{{2 Variables to hold arguments
	int diff = 1;
	bool resize = false;
	int from;
	int to;

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
	int pos_arg_i = 0;
	for (int i = num_opts+1; i < argc; i++) {
		pos_arg_i++;

		switch (pos_arg_i) {
			// FROM argument
			case 1:
				from = atoi(argv[i]);
				break;

			// TO argument
			case 2:
				to = atoi(argv[i]);
				break;

			default:
				fprintf(stderr, "Only takes 2 positional "
						"arguments\n");
				print_help();
				break;
		}
	}

	if (pos_arg_i < 2) {
		fprintf(stderr, "FROM TO position arguments must be provided\n");
		print_help();
	}

	return 0;
}
