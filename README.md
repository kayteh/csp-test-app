# CSP Test App

generates a CSP header or meta tag and tries to do something illegal.

https://csp-test-app.katacat.now.sh/

accepts query strings:
- `?csp=` uses some CSP string
- `?meta=1` will inject meta tags instead of headers

if you see `"ahaha jk :)"`, CSP allowed an inline script.