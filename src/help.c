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
	       "                        between prefixes. Defaults to 1.\n"
	       "    --resize,-r         Resize number of digits in numeric \n"
	       "                        prefix to hold new TO number. By \n"
	       "                        default disabled.\n"
	       "    --same,-s           Allow files with same numeric prefix\n"
	       "                        to exist. By default disabled.\n"
	       "    --help,-h           Print help text\n"
	       "\n"
	       "ARGUMENTS\n"
	       "\n"
	       "    FROM        Selector of files to move\n"
	       "    TO          Selector of location to move files to\n"
	       "\n"
	       "SELECTORS\n"
	       "\n"
	       "    The FROM and TO arguments are selectors.\n"
	       "\n"
	       "    Selectors contain an optional file path and a required\n"
	       "    numeric prefix:\n"
	       "\n"
	       "        [PATH/]N_PREFIX\n"
	       "\n"
	       "    If the file path is not given it defaults to the \n"
	       "    current working directory. If given the file path must\n"
	       "    end with a forward slash.\n");
	exit(1);
}
