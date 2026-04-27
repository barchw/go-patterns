package flyweight

import "fmt"

type Glyph struct {
	Char       rune
	FontFamily string
}

type GlyphFactory struct {
	glyphs map[string]*Glyph
}

func NewGlyphFactory() *GlyphFactory {
	return &GlyphFactory{glyphs: make(map[string]*Glyph)}
}

func (f *GlyphFactory) GetGlyph(char rune, fontFamily string) *Glyph {
	key := fmt.Sprintf("%c:%s", char, fontFamily)
	if g, ok := f.glyphs[key]; ok {
		return g
	}
	g := &Glyph{Char: char, FontFamily: fontFamily}
	f.glyphs[key] = g
	return g
}

func (f *GlyphFactory) GlyphCount() int {
	return len(f.glyphs)
}

type RenderedCharacter struct {
	Glyph *Glyph
	X     int
	Y     int
}

type CharacterRenderer struct {
	factory    *GlyphFactory
	characters []RenderedCharacter
}

func NewCharacterRenderer(factory *GlyphFactory) *CharacterRenderer {
	return &CharacterRenderer{factory: factory}
}

func (r *CharacterRenderer) Render(char rune, fontFamily string, x, y int) string {
	glyph := r.factory.GetGlyph(char, fontFamily)
	r.characters = append(r.characters, RenderedCharacter{Glyph: glyph, X: x, Y: y})
	return fmt.Sprintf("'%c' [%s] at (%d,%d)", glyph.Char, glyph.FontFamily, x, y)
}

func (r *CharacterRenderer) CharacterCount() int {
	return len(r.characters)
}
