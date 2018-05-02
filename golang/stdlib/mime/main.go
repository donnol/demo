package main

import (
	"log"
	"mime"
)

func main() {
	err := mime.AddExtensionType(".jds", "jds")
	if err != nil {
		log.Fatal(err) // mime: extension "jds" missing leading dot
	}

	for _, typ := range []string{
		"Jds",  // [.jds]
		"jds",  // [.jds]
		"hml",  // []
		"html", // []
	} {
		exts, err := mime.ExtensionsByType(typ)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(exts)
	}

	for _, mt := range []struct {
		t     string
		param map[string]string
	}{
		{"", make(map[string]string)},
		{"jds", make(map[string]string)},
		{"jds", map[string]string{"jds": "jds"}},
		{"Application/json", map[string]string{"MaxAge": "86400"}},
	} {
		v := mime.FormatMediaType(mt.t, mt.param)
		log.Println(v)

		vmt, vparam, err := mime.ParseMediaType(v)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(vmt, vparam)
	}

	typ := mime.TypeByExtension(".jds")
	log.Println(typ)

	// encData := mime.QEncoding.Encode("utf-8", "Hello, I am jd") // mime: invalid RFC 2047 encoded-word
	encData := mime.QEncoding.Encode("utf-8", "¡Hola, señor!")
	log.Println(encData)

	dec := new(mime.WordDecoder)
	decData, err := dec.Decode(encData)
	if err != nil {
		log.Fatal(err) // mime: invalid RFC 2047 encoded-word
	}
	log.Println(decData)
}
