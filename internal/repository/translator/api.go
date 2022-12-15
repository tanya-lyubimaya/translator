package translator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tanya-lyubimaya/translator/internal/domain"
	"io/ioutil"
	"net/http"
	"strings"
)

var languageMap = map[string]string{
	"as":         "Assamese",
	"ay":         "Aymara",
	"bho":        "Bhojpuri",
	"bm":         "Bambara",
	"ceb":        "Cebuano",
	"doi":        "Dogri",
	"dv":         "Divehi",
	"gom":        "Goan Konkani",
	"he":         "Hebrew",
	"hmn":        "Hmong",
	"ilo":        "Iloko",
	"jv":         "Javanese",
	"lb":         "Luxembourgish",
	"lus":        "Lushai",
	"mai":        "Maithili",
	"my":         "Burmese",
	"pt":         "Portuguese",
	"sa":         "Sanskrit",
	"sm":         "Samoan",
	"ts":         "Tsonga",
	"zsh":        "Zhuang",
	"af":         "Afrikaans",
	"ak":         "Akan",
	"sq":         "Albanian",
	"am":         "Amharic",
	"ar":         "Arabic",
	"hy":         "Armenian",
	"az":         "Azerbaijani",
	"eu":         "Basque",
	"be":         "Belarusian",
	"bem":        "Bemba",
	"bn":         "Bengali",
	"bh":         "Bihari",
	"xx-bork":    "Bork, bork, bork!",
	"bs":         "Bosnian",
	"br":         "Breton",
	"bg":         "Bulgarian",
	"km":         "Cambodian",
	"ca":         "Catalan",
	"chr":        "Cherokee",
	"ny":         "Chichewa",
	"zh-CN":      "Chinese (Simplified)",
	"zh-TW":      "Chinese (Traditional)",
	"co":         "Corsican",
	"hr":         "Croatian",
	"cs":         "Czech",
	"da":         "Danish",
	"nl":         "Dutch",
	"xx-elmer":   "Elmer Fudd",
	"en":         "English",
	"eo":         "Esperanto",
	"et":         "Estonian",
	"ee":         "Ewe",
	"fo":         "Faroese",
	"tl":         "Filipino",
	"fi":         "Finnish",
	"fr":         "French",
	"fy":         "Frisian",
	"gaa":        "Ga",
	"gl":         "Galician",
	"ka":         "Georgian",
	"de":         "German",
	"el":         "Greek",
	"gn":         "Guarani",
	"gu":         "Gujarati",
	"xx-hacker":  "Hacker",
	"ht":         "Haitian Creole",
	"ha":         "Hausa",
	"haw":        "Hawaiian",
	"iw":         "Hebrew",
	"hi":         "Hindi",
	"hu":         "Hungarian",
	"is":         "Icelandic",
	"ig":         "Igbo",
	"id":         "Indonesian",
	"ia":         "Interlingua",
	"ga":         "Irish",
	"it":         "Italian",
	"ja":         "Japanese",
	"jw":         "Javanese",
	"kn":         "Kannada",
	"kk":         "Kazakh",
	"rw":         "Kinyarwanda",
	"rn":         "Kirundi",
	"xx-klingon": "Klingon",
	"kg":         "Kongo",
	"ko":         "Korean",
	"kri":        "Krio (Sierra Leone)",
	"ku":         "Kurdish",
	"ckb":        "Kurdish (Soranî)",
	"ky":         "Kyrgyz",
	"lo":         "Laothian",
	"la":         "Latin",
	"lv":         "Latvian",
	"ln":         "Lingala",
	"lt":         "Lithuanian",
	"loz":        "Lozi",
	"lg":         "Luganda",
	"ach":        "Luo",
	"mk":         "Macedonian",
	"mg":         "Malagasy",
	"ms":         "Malay",
	"ml":         "Malayalam",
	"mt":         "Maltese",
	"mi":         "Maori",
	"mr":         "Marathi",
	"mfe":        "Mauritian Creole",
	"mo":         "Moldavian",
	"mn":         "Mongolian",
	"sr-ME":      "Montenegrin",
	"ne":         "Nepali",
	"pcm":        "Nigerian Pidgin",
	"nso":        "Northern Sotho",
	"no":         "Norwegian",
	"nn":         "Norwegian (Nynorsk)",
	"oc":         "Occitan",
	"or":         "Oriya",
	"om":         "Oromo",
	"ps":         "Pashto",
	"fa":         "Persian",
	"xx-pirate":  "Pirate",
	"pl":         "Polish",
	"pt-BR":      "Portuguese (Brazil)",
	"pt-PT":      "Portuguese (Portugal)",
	"pa":         "Punjabi",
	"qu":         "Quechua",
	"ro":         "Romanian",
	"rm":         "Romansh",
	"nyn":        "Runyakitara",
	"ru":         "Russian",
	"gd":         "Scots Gaelic",
	"sr":         "Serbian",
	"sh":         "Serbo-Croatian",
	"st":         "Sesotho",
	"tn":         "Setswana",
	"crs":        "Seychellois Creole",
	"sn":         "Shona",
	"sd":         "Sindhi",
	"si":         "Sinhalese",
	"sk":         "Slovak",
	"sl":         "Slovenian",
	"so":         "Somali",
	"es":         "Spanish",
	"es-419":     "Spanish (Latin American)",
	"su":         "Sundanese",
	"sw":         "Swahili",
	"sv":         "Swedish",
	"tg":         "Tajik",
	"ta":         "Tamil",
	"tt":         "Tatar",
	"te":         "Telugu",
	"th":         "Thai",
	"ti":         "Tigrinya",
	"to":         "Tonga",
	"lua":        "Tshiluba",
	"tum":        "Tumbuka",
	"tr":         "Turkish",
	"tk":         "Turkmen",
	"tw":         "Twi",
	"ug":         "Uighur",
	"uk":         "Ukrainian",
	"ur":         "Urdu",
	"uz":         "Uzbek",
	"vi":         "Vietnamese",
	"cy":         "Welsh",
	"wo":         "Wolof",
	"xh":         "Xhosa",
	"yi":         "Yiddish",
	"yo":         "Yoruba",
	"zu":         "Zulu",
}

type repository struct {
	translateUrl string
	apiKey       string
}

func New(translateUrl string, apiKey string) (*repository, error) {
	return &repository{
		translateUrl: translateUrl,
		apiKey:       apiKey,
	}, nil
}

func (r *repository) GetLanguages(ctx context.Context) ([]string, error) {
	req, _ := http.NewRequest("GET", r.translateUrl+"/languages", nil)

	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", r.apiKey)
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return languagesRespToSlice(body)
}

func (r *repository) GetTranslation(ctx context.Context, input domain.TranslateRequestModel) (string, error) {
	targetLang, ok := mapkey(languageMap, input.TargetLang)
	if !ok {
		return "", errors.New("Unknown language to translate to")
	}

	sourceLang, ok := mapkey(languageMap, input.SourceLang)
	if !ok {
		return "", errors.New("Unknown language to translate from")
	}

	payload := strings.NewReader("q=" + input.SourceText + "&target=" + targetLang + "&source=" + sourceLang)
	req, err := http.NewRequest("POST", r.translateUrl, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", r.apiKey)
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body1, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		return "", errors.New("Too many requests")
	}

	//fmt.Println(translateRespToString(body1))
	return translateRespToString(body1)
}

func languagesRespToSlice(body []byte) ([]string, error) {
	type languagesResp struct {
		Data struct {
			Languages []struct {
				Language string `json:"language"`
			} `json:"languages"`
		} `json:"data"`
	}
	var languages []string
	var lr languagesResp

	err := json.Unmarshal(body, &lr)
	if err != nil {
		return languages, err
	}
	for _, v := range lr.Data.Languages {
		if val, ok := languageMap[v.Language]; val != "" || !ok {
			languages = append(languages, languageMap[v.Language])
		} else {
			fmt.Println("Found unsupported language: ", v.Language)
		}
	}

	languages = removeDuplicates(languages)

	return languages, nil
}

func translateRespToString(body []byte) (string, error) {
	type translateResp struct {
		Data struct {
			Translations []struct {
				TranslatedText string `json:"translatedText"`
			} `json:"translations"`
		} `json:"data"`
	}
	var tr translateResp
	var translatedText string

	err := json.Unmarshal(body, &tr)
	if err != nil {
		return translatedText, err
	}
	translatedText += tr.Data.Translations[0].TranslatedText
	return translatedText, nil
}

func removeDuplicates[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func mapkey(m map[string]string, value string) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
