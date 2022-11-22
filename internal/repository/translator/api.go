package translator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tanya-lyubimaya/translate/internal/domain"
	"io/ioutil"
	"net/http"
	"strings"
)

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
	req, _ := http.NewRequest("GET", r.translateUrl, nil)

	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", r.apiKey)
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return languagesRespToSlice(body)
}

func (r *repository) GetTranslation(ctx context.Context, input domain.TranslateRequestModel) (string, error) {
	payload := strings.NewReader("q=" + input.SourceText + "&target=" + input.TargetLang + "&source=" + input.SourceLang)
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

	fmt.Println(translateRespToString(body1))
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
		languages = append(languages, v.Language)
	}

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
