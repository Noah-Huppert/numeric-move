#include <stdio.h>
#include <stdlib.h>

#include "help.h"

void print_help() {
	printf("nmv - Numeric move tool\n"
		"\n"
	       "USAGE\n"
	       "\n"
	       "    nmv [OPTIONS] FROM TO\n"
	       "\n"
	       "OPTIONS\n"
	       "\n"
	       "    --diff,-d DIFF      Numeric difference to enforce \n"
	       "                        between prefixes\n"
	       "    --resize,-r         Resize number of digits in numeric \n"
	       "                        prefix to\n"
	       "                        hold new TO number\n"
	       "    --help,-h           Print help text\n"
	       "\n"
	       "ARGUMENTS\n"
	       "\n"
	       "    FROM    Numeric prefix of files to move\n"
	       "    TO      New numeric prefix for files\n");
	exit(1);
}
