package tree_sitter_markdown_inline_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-markdown_inline"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_markdown_inline.Language())
	if language == nil {
		t.Errorf("Error loading MarkdownInline grammar")
	}
}
