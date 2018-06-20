package fyne

import "os"

import "testing"
import "github.com/stretchr/testify/assert"

var imgName = "dot.png"
var imgBytes = []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 1, 0, 0, 0, 1, 8, 2, 0, 0, 0, 144, 119, 83, 222, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 11, 19, 0, 0, 11, 19, 1, 0, 154, 156, 24, 0, 0, 0, 7, 116, 73, 77, 69, 7, 226, 6, 5, 15, 49, 19, 102, 121, 20, 234, 0, 0, 0, 29, 105, 84, 88, 116, 67, 111, 109, 109, 101, 110, 116, 0, 0, 0, 0, 0, 67, 114, 101, 97, 116, 101, 100, 32, 119, 105, 116, 104, 32, 71, 73, 77, 80, 100, 46, 101, 7, 0, 0, 0, 12, 73, 68, 65, 84, 8, 215, 99, 248, 207, 192, 0, 0, 3, 1, 1, 0, 24, 221, 141, 176, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}
var imgGo = "&fyne.Resource{\n\tName:    \"dot.png\",\n\tContent: []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 1, 0, 0, 0, 1, 8, 2, 0, 0, 0, 144, 119, 83, 222, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 11, 19, 0, 0, 11, 19, 1, 0, 154, 156, 24, 0, 0, 0, 7, 116, 73, 77, 69, 7, 226, 6, 5, 15, 49, 19, 102, 121, 20, 234, 0, 0, 0, 29, 105, 84, 88, 116, 67, 111, 109, 109, 101, 110, 116, 0, 0, 0, 0, 0, 67, 114, 101, 97, 116, 101, 100, 32, 119, 105, 116, 104, 32, 71, 73, 77, 80, 100, 46, 101, 7, 0, 0, 0, 12, 73, 68, 65, 84, 8, 215, 99, 248, 207, 192, 0, 0, 3, 1, 1, 0, 24, 221, 141, 176, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}}"

func TestToFromFile(t *testing.T) {
	path := cachePath(imgName)

	res := NewResource(imgName, imgBytes)
	toFile(res)
	assert.True(t, path != "")
	if !pathExists(path) {
		t.Errorf("%s was not written", imgName)
		return
	}

	res2 := fromFile(imgName)
	assert.True(t, res2 != nil)
	assert.Equal(t, imgName, res2.Name)
	assert.Equal(t, imgBytes, res2.Content)

	os.Remove(path)
}

func TestToGo(t *testing.T) {
	res := NewResource(imgName, imgBytes)

	assert.Equal(t, imgGo, ToGo(res))
}