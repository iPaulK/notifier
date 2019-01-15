package notifier

import (
	"bytes"
	"text/template"
)

var (
	mailLayout = template.Must(template.New("mail").Parse(`
<html>
<head>
<title></title>
<style type="text/css">
body {
background: #fff; margin-top: 30px; padding: 0;
}
</style>
</head>
<body style="background: #fff; margin-top: 30px; padding: 0;">
{{ .Body }}
</body>
</html>
`))
)

type Layout struct {
	Body string
}

// NewLayout creates new layout
func NewLayout() *Layout {
	return &Layout{}
}

func (l *Layout) SetBody(body string) *Layout {
	l.Body = body
	return l
}

func (l *Layout) GetBody() string {
	return l.Body
}

func (l *Layout) Render() (string, error) {
	var buf bytes.Buffer

	if err := mailLayout.Execute(&buf, l); err != nil {
		return "", err
	}
	return buf.String(), nil
}
