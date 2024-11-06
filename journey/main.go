package main

import (
	"fmt"
	"sync"
	"time"

	j "github.com/foadmom/myCoach/services/journey"
)

// ============================================================================
// test data
// ============================================================================
var _Node_digbeth j.Location = j.Location{ID: "BIR000", Name: "Digbeth Coach Station", GeoLocation: j.GeoLocation{Lat: 52.47539886357076, Lng: -1.8884546163188745}, Type: j.COACH_STATION}
var _Node_coventry j.Location = j.Location{ID: "COV000", Name: "Coventry Main", GeoLocation: j.GeoLocation{Lat: 52.41142440238643, Lng: -1.5039723332943575}, Type: j.BUS_STOP}
var _Node_northhampton j.Location = j.Location{ID: "NOR000", Name: "Northhampton Coach Station", GeoLocation: j.GeoLocation{Lat: 52.238694983691886, Lng: -0.8978988047538816}, Type: j.COACH_AND_BUS_STATION}
var _Node_milton_keynes j.Location = j.Location{ID: "MKN000", Name: "Milton Keynes Station", GeoLocation: j.GeoLocation{Lat: 52.034607377329756, Lng: -0.7738703731173189}, Type: j.COACH_AND_BUS_STATION}
var _Node_lon_vic j.Location = j.Location{ID: "LON000", Name: "London Victoria Coach Station", GeoLocation: j.GeoLocation{Lat: 51.49254736143251, Lng: -0.14819033650098146}, Type: j.COACH_STATION}
var _Node_manchester j.Location = j.Location{ID: "MNA000", Name: "Manchester Airport Coach Station", GeoLocation: j.GeoLocation{Lat: 53.365127815527146, Lng: -2.273351127768619}, Type: j.COACH_STATION}
var _Node_liverpool j.Location = j.Location{ID: "LIV000", Name: "LIVERPOOL One Bus Station, Canning Place", GeoLocation: j.GeoLocation{Lat: 53.40214866002345, Lng: -2.987504445758943}, Type: j.COACH_STATION}
var _Node_bristol j.Location = j.Location{ID: "BRI000", Name: "BRISTOL Bus & Coach Station", GeoLocation: j.GeoLocation{Lat: 51.459292239749075, Lng: -2.5926834958170972}, Type: j.COACH_STATION}
var _Node_brighton j.Location = j.Location{ID: "GRG000", Name: "Brighton Coach Station", GeoLocation: j.GeoLocation{Lat: 50.81979136982045, Lng: -0.13800739665239234}, Type: j.COACH_STATION}
var _Node_hull j.Location = j.Location{ID: "HUL000", Name: "HULL (City Centre)", GeoLocation: j.GeoLocation{Lat: 53.74479743290693, Lng: -0.34846887628487283}, Type: j.COACH_STATION}
var _Node_leeds j.Location = j.Location{ID: "LEE000", Name: "LEEDS Coach Station, Dyer Street", GeoLocation: j.GeoLocation{Lat: 53.79716973304426, Lng: -1.5362002727619775}, Type: j.COACH_STATION}
var _Node_Portsmouth j.Location = j.Location{ID: "PTM000", Name: "Portsmouth, The Hard", GeoLocation: j.GeoLocation{Lat: 50.797635, Lng: -1.106151}, Type: j.COACH_STATION}
var _Node_Heathrow_T2A3 j.Location = j.Location{ID: "HTH2A3", Name: "Haethrow Terminal 2&3", GeoLocation: j.GeoLocation{Lat: 51.46934739483771, Lng: -0.4517271284645256}, Type: j.COACH_STATION}
var _Node_Heathrow_T5 j.Location = j.Location{ID: "HTH005", Name: "Haethrow Terminal 5", GeoLocation: j.GeoLocation{Lat: 51.47130116397122, Lng: -0.4895683432946781}, Type: j.COACH_STATION}
var _Node_Newcatle_upon_Tyne j.Location = j.Location{ID: "NWC000", Name: "NEWCASTLE upon Tyne, St James Boulevard", GeoLocation: j.GeoLocation{Lat: 54.967511215272715, Lng: -1.6228254701501645}, Type: j.COACH_STATION}
var _Node_Edinburgh j.Location = j.Location{ID: "EDN000", Name: "Edinburgh", GeoLocation: j.GeoLocation{Lat: 55.95548817578951, Lng: -3.1913901320202567}, Type: j.COACH_STATION}

var leg_NX400_S *j.Leg
var leg_NX400_N *j.Leg
var leg_NX171_S *j.Leg
var leg_NX040_E *j.Leg
var leg_NX025_N *j.Leg
var leg_NX025_S *j.Leg
var leg_NX152_E *j.Leg
var leg_NX561_N *j.Leg
var leg_NX133_E *j.Leg
var leg_NX030_W *j.Leg
var leg_NX172_N *j.Leg
var leg_NX598_N *j.Leg

// ============================================================================
//
// ============================================================================
func main() {
	fmt.Println("starting")
	createTestLegs()

	singleRoundTest()

	fmt.Println("exiting")

}

// ============================================================================
// ============================================================================
// ============================================================================
// ============================================================================

func createTestLegs() {
	leg_NX400_S, _ = j.MakeALeg("NX400_S", &_Node_digbeth, &_Node_coventry, &_Node_northhampton, &_Node_milton_keynes, &_Node_lon_vic) // Birmingham to London
	leg_NX400_N, _ = j.MakeALeg("NX400_N", &_Node_lon_vic, &_Node_milton_keynes, &_Node_northhampton, &_Node_coventry, &_Node_digbeth) // London to Birmingham
	leg_NX171_S, _ = j.MakeALeg("NX171_S", &_Node_liverpool, &_Node_manchester, &_Node_digbeth, &_Node_milton_keynes, &_Node_lon_vic)  // Liverpool to London
	leg_NX561_N, _ = j.MakeALeg("NX561_N", &_Node_lon_vic, &_Node_milton_keynes, &_Node_leeds)                                         // London to Leeds
	leg_NX561_N, _ = j.MakeALeg("NX561_EN", &_Node_lon_vic, &_Node_leeds)                                                              // London to Leeds express
	leg_NX040_E, _ = j.MakeALeg("NX040_E", &_Node_bristol, &_Node_lon_vic)                                                             // bristol to london
	leg_NX025_N, _ = j.MakeALeg("NX025_N", &_Node_brighton, &_Node_lon_vic)                                                            // brighton to london
	leg_NX025_S, _ = j.MakeALeg("NX025_S", &_Node_lon_vic, &_Node_brighton)                                                            // London to brighton
	leg_NX152_E, _ = j.MakeALeg("NX152_E", &_Node_digbeth, &_Node_hull)                                                                // birmingham to hull
	leg_NX133_E, _ = j.MakeALeg("NX133_E", &_Node_digbeth, &_Node_leeds)                                                               // Birmingham to leeds
	leg_NX030_W, _ = j.MakeALeg("NX030_W", &_Node_lon_vic, &_Node_Heathrow_T2A3, &_Node_Heathrow_T5, &_Node_Portsmouth)                // London to Portmouth
	leg_NX172_N, _ = j.MakeALeg("NX172_N", &_Node_leeds, &_Node_Newcatle_upon_Tyne)                                                    // leeds to newcastle
	leg_NX598_N, _ = j.MakeALeg("NX598_N", &_Node_Newcatle_upon_Tyne, &_Node_Edinburgh)                                                // newcastle to Edinburgh
}

func singleRoundTest() {
	singlePathTest(&_Node_coventry, &_Node_leeds)
	singlePathTest(&_Node_digbeth, &_Node_leeds)
	singlePathTest(&_Node_digbeth, &_Node_Edinburgh)
}

func singlePathTest(from, to *j.Location) {
	_jm, _ := j.InitialiseJourneyMap(from, to)
	var _conn j.Connection = j.Connection{NestedLevel: 0, FromNode: from}
	_jm.ConnectionTree = append(_jm.ConnectionTree, &_conn)
	start := time.Now()
	callFindRoute(_jm, 0, &_conn)
	elapsed := time.Since(start)
	fmt.Printf("search took %s\n", elapsed)
	_jm.ShowResultingLegs()
	_jm = nil
}

func testNestedConnections() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 1; i > 0; i-- {
		wg.Add(1)
		go testFindPaths(&_Node_coventry, &_Node_leeds, &wg, i)
		wg.Add(1)
		go testFindPaths(&_Node_digbeth, &_Node_leeds, &wg, i)
		wg.Add(1)
		go testFindPaths(&_Node_digbeth, &_Node_Edinburgh, &wg, i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("**** overall search took %s\n", elapsed)

	time.Sleep(1000 * time.Millisecond)

}

func testFindPaths(from, to *j.Location, wg *sync.WaitGroup, count int) {
	_jm, _ := j.InitialiseJourneyMap(from, to)
	var _conn j.Connection = j.Connection{NestedLevel: 0, FromNode: from}
	_jm.ConnectionTree = append(_jm.ConnectionTree, &_conn)
	start := time.Now()
	callFindRoute(_jm, 0, &_conn)
	elapsed := time.Since(start)
	fmt.Printf("search took %s\n", elapsed)
	_jm.ShowResultingLegs()
	_jm = nil
	wg.Done()

}

func callFindRoute(jm *j.JourneyMap, level int, connection *j.Connection) {
	jm.FindConnectingNodes_v5(level, connection)
}
