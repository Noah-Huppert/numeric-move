#ifndef __SELECTOR_H__
#define __SELECTOR_H__

/**
 * Holds information about files to operate on
 */
typedef struct {
	/**
	 * Path to files directory
	 */
	char *directory;

	/**
	 * Number of digits used to express numeric prefix
	 */
	int num_digs;

	/**
	 * Numeric prefix
	 */
	int prefix;
} Selector;

/**
 * Create selector from selector string
 */
Selector* selector_init(char *sstr);

/**
 * Free selector
 */
void selector_free(Selector *s);

#endif
