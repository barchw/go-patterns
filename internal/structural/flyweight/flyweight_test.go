package flyweight

import "testing"

func TestGlyphSharing(t *testing.T) {
	factory := NewGlyphFactory()

	g1 := factory.GetGlyph('A', "Arial")
	g2 := factory.GetGlyph('A', "Arial")

	if g1 != g2 {
		t.Fatal("same char+font should return the same glyph pointer")
	}
}

func TestDifferentGlyphsNotShared(t *testing.T) {
	factory := NewGlyphFactory()

	g1 := factory.GetGlyph('A', "Arial")
	g2 := factory.GetGlyph('B', "Arial")
	g3 := factory.GetGlyph('A', "Times")

	if g1 == g2 {
		t.Fatal("different chars should not share a glyph")
	}
	if g1 == g3 {
		t.Fatal("different fonts should not share a glyph")
	}
}

func TestGlyphCountReflectsSharing(t *testing.T) {
	factory := NewGlyphFactory()

	factory.GetGlyph('H', "Arial")
	factory.GetGlyph('e', "Arial")
	factory.GetGlyph('l', "Arial")
	factory.GetGlyph('l', "Arial")
	factory.GetGlyph('o', "Arial")

	if factory.GlyphCount() != 4 {
		t.Fatalf("expected 4 unique glyphs (H,e,l,o), got %d", factory.GlyphCount())
	}
}

func TestCharacterRenderer(t *testing.T) {
	factory := NewGlyphFactory()
	renderer := NewCharacterRenderer(factory)

	result := renderer.Render('A', "Arial", 10, 20)
	if result != "'A' [Arial] at (10,20)" {
		t.Fatalf("unexpected render output: %s", result)
	}

	renderer.Render('A', "Arial", 30, 40)
	if renderer.CharacterCount() != 2 {
		t.Fatalf("expected 2 rendered characters, got %d", renderer.CharacterCount())
	}
	if factory.GlyphCount() != 1 {
		t.Fatalf("expected 1 unique glyph, got %d", factory.GlyphCount())
	}
}
