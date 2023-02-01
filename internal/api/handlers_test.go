package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSong(t *testing.T) {

	testCases := []struct {
		name       string
		langHeader string
		response   string
	}{
		{
			name:       "first API call for en-US",
			langHeader: "en-US",
			response: `10 green bottles standing on a wall,
And if 1 green bottle should accidentally fall,
There’ll be 9 green bottles standing on a wall.

This song is gonna be sung first time!
`,
		},
		{
			name:       "second API call for en-US",
			langHeader: "en-US",
			response: `10 green bottles standing on a wall,
And if 1 green bottle should accidentally fall,
There’ll be 9 green bottles standing on a wall.

This song has already been sung 1 time!
`,
		},
		{
			name:       "third API call for en-US",
			langHeader: "en-US",
			response: `10 green bottles standing on a wall,
And if 1 green bottle should accidentally fall,
There’ll be 9 green bottles standing on a wall.

This song has already been sung 2 times!
`,
		},
		{
			name:       "API call for fr-FR",
			langHeader: "fr-FR",
			response: `10 bouteilles vertes suspendues au mur,
Et si 9 bouteilles vertes venait à tomber,
Il n’y aurait plus que 9 bouteilles vertes suspendues au mur.

Cette chanson a déjà été chantée 3 fois !
`,
		},
		{
			name:       "unsupported language",
			langHeader: "de-DE",
			response: `10 green bottles standing on a wall,
And if 1 green bottle should accidentally fall,
There’ll be 9 green bottles standing on a wall.

This song has already been sung 4 times!
`,
		},
		{
			name:       "invalid header",
			langHeader: "243242!",
			response: `10 green bottles standing on a wall,
And if 1 green bottle should accidentally fall,
There’ll be 9 green bottles standing on a wall.

This song has already been sung 5 times!
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Add("Accept-Language", tc.langHeader)
			rec := httptest.NewRecorder()

			assert.NoError(t, Song(&LanguageContext{e.NewContext(req, rec)}))
			assert.Equal(t, tc.response, rec.Body.String())
		})
	}
}
