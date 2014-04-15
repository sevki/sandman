package sandman

import "testing"

func TestHighlight_KNR(t *testing.T) {

	knr := `#include <stdio.h>
#include <stdlib.h>

int main(void)
{
    printf("Hello world!\n");
    return EXIT_SUCCESS;
}`

	v := Highlight(knr, "c")
	if v == "" {
		t.Fail()
	}

}

func TestHighlight_UTF(t *testing.T) {
	goc := `// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}`

	v := Highlight(goc, "go")
	if v == "" {
		t.Fail()
	}

}
