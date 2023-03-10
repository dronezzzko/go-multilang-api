// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package translation

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
		"fr_FR": &dictionary{index: fr_FRIndex, data: fr_FRData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"This song has already been sung %d times!": 0,
	"ten-green-bottles-song":                    1,
}

var en_USIndex = []uint32{ // 3 elements
	0x00000000, 0x0000008c, 0x00000113,
} // Size: 36 bytes

const en_USData string = "" + // Size: 275 bytes
	"\x14\x01\x81\x01\x00=\x00'\x02This song is gonna be sung first time!\x02" +
	",\x02This song has already been sung %[1]d time!\x00-\x02This song has a" +
	"lready been sung %[1]d times!\x0210 green bottles standing on a wall," +
	"\x0aAnd if 1 green bottle should accidentally fall,\x0aThere’ll be 9 gre" +
	"en bottles standing on a wall."

var fr_FRIndex = []uint32{ // 3 elements
	0x00000000, 0x000000a9, 0x0000013d,
} // Size: 36 bytes

const fr_FRData string = "" + // Size: 317 bytes
	"\x14\x01\x81\x01\x00=\x009\x02Cette chanson va être chantée pour la prem" +
	"ière fois !\x021\x02Cette chanson a déjà été chantée une fois !\x003\x02" +
	"Cette chanson a déjà été chantée %[1]d fois !\x0210 bouteilles vertes su" +
	"spendues au mur,\x0aEt si 9 bouteilles vertes venait à tomber,\x0aIl n’y" +
	" aurait plus que 9 bouteilles vertes suspendues au mur."

	// Total table size 664 bytes (0KiB); checksum: D6612198
