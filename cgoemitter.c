#include "cgoemitter.h"

cgoemitter_args_t cgoemitter_new_args(int args_cap) {
	cgoemitter_args_t cgoemitter_args;
	cgoemitter_args.args = malloc(args_cap * sizeof(void*));
	cgoemitter_args.args_cap = args_cap;
	cgoemitter_args.args_len = 0;
	return cgoemitter_args;
}

void* cgoemitter_args_halloc_arg(void* value, size_t size) {
	void* value_pointer = malloc(size);
	if (value_pointer != NULL) memcpy(value_pointer, value, size);
	return value_pointer;
}

int cgoemitter_args_halloc_argp(void** ptr, void* value, size_t size) { 
	*ptr = cgoemitter_args_halloc_arg(value, size);
	if (*ptr == NULL) return EXIT_FAILURE;
	return EXIT_SUCCESS;
}

int cgoemitter_args_add_arg(cgoemitter_args_t* cgoemitter_args, void* value) {
	if ((*cgoemitter_args).args_len == (*cgoemitter_args).args_cap) return EXIT_FAILURE;
	(*cgoemitter_args).args[(*cgoemitter_args).args_len++] = value;
	return EXIT_SUCCESS;
}
