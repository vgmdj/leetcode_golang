package Count_and_Say

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSay(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(countAndSay(1), "1")
	ast.Equal(countAndSay(6), "312211")
	ast.Equal(countAndSay(20), "1113122113121113222123211211131211121311121321123113213221121113122"+
		"1131211221321123113213221123113112221131112311332211211131221131211132211121312211231131112311211232"+
		"2211213211321322113311213212312311211131122211213211331121321123123211231131122211211131221131211131"+
		"23112112322111213211322211312113211")

}
