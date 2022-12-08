package cmd

import "github.com/tanya-lyubimaya/translator/internal/app"

func Execute() error {
	application, err := app.New()
	if err != nil {
		return err
	}
	err = application.Serve(":8080")
	if err != nil {
		return err
	}
	return nil
}
