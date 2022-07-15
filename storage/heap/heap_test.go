package heap

import "testing"

func TestExists(t *testing.T) {
	var h HeapInstance
	h.Setup()
	h.IDs[52345234312] = true
	var id int64 = 52345234312
	want := true
	res := h.Exists(id)
	if !want == res {
		t.Fatalf("got %v, want %v", res, want)
	}
}

func TestNotExists(t *testing.T) {
	var h HeapInstance
	h.Setup()
	h.IDs[52345234312] = true
	var id int64 = 4125346324
	want := false
	res := h.Exists(id)
	if !want == res {
		t.Fatalf("got %v, want %v", res, want)
	}
}

func TestGet(t *testing.T) {
	var h HeapInstance
	h.Setup()
	h.ShortyToUrl["G1rW1b_wva"] = "https://www.youtube.com/watch?v=8YJxWV5yPx0"
	wantRes := "https://www.youtube.com/watch?v=8YJxWV5yPx0"
	res, err := h.Get("G1rW1b_wva")
	if err != nil {
		t.Fatalf("unexpected error uccured: %s", err.Error())
	}
	if res != wantRes {
		t.Fatalf("wrong shortUrl: got %s, want %s", res, wantRes)
	}
}

func TestGetNotFound(t *testing.T) {
	var h HeapInstance
	h.Setup()
	h.ShortyToUrl["G1rW1b_wva"] = "https://www.youtube.com/watch?v=8YJxWV5yPx0"
	wantRes, wantErrMsg := "", "shorty not found"
	res, err := h.Get("27I5tQ6N4E")
	if err == nil {
		t.Fatalf("error is %s, want nil", err.Error())
	}
	if err.Error() != "shorty not found" {
		t.Fatalf("wrong error: got %s, want %s", err.Error(), wantErrMsg)
	}
	if res != wantRes {
		t.Fatalf("wrong url: got %s, want %s", res, wantRes)
	}
}

func TestSave(t *testing.T) {
	var h HeapInstance
	h.Setup()
	url := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	short, resErr := h.Save(url)
	if resErr != nil {
		t.Fatal(resErr)
	}
	getResult, getErr := h.Get(short)
	if getErr != nil {
		t.Fatalf("shortUrl have to exists but it doesn't")
	}
	if getResult != url {
		t.Fatalf("url saved wrong: got %s, want %s", getResult, url)
	}
}
