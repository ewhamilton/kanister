package param

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/pkg/errors"

	crv1alpha1 "github.com/kanisterio/kanister/pkg/apis/cr/v1alpha1"
)

// RenderArgs function renders the arguments required for execution
func RenderArgs(args []string, tp TemplateParams) ([]string, error) {
	ras := make([]string, 0, len(args))
	for _, a := range args {
		ra, err := render(a, tp)
		if err != nil {
			return nil, err
		}
		ras = append(ras, ra)
	}
	return ras, nil
}

// RenderArtifacts function renders the artifacts required for execution
func RenderArtifacts(arts map[string]crv1alpha1.Artifact, tp TemplateParams) (map[string]crv1alpha1.Artifact, error) {
	rarts := make(map[string]crv1alpha1.Artifact, len(arts))
	for name, a := range arts {
		ra := crv1alpha1.Artifact{}
		for k, v := range a.KeyValue {
			rv, err := render(v, tp)
			if err != nil {
				return nil, err
			}
			if ra.KeyValue == nil {
				ra.KeyValue = make(map[string]string, len(a.KeyValue))
			}
			ra.KeyValue[k] = rv
		}
		rarts[name] = ra
	}
	return rarts, nil
}

func render(arg string, tp TemplateParams) (string, error) {
	t, err := template.New("config").Funcs(sprig.TxtFuncMap()).Parse(arg)
	if err != nil {
		return "", errors.WithStack(err)
	}
	buf := bytes.NewBuffer(nil)
	if err = t.Execute(buf, tp); err != nil {
		return "", errors.WithStack(err)
	}
	return buf.String(), nil
}
