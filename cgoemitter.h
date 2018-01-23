#ifndef CGOEMITTER_H_
#define CGOEMITTER_H_

#include <stdlib.h> //EXIT_SUCCESS, EXIT_FAILURE
#include <stddef.h> //NULL
#include <string.h> //memcpy

typedef struct cgoemitter_args_t {
	int args_cap;
	int args_len;
	void** args;
} cgoemitter_args_t;

void emit(char* event_name, cgoemitter_args_t* args);

cgoemitter_args_t cgoemitter_new_args(int args_cap);

void* cgoemitter_args_halloc_arg(void* value, size_t size);

int cgoemitter_args_halloc_argp(void **ptr, void* value, size_t size);

int cgoemitter_args_add_arg(cgoemitter_args_t* cgoemitter_args, void* value);

#endif