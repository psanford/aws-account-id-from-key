package awskey

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAwskey(t *testing.T) {
	k, err := Decode("ASIAQNZGKIQY56JQ7WML")
	if err != nil {
		t.Fatal(err)
	}

	expect := Key{
		Type:      STSKey,
		AccountID: "029608264753",
	}
	if !cmp.Equal(k, &expect) {
		t.Fatal(cmp.Diff(k, expect))
	}

}
