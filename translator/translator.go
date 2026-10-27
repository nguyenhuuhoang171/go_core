package translator

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"google.golang.org/grpc/metadata"
)

type AppTranslator struct {
	bundle          *i18n.Bundle
	defaultLanguage string
}

// NewTranslator load file <localeDir>/<lang>.json cho từng ngôn ngữ trong
// languages. defaultLanguage được dùng khi request không gửi ngôn ngữ, và là
// ngôn ngữ fallback khi một key chưa có bản dịch — nên bắt buộc phải nằm
// trong languages.
func NewTranslator(localeDir string, languages []string, defaultLanguage string) (*AppTranslator, error) {
	defaultTag, err := language.Parse(defaultLanguage)
	if err != nil {
		return nil, fmt.Errorf("invalid default language %q: %w", defaultLanguage, err)
	}
	if !slices.Contains(languages, defaultLanguage) {
		return nil, fmt.Errorf("default language %q is not in languages %v", defaultLanguage, languages)
	}

	translator := &AppTranslator{
		defaultLanguage: defaultLanguage,
	}

	// Tạo bundle với ngôn ngữ fallback là defaultLanguage
	translator.bundle = i18n.NewBundle(defaultTag)
	translator.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	// Load tất cả file translation
	for _, lang := range languages {
		file := filepath.Join(localeDir, lang+".json")
		if _, err := translator.bundle.LoadMessageFile(file); err != nil {
			return nil, fmt.Errorf("failed to load %s: %w", file, err)
		}
	}

	return translator, nil
}

func (t *AppTranslator) T(lang, key string) string {
	if lang == "" {
		lang = t.defaultLanguage
	}

	localizer := i18n.NewLocalizer(t.bundle, lang)

	// Khi lang thiếu key, go-i18n vẫn trả bản dịch của defaultLanguage kèm
	// MessageNotFoundErr — chỉ trả về key khi không có bản dịch nào.
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})
	if err != nil && msg == "" {
		return key
	}
	return msg
}

// GetLanguageFromContext extracts language from gRPC metadata in context
// Returns the language code (e.g., "en", "vi") or empty string if not found
func (t *AppTranslator) GetLanguageFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	// Check for "language" key in metadata (case-insensitive)
	langs := md.Get("language")
	if len(langs) == 0 {
		// Also check for "lang" as alternative
		langs = md.Get("lang")
	}
	if len(langs) == 0 {
		return ""
	}

	// Get first language value and normalize (lowercase, trim)
	lang := strings.ToLower(strings.TrimSpace(langs[0]))
	return lang
}
