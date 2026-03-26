package main

import (
	"context"
	"encoding/base64"
	"html"
	"net/url"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Encode : フロントエンドから呼ばれるエンコード処理
func (a *App) Encode(mode int, input string) string {
	switch mode {
	case 0:
		return url.QueryEscape(input)
	case 1:
		return base64.StdEncoding.EncodeToString([]byte(input))
	case 2:
		return html.EscapeString(input)
	case 3:
		b64 := base64.StdEncoding.EncodeToString([]byte(input))
		return url.QueryEscape(b64)
	default:
		return input
	}
}

// Decode : フロントエンドから呼ばれるデコード処理
func (a *App) Decode(mode int, input string) string {
	switch mode {
	case 0:
		str, err := url.QueryUnescape(input)
		if err != nil {
			return "error: " + err.Error()
		}
		return str
	case 1:
		str, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			return "error: " + err.Error()
		}
		return string(str)
	case 2:
		return html.UnescapeString(input)
	case 3:
		str1, err := url.QueryUnescape(input)
		if err != nil {
			return "error: " + err.Error()
		}
		str2, err := base64.StdEncoding.DecodeString(str1)
		if err != nil {
			return "error: " + err.Error()
		}
		return string(str2)
	default:
		return input
	}
}
