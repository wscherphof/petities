package secure

import (
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

func getCookie(r *http.Request) (session *sessions.Session) {
	session, _ = store.Get(r, tokenName)
	return
}

func create(w http.ResponseWriter, r *http.Request, record interface{}, redirect bool) (err error) {
	session := getCookie(r)
	if session.Values[createdField] == nil {
		session.Values[createdField] = time.Now()
	}
	session.Values[recordField] = record
	session.Values[validatedField] = time.Now()
	if r.TLS == nil {
		err = ErrNoTLS
	} else if e := session.Save(r, w); e != nil {
		err = ErrTokenNotSaved
	} else if redirect {
		path := session.Values[returnField]
		if path == nil {
			path = config.LogOutPath
		}
		http.Redirect(w, r, path.(string), http.StatusSeeOther)
	}
	return
}

// LogIn creates the cookie and sets the cookie. It redirects back to the path
// where Authenticate() was called.
//
// 'record' is the authentication data to store in the cookie, as returned by
// Authentication()
func LogIn(w http.ResponseWriter, r *http.Request, record interface{}) (err error) {
	return create(w, r, record, true)
}

// Update updates the authentication data in the cookie.
func Update(w http.ResponseWriter, r *http.Request, record interface{}) (err error) {
	return create(w, r, record, false)
}

func sessionCurrent(session *sessions.Session) (current bool) {
	if created := session.Values[createdField]; created == nil {
	} else {
		current = time.Since(created.(time.Time)) < config.CookieTimeOut
	}
	return
}

func accountCurrent(session *sessions.Session, w http.ResponseWriter, r *http.Request) (current bool) {
	if validated := session.Values[validatedField]; validated == nil {
	} else if cur := time.Since(validated.(time.Time)) < config.SyncInterval; cur {
		current = true
	} else if record, cur := validate(session.Values[recordField]); cur {
		session.Values[recordField] = record
		session.Values[validatedField] = time.Now()
		_ = session.Save(r, w)
		current = true
	}
	return
}

type contextKey int

const authKey contextKey = 0

func authenticate(w http.ResponseWriter, r *http.Request, optional ...bool) (authenticated bool) {
	enforce := true
	if len(optional) > 0 {
		enforce = !optional[0]
	}
	session := getCookie(r)
	if !session.IsNew && sessionCurrent(session) && accountCurrent(session, w, r) {
		context.Set(r, authKey, session.Values[recordField])
		authenticated = true
	} else if enforce {
		session = clearCookie(r)
		session.Values[returnField] = r.URL.Path
		_ = session.Save(r, w)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`<!DOCTYPE html>
			<html>
				<head>
					<meta charset="utf-8">
					<meta http-equiv="refresh" content="0; url=` + config.LogInPath + `">
				</head>
				<body>
					<h2>Forbidden</h2>
					<a id="location" href="` + config.LogInPath + `">Log in</a>
				</body>
			</html>
		`))
	}
	return
}

/*
Authentication returns the record that was stored in the cookie on LogIn().

Call from a Handle wrapped in secure.Handle or secure.IfHandle.
*/
func Authentication(r *http.Request) interface{} {
	return context.Get(r, authKey)
}

func clearCookie(r *http.Request) (session *sessions.Session) {
	session = getCookie(r)
	delete(session.Values, recordField)
	delete(session.Values, createdField)
	delete(session.Values, validatedField)
	return
}

// LogOut deletes the cookie. If 'redirect' is 'true', the request is redirected
// to config.LogOutPath.
func LogOut(w http.ResponseWriter, r *http.Request, redirect bool) {
	session := clearCookie(r)
	session.Options = &sessions.Options{
		MaxAge: -1,
	}
	_ = session.Save(r, w)
	if redirect {
		http.Redirect(w, r, config.LogOutPath, http.StatusSeeOther)
	}
}
