package parsename

import (
	"testing"
)

var names = []Name{
	{"Edwood Ocasio", "Edwood", "", "Ocasio", ""},
	{"Edwood Ocasio Vicente", "Edwood", "", "Ocasio", "Vicente"},
	{"Manuel Joel Vicente Yera", "Manuel", "Joel", "Vicente", "Yera"},
	{"Maria L.M. Rivera Torres", "Maria", "L.M.", "Rivera", "Torres"},
	{"Nestor García de Jesus", "Nestor", "", "García", "de Jesus"},
	{"Jose Colón de la Torre", "Jose", "", "Colón", "de la Torre"},
	{"Jose De los Angeles De Jesus", "Jose", "", "De los Angeles", "De Jesus"},
	{"Mario C. De los Angeles De la Torres", "Mario", "C.", "De los Angeles", "De la Torres"},
	{"MARIA DEL C COSTA", "MARIA", "DEL C", "COSTA", ""},
	{"MARIA DEL C. COSTA", "MARIA", "DEL C.", "COSTA", ""},
	{"MARIA DEL CARMEN COSTA", "MARIA", "", "DEL CARMEN", "COSTA"},
	{"CYNTHIA M. CHARLESTONE", "CYNTHIA", "M.", "CHARLESTONE", ""},
	{"JOHN JAMES GONZALEZ ORTIZ", "JOHN", "JAMES", "GONZALEZ", "ORTIZ"},
	{"SONIALY A MC CLINTOSH", "SONIALY", "A", "MC", "CLINTOSH"},
	{"MARIA DE TORRES CRUZ", "MARIA", "", "DE TORRES", "CRUZ"},
	{"LUIS LA TORRE", "LUIS", "", "LA TORRE", ""},
	{"VENICIO DEL TORO", "VENICIO", "", "DEL TORO", ""},
	{"VENICIO DE ARMAS", "VENICIO", "", "DE ARMAS", ""},
	{"José E. Colón Rodríguez", "José", "E.", "Colón", "Rodríguez"},
	{"María de los Angeles Merced", "María", "", "de los Angeles", "Merced"},
	{Input: "LUIS T."},
	{Input: "LUIS DE"},
	{Input: "LUIS LA"},
	{Input: "LUIS A"},
	{Input: " "},
}

func TestParseName(t *testing.T) {
	for _, want := range names {
		got, err := New(want.Input)
		if err != nil {
			t.Logf("Parse error for input %q, %v", want.Input, err)
			continue
		}
		if want.FirstName != got.FirstName || want.MiddleName != got.MiddleName ||
			want.LastName != got.LastName || want.MaidenName != got.MaidenName {
			t.Errorf("Parse error: wanted %#v , got %#v", want, got)
			continue
		}
		t.Logf("%s", got)
	}
}

func BenchmarkParseName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(names[0].Input)
	}
}
