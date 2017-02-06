package lib

import(
	"strings"
	"testing"
)
func TestIndex(t *testing.T){
	const s,sep,want ="chicken","ken",3
	got:=strings.Index(s,sep)
	if got !=want{
		t.Errorf("Index(%q,%q)=%v",s,sep,got,want)
	}
}