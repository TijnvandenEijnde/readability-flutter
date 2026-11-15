#include "readability/readability.h"

// The Go library is a static archive. Dart resolves Parse/FreeArticle at
// runtime through dlsym, so nothing in the app references them at link time.
// Referencing enforce_binding pulls the Go object file into the final binary
// and keeps the linker from discarding it.
__attribute__((used)) void *readability_keep_symbols = (void *)&enforce_binding;
