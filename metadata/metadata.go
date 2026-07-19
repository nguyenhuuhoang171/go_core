package metadata

import "context"

type langKey struct{}

// WithLang attaches a request's language to the context.
func WithLang(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, langKey{}, lang)
}

// GetLanguage returns the language previously attached via WithLang, or ""
// if none was set.
func GetLanguage(ctx context.Context) string {
	lang, _ := ctx.Value(langKey{}).(string)
	return lang
}
