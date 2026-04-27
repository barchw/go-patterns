/*
Package flyweight implements the Flyweight design pattern.

What is it?
Flyweight lets you fit more objects into available memory by sharing common parts
of state between multiple objects. The flyweight factory checks whether an object
with the given parameters already exists, and if so -- returns the existing instance. In this
package, the pattern models character rendering -- glyphs (Glyph) are shared,
and the position (x, y) is stored separately in RenderedCharacter.

When to use?
  - When the program must handle a huge number of similar objects that consume a lot of memory.
  - When a large part of the state can be shared (intrinsic state), and the rest depends on context.
  - When creating new objects is expensive, and many of them are identical.
  - When you want to reduce memory usage without changing business logic.

When NOT to use?
  - When objects have little shared state and sharing won't bring savings.
  - When there are few objects and memory usage isn't a concern.
  - When it's hard to separate state into intrinsic (shared) and extrinsic (unique).

Tips and pitfalls:
  - Flyweights MUST be immutable -- changing a shared Glyph will affect all objects.
  - The current GlyphFactory implementation is not thread-safe; in a concurrent program
    use sync.Mutex or sync.Map.
  - Flyweight shares parts of an object's state; Cache stores full computation results;
    Pool manages reuse of resources -- these are different concepts.
*/
package flyweight

import "fmt"

// Glyph stores the intrinsic (shared) state of a flyweight -- the character and font family.
type Glyph struct {
	Char       rune
	FontFamily string
}

// GlyphFactory manages a pool of flyweights, creating new glyphs or returning existing ones.
type GlyphFactory struct {
	glyphs map[string]*Glyph
}

// NewGlyphFactory creates a new flyweight factory with an empty glyph map.
func NewGlyphFactory() *GlyphFactory {
	return &GlyphFactory{glyphs: make(map[string]*Glyph)}
}

// GetGlyph returns an existing glyph or creates a new one if it doesn't exist in the pool.
func (f *GlyphFactory) GetGlyph(char rune, fontFamily string) *Glyph {
	key := fmt.Sprintf("%c:%s", char, fontFamily)
	if g, ok := f.glyphs[key]; ok {
		return g
	}
	g := &Glyph{Char: char, FontFamily: fontFamily}
	f.glyphs[key] = g
	return g
}

// GlyphCount returns the number of unique glyphs stored in the factory.
func (f *GlyphFactory) GlyphCount() int {
	return len(f.glyphs)
}

// RenderedCharacter stores the extrinsic (unique) state -- the character's position on screen
// and a pointer to the shared Glyph.
type RenderedCharacter struct {
	Glyph *Glyph
	X     int
	Y     int
}

// CharacterRenderer combines the flyweight factory with extrinsic state, rendering characters on screen.
type CharacterRenderer struct {
	factory    *GlyphFactory
	characters []RenderedCharacter
}

// NewCharacterRenderer creates a new character renderer using the given flyweight factory.
func NewCharacterRenderer(factory *GlyphFactory) *CharacterRenderer {
	return &CharacterRenderer{factory: factory}
}

// Render retrieves (or creates) a glyph from the factory and assigns it a position on screen.
func (r *CharacterRenderer) Render(char rune, fontFamily string, x, y int) string {
	glyph := r.factory.GetGlyph(char, fontFamily)
	r.characters = append(r.characters, RenderedCharacter{Glyph: glyph, X: x, Y: y})
	return fmt.Sprintf("'%c' [%s] at (%d,%d)", glyph.Char, glyph.FontFamily, x, y)
}

// CharacterCount returns the total number of rendered characters.
func (r *CharacterRenderer) CharacterCount() int {
	return len(r.characters)
}
