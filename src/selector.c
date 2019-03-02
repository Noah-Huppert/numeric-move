#include "selector.h"
#include <stdio.h>
#include <ctype.h>

Selector* selector_init(char *sstr) {
	// {{{1 Find where the path section ends
	int sstr_len = 0;
	int path_end = 0;

	for (int i = 0; sstr[i] != '\0'; i++) {
		if (sstr[i] == '/') {
			path_end = i;
		}

		sstr_len++;
	}

	// {{{1 Determine size of numeric prefix
	int num_digs = sstr_len - (path_end+1);

	// {{{1 Check numeric prefix is made up of numbers
	for (int i = 0; i < num_digs; i++) {
		printf("%d: %c\n", path_end+i, sstr[path_end+i]);
	}

	printf("sstr_len: %d, path_end: %d, num_digs: %d\n", sstr_len, 
			path_end, num_digs);

	return NULL;
}

void selector_free(Selector *s) {
}
