package gui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Page struct {
	elems     map[string]*E
	htmlCache []byte      // static html content, rendered only once
	haveJS    bool        // have called JS()?
	data      interface{} // any additional data to be passed to template
	onUpdate  func()
}

func NewPage(htmlTemplate string, data interface{}) *Page {
	d := &Page{elems: make(map[string]*E),
		data: data}

	// exec template (once)
	t := template.Must(template.New("").Parse(htmlTemplate))
	cache := bytes.NewBuffer(nil)
	check(t.Execute(cache, d))
	d.htmlCache = cache.Bytes()

	// check if template contains {{.JS}}
	if !d.haveJS {
		log.Panic("template should call {{.JS}}")
	}
	return d
}

func (d *Page) Set(id string, v interface{}) {
	d.elem(id).set(v)
}

// Set func to be executed each time javascript polls for updates
func (d *Page) OnUpdate(f func()) {
	d.onUpdate = f
}

// {{.JS}} should always be embedded in the template <head>.
// Expands to needed JavaScript code.
func (d *Page) JS() string {
	d.haveJS = true
	return JS
}

// {{.ErrorBox}} should be embedded in the template where errors are to be shown.
// CSS rules for class ErrorBox may be set, e.g., to render errors in red.
func (t *Page) ErrorBox() string {
	return `<span id=ErrorBox class=ErrorBox></span> <span id=MsgBox class=ErrorBox></span>`
}

// {{.UpdateButton}} adds a page Update button
func (t *Page) UpdateButton() string {
	return `<button onclick="update();"> &#x21bb; </button>`
}

// {{.UpdateBox}} adds an auto update checkbox
func (t *Page) UpdateBox() string {
	return `<input type=checkbox id=UpdateBox class=CheckBox checked=true onchange="autoUpdate=elementById('UpdateBox').checked").checked">auto update</input>`
}

// {{.Data}} returns the extra data that was passed to NewPage
func (t *Page) Data() interface{} {
	return t.data
}

// return elem[id], panic if non-existent
func (d *Page) elem(id string) *E {
	if e, ok := d.elems[id]; ok {
		return e
	} else {
		panic("no element with id: " + id)
	}
}

// elem[id] = e, panic if already defined
func (d *Page) addElem(id string, e El) {
	if _, ok := d.elems[id]; ok {
		panic("addElem: already defined: " + id)
	} else {
		d.elems[id] = &E{el: e, dirty: true}
	}
}

// ServeHTTP implements http.Handler.
func (d *Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(w, "not allowed: "+r.Method+" "+r.URL.Path, http.StatusForbidden)
	case "GET":
		d.serveContent(w, r)
	case "POST":
		d.serveUpdate(w, r)
	case "PUT":
		d.serveEvent(w, r)
	}
}

// serves the html content.
func (d *Page) serveContent(w http.ResponseWriter, r *http.Request) {
	for _, e := range d.elems {
		e.dirty = true
	}
	w.Write(d.htmlCache)
}

// HTTP handler for event notifications by button clicks etc
func (d *Page) serveEvent(w http.ResponseWriter, r *http.Request) {
	var ev event
	check(json.NewDecoder(r.Body).Decode(&ev))
	fmt.Println(ev)
	//el := d.elem(ev.ID)
	//el.setValue(ev.Arg)
	//if el.onevent != nil {
	//	el.onevent()
	//}
}

type event struct {
	ID  string
	Arg interface{}
}

// HTTP handler for updating the dynamic elements
func (d *Page) serveUpdate(w http.ResponseWriter, r *http.Request) {
	d.onUpdate()

	buf := make([]byte, 1024)
	r.Body.Read(buf)
	fmt.Println(string(buf))

	calls := make([]jsCall, 0, len(d.elems))
	for id, e := range d.elems {
		if e.dirty {
			calls = append(calls, e.el.update(id))
			e.dirty = false
		}
	}
	check(json.NewEncoder(w).Encode(calls))
}

// javascript call
type jsCall struct {
	F    string        // function to call
	Args []interface{} // function arguments
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
