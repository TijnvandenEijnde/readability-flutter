#include "readability/readability.h"

// The real signatures of Parse and FreeArticle are in libreadability.h. Only
// their addresses are needed here, so they are declared without parameters.
extern void Parse(void);
extern void FreeArticle(void);

// The Go library is a static archive. Dart resolves Parse/FreeArticle at
// runtime through dlsym, so nothing in the app references them at link time.
// Under Swift Package Manager the archive is linked into the app executable,
// and release builds dead-strip every export that nothing references.
// Referencing enforce_binding pulls the Go object file into the final binary,
// and referencing Parse and FreeArticle keeps the linker from discarding them.
__attribute__((used)) void *readability_keep_symbols = (void *)&enforce_binding;
__attribute__((used)) void *readability_keep_parse = (void *)&Parse;
__attribute__((used)) void *readability_keep_free_article = (void *)&FreeArticle;
