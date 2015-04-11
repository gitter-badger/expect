package expect

import "testing"

// TODO(Ariel): Create mock that implement TB interface
// and stub `Error` and `Fatal`

func TestStartWith(t *testing.T) {
	expect := New(t)
	expect("foo").To.StartWith("f")
	expect("foo").Not.To.StartWith("bar")
}

func TestEndWith(t *testing.T) {
	expect := New(t)
	expect("bar").To.EndWith("ar")
	expect("bar").Not.To.EndWith("az")
}

func TestContains(t *testing.T) {
	expect := New(t)
	expect("foobar").To.Contains("ba")
	expect("foobar").Not.To.Contains("ga")
}

func TestMatch(t *testing.T) {
	expect := New(t)
	expect("Foo").To.Match("(?i)foo")
}

func TestToChaining(t *testing.T) {
	expect := New(t)
	expect("foobarbaz").To.StartWith("foo").And.EndWith("baz").And.Contains("bar")
	expect("foo").Not.To.StartWith("bar").And.EndWith("baz").And.Contains("bob")
	expect("foo").To.Match("f").And.Match("(?i)F")
}