package csptester

import (
	"html/template"
	"net/http"
)

var (
	testHTML = template.Must(template.New("test.html").Parse(`
<!doctype html>
<head>
	<meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
{{if .UseMeta}}
    <meta http-equiv="X-Content-Security-Policy" content="{{.CSPString}}">
	<meta http-equiv="Content-Security-Policy" content="{{.CSPString}}">
{{end}}
</head>
<body>
    <p>hello world! i'm safe from injected code!</p>
    <script>
        setTimeout(function() {
            document.querySelector('p').innerHTML = 'ahaha jk :)'
        }, 1000)
	</script>
	
	<hr>
	<pre>Content-Security-Policy: {{.CSPString}}</pre>
</body>
	`))
)

type data struct {
	UseMeta   bool
	CSPString string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	useMeta := r.URL.Query().Get("meta") == "1"

	cspString := r.URL.Query().Get("csp")
	if cspString == "" {
		cspString = "script-src 'none'"
	}

	if !useMeta {
		w.Header().Add("X-Content-Security-Policy", cspString)
		w.Header().Add("Content-Security-Policy", cspString)
	}

	testHTML.Execute(w, data{UseMeta: useMeta, CSPString: cspString})
}
