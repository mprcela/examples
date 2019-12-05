package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

var heads  = map[int]string{
	0: "",
	1: `
<meta name="apple-mobile-web-app-capable" content="yes" />
`,
	2: `
<meta name="apple-mobile-web-app-capable" content="yes" />
<script type="text/javascript">
(function(document,navigator,standalone) {
    // prevents links from apps from oppening in mobile safari
    // this javascript must be the first script in your <head>
    if ((standalone in navigator) && navigator[standalone]) {
        var curnode, location=document.location, stop=/^(a|html)$/i;
        document.addEventListener('click', function(e) {
            curnode=e.target;
            while (!(stop).test(curnode.nodeName)) {
                curnode=curnode.parentNode;
            }
            // Condidions to do this only on links to your own app
            // if you want all links, use if('href' in curnode) instead.
            if('href' in curnode && ( curnode.href.indexOf('http') || ~curnode.href.indexOf(location.host) ) ) {
                e.preventDefault();
                location.href = curnode.href;
            }
        },false);
    }
})(document,window.navigator,'standalone');
</script>
`,
	3: `
<meta name="apple-mobile-web-app-capable" content="yes" />
<link rel="manifest" href="/home.webmanifest">
<script type="text/javascript">
(function(document,navigator,standalone) {
    // prevents links from apps from oppening in mobile safari
    // this javascript must be the first script in your <head>
    if ((standalone in navigator) && navigator[standalone]) {
        var curnode, location=document.location, stop=/^(a|html)$/i;
        document.addEventListener('click', function(e) {
            curnode=e.target;
            while (!(stop).test(curnode.nodeName)) {
                curnode=curnode.parentNode;
            }
            // Condidions to do this only on links to your own app
            // if you want all links, use if('href' in curnode) instead.
            if('href' in curnode && ( curnode.href.indexOf('http') || ~curnode.href.indexOf(location.host) ) ) {
                e.preventDefault();
                location.href = curnode.href;
            }
        },false);
    }
})(document,window.navigator,'standalone');
</script>
`,
4:`
<link rel="manifest" href="/home.webmanifest">
<script type="text/javascript">
(function(document,navigator,standalone) {
    // prevents links from apps from oppening in mobile safari
    // this javascript must be the first script in your <head>
    if ((standalone in navigator) && navigator[standalone]) {
        var curnode, location=document.location, stop=/^(a|html)$/i;
        document.addEventListener('click', function(e) {
            curnode=e.target;
            while (!(stop).test(curnode.nodeName)) {
                curnode=curnode.parentNode;
            }
            // Condidions to do this only on links to your own app
            // if you want all links, use if('href' in curnode) instead.
            if('href' in curnode && ( curnode.href.indexOf('http') || ~curnode.href.indexOf(location.host) ) ) {
                e.preventDefault();
                location.href = curnode.href;
            }
        },false);
    }
})(document,window.navigator,'standalone');
</script>
`,
}

var document = `
<html>
	<head>
		<title>Hello world</title>
    %s 
	</head>
  <body style="font-size:4em">
		<p><a href="/">Go home</a></p>
		<p><a href="/bar">Go to the bar</a></p>
		<p><a href="/foo">Go to the foo</a></p>
		<hr/>
    <h4>Path:</h4>
		<p>%s</p>
    <h4>Head:</h4>
		<p>%s</p>
  </body>
</html>
`;

func writeHeaders(w http.ResponseWriter){
	w.Header().Add("Cache-Control", "no-cache, no-store")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("Expires", "Fri, 01 Jan 1990 00:00:00 GMT")
	w.Header().Add("max-age", "0")
}

var currentHead = 4;

func main() {
	http.HandleFunc("/home.webmanifest", func(w http.ResponseWriter, r *http.Request) {
		writeHeaders(w)
		w.Header().Add("Content-Type","application/manifest+json")
		fmt.Fprintf(w, `{"name": "Hello world", "display": "standalone"}`)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeHeaders(w)
		head := heads[currentHead];
		rsp := fmt.Sprintf(document, head, html.EscapeString(r.URL.Path), html.EscapeString(head))
		fmt.Fprintf(w, rsp)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
