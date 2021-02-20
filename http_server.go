package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func http_server_fallback(w http.ResponseWriter, req *http.Request) {
	switch strings.ToLower(req.Host) {
	case "captive.apple.com":
		hotspot_detect_html(w, req)
		return

	case "capnet.elementary.io":
		capnet(w, req)
		return

	case "www.archlinux.org":
		// mock the redirect
		w.Header().Add("Location", "https://archlinux.org"+req.RequestURI)
		w.WriteHeader(http.StatusMovedPermanently)
		fmt.Fprint(w, "<html>\n<head><title>301 Moved Permanently</title></head>\n<body>\n<center><h1>301 Moved Permanently</h1></center>\n<hr><center>nginx</center>\n</body>\n</html>\n")
		return

	default:
		// fallback to a 404 page
		log.Printf("[HTTP] not implemented: %s %s => \"%s%s\"\n", req.Method, req.RemoteAddr, req.Host, req.RequestURI)
		w.WriteHeader(http.StatusNotFound)
	}
}

// /robots.txt
func robots_txt(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User-agent: *\nDisallow: /\n")
}

// http://www.msftncsi.com/ncsi.txt
func ncsi_txt(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Microsoft NCSI")
}

// http://www.msftconnecttest.com/connecttest.txt
// http://ipv6.msftconnecttest.com/connecttest.txt
func connecttest(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Microsoft Connect Test")
}

// http://www.msftconnecttest.com/redirect
func redirect(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Location", "http://go.microsoft.com/fwlink/?LinkID=219472&clcid=0x409")
	w.Header().Add("Server", "Microsoft-IIS/114.514")
	w.Header().Add("Content-Length", "0")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte{})
}

// http://captive.apple.com/hotspot-detect.html
func hotspot_detect_html(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML>")
}

// http://clients3.google.com/generate_204
// http://connectivitycheck.gstatic.com/generate_204
// http://connectivitycheck.android.com/generate_204
// http://connect.rom.miui.com/generate_204
// http://connectivitycheck.platform.hicloud.com/generate_204
func generate_204(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Length", "0")
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte{})
}

// http://network-test.debian.org/nm
// http://nmcheck.gnome.org/check_network_status.txt
// http://www.archlinux.org/check_network_status.txt
func nm(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("X-NetworkManager-Status", "online")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "NetworkManager is online\n")
}

// http://detectportal.firefox.com/success.txt
func success_txt(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "success\n")
}

// http://start.ubuntu.com/connectivity-check.html
func connectivity_check_html(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<!DOCTYPE HTML PUBLIC \"-//W3C//DTD HTML 3.2//EN\">\n<HTML>\n<HEAD>\n<TITLE>Lorem Ipsum</TITLE>\n</HEAD>\n<BODY>\n<P>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</P>\n</BODY>\n</HTML>\n")
}

// http://capnet.elementary.io
// warning: relative resources exist
func capnet(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "<!doctype html>\n<!--[if IE]><html lang=\"en\" class=\"ie-legacy\"><![endif]-->\n<!--[if !IE]><!--><html lang=\"en\"><!--<![endif]-->\n<head>\n<meta charset=\"UTF-8\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0, minimum-scale=1.0\">\n<meta name=\"description\" content=\"The fast, open, and privacy-respecting replacement for Windows and macOS\">\n<meta name=\"author\" content=\"elementary, Inc.\">\n<meta name=\"theme-color\" content=\"#3689e6\">\n<meta name=\"twitter:card\" content=\"summary_large_image\">\n<meta name=\"twitter:site\" content=\"@elementary\">\n<meta name=\"twitter:creator\" content=\"@elementary\">\n<meta property=\"og:title\" content=\"You're connected! &sdot; elementary\" />\n<meta property=\"og:description\" content=\"The fast, open, and privacy-respecting replacement for Windows and macOS\" />\n<meta property=\"og:image\" content=\"https://elementary.io/images/preview.png\" />\n<meta itemprop=\"name\" content=\"You're connected! &sdot; elementary\" />\n<meta itemprop=\"description\" content=\"The fast, open, and privacy-respecting replacement for Windows and macOS\" />\n<meta itemprop=\"image\" content=\"https://elementary.io/images/preview.png\" />\n<meta name=\"apple-mobile-web-app-title\" content=\"elementary\">\n<link rel=\"manifest\" href=\"/manifest.json\">\n<title>You're connected! &sdot; elementary</title>\n<base href=\"/\">\n<link rel=\"shortcut icon\" href=\"favicon.ico\">\n<link rel=\"apple-touch-icon\" href=\"images/launcher-icons/apple-touch-icon.png\">\n<link rel=\"icon\" type=\"image/png\" href=\"images/favicon.png\" sizes=\"256x256\">\n<link rel=\"alternate\" type=\"text/html\" hreflang=\"en\" href=\"/capnet-assist\">\n<link rel=\"stylesheet\" href=\"https://pro.fontawesome.com/releases/v5.9.0/css/all.css\" integrity=\"sha384-vlOMx0hKjUCl4WzuhIhSNZSm2yQCaf0mOU1hEDK/iztH3gU4v5NMmJln9273A6Jz\" crossorigin=\"anonymous\">\n<link rel=\"stylesheet\" type=\"text/css\" media=\"all\" href=\"styles/main.css\">\n<link rel=\"stylesheet\" type=\"text/css\" media=\"all\" href=\"styles/capnet-assist.css\">\n</head>\n<body class=\"page-capnet-assist\">\n<nav>\n<div class=\"nav-content\">\n<ul>\n<li><a href=\"/\" class=\"logomark\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"22\" width=\"22\"><path d=\"M11 0C4.926 0 0 4.926 0 11s4.926 11 11 11 11-4.926 11-11S17.074 0 11 0zm-.02 1.049h.002a.07.07 0 00.018 0c4.548 0 8.384 3.052 9.57 7.217a17.315 17.315 0 01-4.213 5.496c-.847.733-1.773 1.383-2.79 1.842-1.018.458-2.136.719-3.247.656a5.6 5.6 0 01-2.271-.633 18.276 18.276 0 005.078-4.12c.956-1.116 1.791-2.39 2.115-3.833.165-.722.19-1.48.04-2.207a4.079 4.079 0 00-.999-1.965 4.013 4.013 0 00-1.855-1.098 4.85 4.85 0 00-2.153-.078c-1.423.261-2.693 1.086-3.705 2.108-1.787 1.8-2.89 4.34-2.687 6.875.102 1.267.524 2.511 1.252 3.554.143.205.302.398.467.584a16.228 16.228 0 01-3.086.739A9.946 9.946 0 0110.98 1.049zm.07 2.02c.004-.001.016.002.012.001.692 0 1.387.211 1.936.627a3.2 3.2 0 011.068 1.49 4.12 4.12 0 01.192 1.848c-.143 1.243-.77 2.388-1.533 3.393-1.354 1.778-3.156 3.203-5.159 4.2-.19.095-.386.187-.582.274a5.114 5.114 0 01-1.05-1.295c-.59-1.044-.784-2.284-.67-3.482.116-1.199.526-2.356 1.082-3.43.643-1.244 1.516-2.418 2.732-3.09a4.14 4.14 0 011.973-.537zm9.83 6.81c.042.367.071.739.071 1.117 0 5.497-4.452 9.953-9.95 9.953a9.917 9.917 0 01-7.59-3.53 18.138 18.138 0 003.17-1.06c.461.346.967.634 1.507.84 1.59.61 3.392.52 4.996-.035 1.603-.555 3.021-1.549 4.256-2.705a18.264 18.264 0 003.54-4.58z\" /></svg></a></li>\n<li><a href=\"/support\">Support</a></li>\n<li><a href=\"https://developer.elementary.io\" target=\"_self\">Developer</a></li>\n<li><a href=\"/get-involved\">Get Involved</a></li>\n<li><a href=\"/store/\">Store</a></li>\n<li><a href=\"https://blog.elementary.io\" target=\"_self\">Blog</a></li>\n</ul>\n<ul class=\"right\">\n<li><a href=\"https://youtube.com/user/elementaryproject\" target=\"_blank\" rel=\"noopener\" data-l10n-off title=\"Youtube\"><i class=\"fab fa-youtube\"></i></a></li>\n<li><a href=\"https://www.facebook.com/elementaryos\" target=\"_blank\" rel=\"noopener\" data-l10n-off title=\"Facebook\"><i class=\"fab fa-facebook-f\"></i></a></li>\n<li><a href=\"https://mastodon.social/@elementary\" target=\"_blank\" rel=\"noopener\" data-l10n-off title=\"Mastodon\"><i class=\"fab fa-mastodon\"></i></a></li>\n<li><a href=\"https://www.reddit.com/r/elementaryos\" target=\"_blank\" rel=\"noopener\" data-l10n-off title=\"Reddit\"><i class=\"fab fa-reddit\"></i></a></li>\n<li><a href=\"https://elementaryos.stackexchange.com\" target=\"_blank\" rel=\"noopener\" data-l10n-off title=\"Stack Exchange\"><i class=\"fab fa-stack-exchange\"></i></a></li>\n<li><a href=\"https://twitter.com/elementary\" target=\"_blank\" rel=\"noopener\" data-l10n-off title=\"Twitter\"><i class=\"fab fa-twitter\"></i></a></li>\n<li><a href=\"https://community-slack.elementary.io/\" target=\"_blank\" rel=\"noopener\" data-l10n-off title=\"Slack\"><i class=\"fab fa-slack\"></i></a></li>\n</ul>\n</div>\n</nav>\n<div id=\"content-container\">\n<div class=\"row\">\n<h1>You're connected!</h1>\n<p>Your Internet connection appears to be working. You can safely close this window and continue using your device.</p>\n<h2>Why is this window appearing?</h2>\n<p>elementary OS automatically checks your Internet connection when you connect to a new Wi-Fi network. If it detects there is not an Internet connection (i.e. if you are connecting to a captive portal at a coffee shop or other public location), this window will appear and display the login page.</p>\n<p>Some networks can appear to be a captive portal at first, triggering this window, then actually end up connecting. In those cases, you'll see this message and can safely close the window.</p>\n</div>\n</div>\n<footer>\n<div>\n<p>\nCopyright &copy; 2021 elementary, Inc. </p>\n<div class=\"popover\">\n<a href=\"#\"><i class=\"far fa-language\"></i> Language</a>\n<div class=\"popover-content\">\n<strong>Change Site Language</strong>\n<ul>\n <li><a href=\"/en/\" rel=\"alternate\" hreflang=\"en\" data-l10n-off>\nEnglish </a></li>\n<hr>\n<li><a href=\"/af/\" rel=\"alternate\" hreflang=\"af\" data-l10n-off>\nAfrikaans </a></li>\n<li><a href=\"/ar/\" rel=\"alternate\" hreflang=\"ar\" data-l10n-off>\nالعَرَبِيَّة‎‎ </a></li>\n<li><a href=\"/ca/\" rel=\"alternate\" hreflang=\"ca\" data-l10n-off>\ncatalà </a></li>\n<li><a href=\"/cs_CZ/\" rel=\"alternate\" hreflang=\"cs-CZ\" data-l10n-off>\nčeština </a></li>\n<li><a href=\"/de/\" rel=\"alternate\" hreflang=\"de\" data-l10n-off>\nDeutsch </a></li>\n<li><a href=\"/es/\" rel=\"alternate\" hreflang=\"es\" data-l10n-off>\nEspañol </a></li>\n<li><a href=\"/fi/\" rel=\"alternate\" hreflang=\"fi\" data-l10n-off>\nFinnish </a></li>\n<li><a href=\"/fr/\" rel=\"alternate\" hreflang=\"fr\" data-l10n-off>\nFrançais </a></li>\n<li><a href=\"/he/\" rel=\"alternate\" hreflang=\"he\" data-l10n-off>\nעִברִית </a></li>\n<li><a href=\"/it/\" rel=\"alternate\" hreflang=\"it\" data-l10n-off>\nItaliano </a></li>\n<li><a href=\"/ja/\" rel=\"alternate\" hreflang=\"ja\" data-l10n-off>\n日本語 </a></li>\n<li><a href=\"/ko/\" rel=\"alternate\" hreflang=\"ko\" data-l10n-off>\n한국어 </a></li>\n<li><a href=\"/lt/\" rel=\"alternate\" hreflang=\"lt\" data-l10n-off>\nLietuvių kalba </a></li>\n<li><a href=\"/ms/\" rel=\"alternate\" hreflang=\"ms\" data-l10n-off>\nbahasa Melayu </a></li>\n<li><a href=\"/mr/\" rel=\"alternate\" hreflang=\"mr\" data-l10n-off>\nमराठी </a></li>\n<li><a href=\"/nb/\" rel=\"alternate\" hreflang=\"nb\" data-l10n-off>\nBokmål </a></li>\n<li><a href=\"/nl/\" rel=\"alternate\" hreflang=\"nl\" data-l10n-off>\nNederlands </a></li>\n<li><a href=\"/pl/\" rel=\"alternate\" hreflang=\"pl\" data-l10n-off>\nPolski </a></li>\n<li><a href=\"/pt_BR/\" rel=\"alternate\" hreflang=\"pt-BR\" data-l10n-off>\nPortuguês (Brasil) </a></li>\n<li><a href=\"/pt/\" rel=\"alternate\" hreflang=\"pt\" data-l10n-off>\nPortuguês (Portugal) </a></li>\n<li><a href=\"/ru/\" rel=\"alternate\" hreflang=\"ru\" data-l10n-off>\nРусский </a></li>\n<li><a href=\"/th/\" rel=\"alternate\" hreflang=\"th\" data-l10n-off>\nThai </a></li>\n<li><a href=\"/sk/\" rel=\"alternate\" hreflang=\"sk\" data-l10n-off>\nSlovak </a></li>\n<li><a href=\"/sv/\" rel=\"alternate\" hreflang=\"sv\" data-l10n-off>\nSwedish </a></li>\n<li><a href=\"/tr_TR/\" rel=\"alternate\" hreflang=\"tr-TR\" data-l10n-off>\nTürkçe </a></li>\n<li><a href=\"/uk/\" rel=\"alternate\" hreflang=\"uk\" data-l10n-off>\nукраїнська </a></li>\n<li><a href=\"/zh_CN/\" rel=\"alternate\" hreflang=\"zh-CN\" data-l10n-off>\n简体中文 </a></li>\n<li><a href=\"/zh_TW/\" rel=\"alternate\" hreflang=\"zh-TW\" data-l10n-off>\n繁體中文 </a></li>\n</ul>\n</div>\n</div>\n</div>\n<ul>\n<li><a href=\"/press\">Press</a></li>\n<li><a href=\"/brand\">Brand</a></li>\n<li><a href=\"/oem\">OEMs</a></li>\n<li><a href=\"/privacy\">Privacy</a></li>\n<li><a href=\"/team\">Team</a></li>\n<li><a href=\"/open-source\">Open Source</a></li>\n</ul>\n</footer>\n</body>\n</html>\n")
}
