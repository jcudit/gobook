package my_font

import (
	"errors"
	"fmt"
	"log"
)

type Font struct {
	family string
	size   int
}

func New(family string, size int) *Font {
	newFont := Font{family, size}
	return &newFont
}

func (f *Font) Familee() string {
	return f.family
}

func (f *Font) Family() string {
	return f.family
}

func (f *Font) Size() int {
	return f.size
}

func (f *Font) SetFamily(family string) error {
	if family != "" {
		f.family = family
	} else {
		msg := "Unacceptable font family name provided"
		log.Println(msg)
		return errors.New(msg)
	}
	return nil
}

func (f *Font) SetSize(size int) error {
	if size >= 5 && size <= 144 {
		f.size = size
	} else {
		msg := "Unacceptable font size provided"
		log.Println(msg)
		return errors.New(msg)
	}
	return nil
}

func (f *Font) String() string {
	return fmt.Sprintf("{font-family: \"%s\"; font-size: %dpt;}", f.family, f.size)
}
