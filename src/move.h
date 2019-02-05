#ifndef __MOVE_H__
#define __MOVE_H__

#include <stdbool.h>

/**
 * Move files with numeric prefix
 * @param from_dir Directory from files are located in, NULL if current 
 *                 working directory
 * @param from Numeric prefix of files to move
 * @param to_dir Directory to files should be located in, NULL if current
 *               working directory
 * @param to New numeric prefix of files
 * @param diff Difference in file numeric prefixes to preserve
 * @param resize If numeric prefixes on files should be resized to accommodate 
 *               to argument
 * @param same If files with the same numeric prefix can exist at the same time
 */
void numeric_move(char *from_dir, int from, char *to_dir, int to, int diff,
		bool resize, bool same);

#endif
