#ifndef __MOVE_H__
#define __MOVE_H__

#include <stdbool.h>

#include "selector.h"

/**
 * Move files with numeric prefix
 * @param from Selector of files to move
 * @param to Selector of location to move files
 * @param diff Difference in file numeric prefixes to preserve
 * @param resize If numeric prefixes on files should be resized to accommodate 
 *               to argument
 * @param same If files with the same numeric prefix can exist at the same time
 */
void numeric_move(selector *from, selector *to, int diff,
		bool resize, bool same);

#endif
