package traffic

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	rx "github.com/pixel365/goreydenx"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestCountries(t *testing.T) {
	data := "{\"request_id\":\"vZ0Q8ia4ZxpSPd_aA_KySA1701417570.156328\",\"cached\":true,\"cache_expires_at\":\"2023-12-01T19:59:31.198782+00:00\",\"result\":[{\"code\":\"RU\",\"quantity\":914859},{\"code\":\"UA\",\"quantity\":618367},{\"code\":\"RO\",\"quantity\":389216},{\"code\":\"RS\",\"quantity\":290427},{\"code\":\"MD\",\"quantity\":149359},{\"code\":\"IT\",\"quantity\":143308},{\"code\":\"ES\",\"quantity\":135521},{\"code\":\"CZ\",\"quantity\":135172},{\"code\":\"PK\",\"quantity\":134425},{\"code\":\"US\",\"quantity\":131194},{\"code\":\"DE\",\"quantity\":126516},{\"code\":\"BY\",\"quantity\":121646},{\"code\":\"HU\",\"quantity\":110520},{\"code\":\"GB\",\"quantity\":106880},{\"code\":\"MX\",\"quantity\":105114},{\"code\":\"FR\",\"quantity\":93601},{\"code\":\"BG\",\"quantity\":92507},{\"code\":\"HR\",\"quantity\":89363},{\"code\":\"CA\",\"quantity\":87051},{\"code\":\"BA\",\"quantity\":80752},{\"code\":\"CO\",\"quantity\":73306},{\"code\":\"IL\",\"quantity\":68847},{\"code\":\"SK\",\"quantity\":61308},{\"code\":\"IN\",\"quantity\":58126},{\"code\":\"UZ\",\"quantity\":56701},{\"code\":\"AR\",\"quantity\":56605},{\"code\":\"AE\",\"quantity\":55678},{\"code\":\"KG\",\"quantity\":54974},{\"code\":\"TR\",\"quantity\":54230},{\"code\":\"PE\",\"quantity\":52043},{\"code\":\"AZ\",\"quantity\":46433},{\"code\":\"AM\",\"quantity\":40848},{\"code\":\"GE\",\"quantity\":40615},{\"code\":\"MK\",\"quantity\":37427},{\"code\":\"AT\",\"quantity\":36721},{\"code\":\"CL\",\"quantity\":36657},{\"code\":\"EC\",\"quantity\":34846},{\"code\":\"ME\",\"quantity\":34232},{\"code\":\"SA\",\"quantity\":30919},{\"code\":\"SE\",\"quantity\":27698},{\"code\":\"BR\",\"quantity\":23761},{\"code\":\"AL\",\"quantity\":22072},{\"code\":\"IE\",\"quantity\":21037},{\"code\":\"BO\",\"quantity\":20186},{\"code\":\"NP\",\"quantity\":19971},{\"code\":\"GR\",\"quantity\":18822},{\"code\":\"MY\",\"quantity\":18709},{\"code\":\"CH\",\"quantity\":18431},{\"code\":\"PH\",\"quantity\":18107},{\"code\":\"NO\",\"quantity\":18042},{\"code\":\"KR\",\"quantity\":17895},{\"code\":\"LT\",\"quantity\":16885},{\"code\":\"PT\",\"quantity\":15889},{\"code\":\"BD\",\"quantity\":15621},{\"code\":\"VE\",\"quantity\":14880},{\"code\":\"DO\",\"quantity\":13796},{\"code\":\"XK\",\"quantity\":13336},{\"code\":\"SI\",\"quantity\":13253},{\"code\":\"KW\",\"quantity\":12881},{\"code\":\"DK\",\"quantity\":12666},{\"code\":\"AU\",\"quantity\":12472},{\"code\":\"SG\",\"quantity\":11822},{\"code\":\"CY\",\"quantity\":11445},{\"code\":\"GT\",\"quantity\":10747},{\"code\":\"JP\",\"quantity\":10593},{\"code\":\"NL\",\"quantity\":10586},{\"code\":\"PA\",\"quantity\":10258},{\"code\":\"QA\",\"quantity\":9750},{\"code\":\"TH\",\"quantity\":9689},{\"code\":\"PL\",\"quantity\":9297},{\"code\":\"ID\",\"quantity\":8838},{\"code\":\"CR\",\"quantity\":8825},{\"code\":\"SV\",\"quantity\":8811},{\"code\":\"KZ\",\"quantity\":7780},{\"code\":\"DZ\",\"quantity\":7588},{\"code\":\"HK\",\"quantity\":6939},{\"code\":\"UY\",\"quantity\":6835},{\"code\":\"MU\",\"quantity\":6674},{\"code\":\"EG\",\"quantity\":6324},{\"code\":\"HN\",\"quantity\":6201},{\"code\":\"TJ\",\"quantity\":6181},{\"code\":\"PY\",\"quantity\":5321},{\"code\":\"LV\",\"quantity\":5289},{\"code\":\"ZA\",\"quantity\":4860},{\"code\":\"IR\",\"quantity\":4222},{\"code\":\"BE\",\"quantity\":4028},{\"code\":\"NZ\",\"quantity\":3930},{\"code\":\"BH\",\"quantity\":3914},{\"code\":\"EE\",\"quantity\":3863},{\"code\":\"PR\",\"quantity\":3760},{\"code\":\"NI\",\"quantity\":3572},{\"code\":\"LK\",\"quantity\":3541},{\"code\":\"VN\",\"quantity\":3533},{\"code\":\"OM\",\"quantity\":3367},{\"code\":\"FI\",\"quantity\":3115},{\"code\":\"MT\",\"quantity\":2941},{\"code\":\"MN\",\"quantity\":2629},{\"code\":\"CN\",\"quantity\":2565},{\"code\":\"RE\",\"quantity\":2023},{\"code\":\"TN\",\"quantity\":1984},{\"code\":\"TW\",\"quantity\":1839},{\"code\":\"JO\",\"quantity\":1833},{\"code\":\"LU\",\"quantity\":1792},{\"code\":\"HT\",\"quantity\":1713},{\"code\":\"CI\",\"quantity\":1624},{\"code\":\"FJ\",\"quantity\":1511},{\"code\":\"SR\",\"quantity\":1397},{\"code\":\"MG\",\"quantity\":1298},{\"code\":\"MV\",\"quantity\":1279},{\"code\":\"KE\",\"quantity\":1199},{\"code\":\"GA\",\"quantity\":1117},{\"code\":\"LB\",\"quantity\":1000},{\"code\":\"TZ\",\"quantity\":980},{\"code\":\"PF\",\"quantity\":969},{\"code\":\"IQ\",\"quantity\":942},{\"code\":\"MQ\",\"quantity\":846},{\"code\":\"MA\",\"quantity\":791},{\"code\":\"TG\",\"quantity\":768},{\"code\":\"IS\",\"quantity\":739},{\"code\":\"GP\",\"quantity\":692},{\"code\":\"BN\",\"quantity\":675},{\"code\":\"ML\",\"quantity\":576},{\"code\":\"TT\",\"quantity\":572},{\"code\":\"CU\",\"quantity\":569},{\"code\":\"CG\",\"quantity\":556},{\"code\":\"MM\",\"quantity\":552},{\"code\":\"KH\",\"quantity\":544},{\"code\":\"NC\",\"quantity\":508},{\"code\":\"NG\",\"quantity\":506},{\"code\":\"GU\",\"quantity\":475},{\"code\":\"TM\",\"quantity\":450},{\"code\":\"JM\",\"quantity\":444},{\"code\":\"CM\",\"quantity\":443},{\"code\":\"ET\",\"quantity\":428},{\"code\":\"MO\",\"quantity\":372},{\"code\":\"RW\",\"quantity\":369},{\"code\":\"BJ\",\"quantity\":361},{\"code\":\"AO\",\"quantity\":353},{\"code\":\"AD\",\"quantity\":351},{\"code\":\"SY\",\"quantity\":321},{\"code\":\"GF\",\"quantity\":297},{\"code\":\"SN\",\"quantity\":291},{\"code\":\"BZ\",\"quantity\":267},{\"code\":\"SO\",\"quantity\":264},{\"code\":\"CW\",\"quantity\":264},{\"code\":\"PS\",\"quantity\":247},{\"code\":\"UG\",\"quantity\":246},{\"code\":\"MZ\",\"quantity\":219},{\"code\":\"GY\",\"quantity\":216},{\"code\":\"YE\",\"quantity\":200},{\"code\":\"ZM\",\"quantity\":191},{\"code\":\"BF\",\"quantity\":187},{\"code\":\"GH\",\"quantity\":185},{\"code\":\"BS\",\"quantity\":174},{\"code\":\"MW\",\"quantity\":174},{\"code\":\"MR\",\"quantity\":169},{\"code\":\"MP\",\"quantity\":169},{\"code\":\"SC\",\"quantity\":167},{\"code\":\"AW\",\"quantity\":158},{\"code\":\"BB\",\"quantity\":156},{\"code\":\"LY\",\"quantity\":155},{\"code\":\"BW\",\"quantity\":152},{\"code\":\"CD\",\"quantity\":151},{\"code\":\"GQ\",\"quantity\":151},{\"code\":\"AF\",\"quantity\":150},{\"code\":\"NA\",\"quantity\":149},{\"code\":\"MC\",\"quantity\":146},{\"code\":\"LA\",\"quantity\":145},{\"code\":\"YT\",\"quantity\":141},{\"code\":\"KY\",\"quantity\":136},{\"code\":\"LI\",\"quantity\":134},{\"code\":\"BT\",\"quantity\":134},{\"code\":\"GN\",\"quantity\":126},{\"code\":\"JE\",\"quantity\":116},{\"code\":\"GI\",\"quantity\":104},{\"code\":\"FO\",\"quantity\":101},{\"code\":\"DJ\",\"quantity\":101},{\"code\":\"ZW\",\"quantity\":93},{\"code\":\"SM\",\"quantity\":81},{\"code\":\"CV\",\"quantity\":81},{\"code\":\"AX\",\"quantity\":76},{\"code\":\"BM\",\"quantity\":74},{\"code\":\"TC\",\"quantity\":70},{\"code\":\"BI\",\"quantity\":66},{\"code\":\"LR\",\"quantity\":66},{\"code\":\"IM\",\"quantity\":61},{\"code\":\"DM\",\"quantity\":59},{\"code\":\"NE\",\"quantity\":54},{\"code\":\"LC\",\"quantity\":52},{\"code\":\"SX\",\"quantity\":50},{\"code\":\"GG\",\"quantity\":49},{\"code\":\"KM\",\"quantity\":46},{\"code\":\"GW\",\"quantity\":42},{\"code\":\"SL\",\"quantity\":41},{\"code\":\"WS\",\"quantity\":37},{\"code\":\"GD\",\"quantity\":37},{\"code\":\"AS\",\"quantity\":34},{\"code\":\"BQ\",\"quantity\":33},{\"code\":\"MF\",\"quantity\":32},{\"code\":\"VI\",\"quantity\":32},{\"code\":\"SZ\",\"quantity\":31},{\"code\":\"VG\",\"quantity\":29},{\"code\":\"GM\",\"quantity\":28},{\"code\":\"PG\",\"quantity\":28},{\"code\":\"PW\",\"quantity\":23},{\"code\":\"VC\",\"quantity\":23},{\"code\":\"SS\",\"quantity\":23},{\"code\":\"AG\",\"quantity\":22},{\"code\":\"VU\",\"quantity\":21},{\"code\":\"MH\",\"quantity\":16},{\"code\":\"GL\",\"quantity\":16},{\"code\":\"LS\",\"quantity\":16},{\"code\":\"TD\",\"quantity\":16},{\"code\":\"KN\",\"quantity\":12},{\"code\":\"FM\",\"quantity\":8},{\"code\":\"SH\",\"quantity\":7},{\"code\":\"PM\",\"quantity\":6},{\"code\":\"BL\",\"quantity\":4},{\"code\":\"TO\",\"quantity\":4},{\"code\":\"SD\",\"quantity\":2},{\"code\":\"ST\",\"quantity\":2},{\"code\":\"AI\",\"quantity\":2},{\"code\":\"FK\",\"quantity\":2},{\"code\":\"ER\",\"quantity\":2},{\"code\":\"CF\",\"quantity\":1},{\"code\":\"TL\",\"quantity\":1},{\"code\":\"MS\",\"quantity\":1}]}"
	client := rx.NewClient("", "")
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(data)),
			Header:     make(http.Header),
		}
	})
	client.Token = &rx.Token{
		AccessToken: "fake token",
		ExpiresIn:   "fake date",
	}

	_, err := Countries(client)
	if err != nil {
		t.Error(err)
	}
}

func TestLanguages(t *testing.T) {
	data := "{\"request_id\":\"9dVY5TjzjEt52qNE459I3Q1701417786.392909\",\"cached\":true,\"cache_expires_at\":\"2023-12-01T20:03:07.426154+00:00\",\"result\":[{\"code\":\"RU\",\"quantity\":2127119},{\"code\":\"EN\",\"quantity\":1262826},{\"code\":\"ES\",\"quantity\":584676},{\"code\":\"RO\",\"quantity\":351696},{\"code\":\"UK\",\"quantity\":229539},{\"code\":\"SR\",\"quantity\":180304},{\"code\":\"HU\",\"quantity\":124098},{\"code\":\"HR\",\"quantity\":121787},{\"code\":\"IT\",\"quantity\":102289},{\"code\":\"CS\",\"quantity\":96808},{\"code\":\"FR\",\"quantity\":87109},{\"code\":\"BG\",\"quantity\":75961},{\"code\":\"DE\",\"quantity\":55428},{\"code\":\"SK\",\"quantity\":52976},{\"code\":\"PT\",\"quantity\":28229},{\"code\":\"TR\",\"quantity\":28059},{\"code\":\"BS\",\"quantity\":19737},{\"code\":\"HE\",\"quantity\":15364},{\"code\":\"LT\",\"quantity\":11833},{\"code\":\"KA\",\"quantity\":10964},{\"code\":\"EL\",\"quantity\":9063},{\"code\":\"MK\",\"quantity\":8460},{\"code\":\"AR\",\"quantity\":7653},{\"code\":\"SV\",\"quantity\":7131},{\"code\":\"PL\",\"quantity\":6722},{\"code\":\"ID\",\"quantity\":5231},{\"code\":\"SQ\",\"quantity\":5062},{\"code\":\"SL\",\"quantity\":4589},{\"code\":\"NL\",\"quantity\":3889},{\"code\":\"NB\",\"quantity\":3592},{\"code\":\"FA\",\"quantity\":3147},{\"code\":\"UZ\",\"quantity\":3003},{\"code\":\"AZ\",\"quantity\":2891},{\"code\":\"LV\",\"quantity\":2242},{\"code\":\"DA\",\"quantity\":2100},{\"code\":\"HY\",\"quantity\":1993},{\"code\":\"TH\",\"quantity\":1935},{\"code\":\"CA\",\"quantity\":1932},{\"code\":\"JA\",\"quantity\":1846},{\"code\":\"KO\",\"quantity\":1449},{\"code\":\"MS\",\"quantity\":1308},{\"code\":\"ZH\",\"quantity\":1133},{\"code\":\"VI\",\"quantity\":1083},{\"code\":\"FI\",\"quantity\":792},{\"code\":\"BE\",\"quantity\":384},{\"code\":\"ET\",\"quantity\":259},{\"code\":\"KY\",\"quantity\":219},{\"code\":\"MN\",\"quantity\":207},{\"code\":\"HI\",\"quantity\":188},{\"code\":\"OS\",\"quantity\":165},{\"code\":\"UR\",\"quantity\":141},{\"code\":\"EU\",\"quantity\":129},{\"code\":\"BN\",\"quantity\":127},{\"code\":\"KK\",\"quantity\":113},{\"code\":\"PA\",\"quantity\":99},{\"code\":\"TE\",\"quantity\":94},{\"code\":\"CE\",\"quantity\":81},{\"code\":\"RM\",\"quantity\":73},{\"code\":\"GU\",\"quantity\":64},{\"code\":\"GL\",\"quantity\":62},{\"code\":\"TK\",\"quantity\":59},{\"code\":\"NO\",\"quantity\":58},{\"code\":\"NN\",\"quantity\":56},{\"code\":\"NE\",\"quantity\":53},{\"code\":\"AF\",\"quantity\":53},{\"code\":\"SA\",\"quantity\":41},{\"code\":\"MR\",\"quantity\":39},{\"code\":\"LI\",\"quantity\":37},{\"code\":\"LA\",\"quantity\":33},{\"code\":\"AM\",\"quantity\":29},{\"code\":\"EG\",\"quantity\":29},{\"code\":\"BO\",\"quantity\":21},{\"code\":\"BA\",\"quantity\":21},{\"code\":\"SN\",\"quantity\":19},{\"code\":\"KN\",\"quantity\":19},{\"code\":\"HT\",\"quantity\":19},{\"code\":\"TG\",\"quantity\":18},{\"code\":\"KS\",\"quantity\":18},{\"code\":\"MF\",\"quantity\":14},{\"code\":\"MY\",\"quantity\":13},{\"code\":\"IS\",\"quantity\":12},{\"code\":\"AS\",\"quantity\":11},{\"code\":\"TA\",\"quantity\":11},{\"code\":\"AK\",\"quantity\":10},{\"code\":\"UN\",\"quantity\":9},{\"code\":\"KW\",\"quantity\":9},{\"code\":\"LO\",\"quantity\":9},{\"code\":\"SI\",\"quantity\":8},{\"code\":\"ML\",\"quantity\":8},{\"code\":\"KB\",\"quantity\":7},{\"code\":\"TT\",\"quantity\":7},{\"code\":\"AV\",\"quantity\":7},{\"code\":\"MG\",\"quantity\":6},{\"code\":\"BH\",\"quantity\":6},{\"code\":\"SO\",\"quantity\":5},{\"code\":\"FU\",\"quantity\":5},{\"code\":\"AN\",\"quantity\":4},{\"code\":\"DV\",\"quantity\":4},{\"code\":\"ME\",\"quantity\":4},{\"code\":\"GS\",\"quantity\":4},{\"code\":\"AD\",\"quantity\":3},{\"code\":\"PS\",\"quantity\":3},{\"code\":\"JO\",\"quantity\":2},{\"code\":\"SW\",\"quantity\":2},{\"code\":\"QU\",\"quantity\":1},{\"code\":\"OR\",\"quantity\":1},{\"code\":\"YI\",\"quantity\":1},{\"code\":\"NY\",\"quantity\":1},{\"code\":\"BM\",\"quantity\":1}]}"
	client := rx.NewClient("", "")
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(data)),
			Header:     make(http.Header),
		}
	})
	client.Token = &rx.Token{
		AccessToken: "fake token",
		ExpiresIn:   "fake date",
	}

	_, err := Languages(client)
	if err != nil {
		t.Error(err)
	}
}

func TestDevices(t *testing.T) {
	data := "{\"request_id\":\"XntHUXu9npR6KQUt-sVXGQ1706619012.52129\",\"cached\":false,\"cache_expires_at\":\"2024-01-30T18:50:13.549227+00:00\",\"result\":[{\"code\":\"desktop\",\"quantity\":1375379},{\"code\":\"mobile\",\"quantity\":1883065},{\"code\":\"tablet\",\"quantity\":35829}]}"
	client := rx.NewClient("", "")
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(data)),
			Header:     make(http.Header),
		}
	})
	client.Token = &rx.Token{
		AccessToken: "fake token",
		ExpiresIn:   "fake date",
	}

	_, err := Devices(client)
	if err != nil {
		t.Error(err)
	}
}
