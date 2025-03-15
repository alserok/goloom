package pages

import (
	"bytes"
	"github.com/alserok/goloom/internal/utils"
	"html/template"
)

type HTMLConstructor interface {
	Render(page uint, data any) ([]byte, error)
}

const (
	PageConfig = iota
	PageDir
	PageState
)

func NewConstructor() HTMLConstructor {
	return &htmlConstructor{
		state:  newStatePage(),
		dir:    newDirPage(),
		config: newConfigPage(),
	}
}

type htmlConstructor struct {
	config *template.Template
	state  *template.Template
	dir    *template.Template
}

func (c *htmlConstructor) Render(page uint, data any) ([]byte, error) {
	output := bytes.NewBuffer([]byte{})

	switch page {
	case PageConfig:
		if err := c.config.ExecuteTemplate(output, "config", data); err != nil {
			return nil, utils.NewError(err.Error(), utils.ErrInternal)
		}
	case PageDir:
		if err := c.dir.ExecuteTemplate(output, "dir", data); err != nil {
			return nil, utils.NewError(err.Error(), utils.ErrInternal)
		}
	case PageState:
		if err := c.state.ExecuteTemplate(output, "state", data); err != nil {
			return nil, utils.NewError(err.Error(), utils.ErrInternal)
		}
	default:
		panic("invalid page type")
	}

	return output.Bytes(), nil
}

type Data map[string]any
