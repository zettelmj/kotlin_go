#include <stdlib.h>
#include <stdio.h>
#include <string.h>

char* helloLib(char* str) {
      int current_len = strlen(str);
      int new_size = current_len + 7;

      char* new_str = realloc(str, new_size * sizeof(char));
      if (new_str == NULL) {
        // Handle realloc failure (e.g., print error message)
        return NULL;
      }

      str = new_str;

      str[current_len] = ' ';
      str[current_len + 1] = 'H';
      str[current_len + 2] = 'a';
      str[current_len + 3] = 'l';
      str[current_len + 4] = 'l';
      str[current_len + 5] = 'o';
      str[current_len + 6] = '\0';

      return str;
}