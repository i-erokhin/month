package month

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMonthJSON(t *testing.T) {
	t.Parallel()

	p := NewCurrent()
	b, err := p.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	jsonExpected := fmt.Sprintf("%q", p)
	jsonMarshaled := string(b)
	if jsonMarshaled != jsonExpected {
		t.Fatalf("JSON is %q, but %q exected.", jsonMarshaled, jsonExpected)
	}

	var p2 Month
	err = json.Unmarshal(b, &p2)
	if err != nil {
		t.Log(string(b))
		t.Fatal(err)
	}

	if p2 != *p {
		t.Errorf("Month is %q, but %q expected.", p2, *p)
	}
}
