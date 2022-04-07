package automaticshit

import (
	"testing"
)

func TestAutomaticShit(t *testing.T) {
	people := []string{"test1", "test2", "test3", "test4", "test5"}
	shit := AutomaticShit(people, 2, 2, 31)
	t.Log(shit)
	t.Log(len(shit))
}
