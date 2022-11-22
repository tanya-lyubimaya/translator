package domain

type TranslateRequestModel struct {
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
	SourceText string `json:"source_text"`
}
