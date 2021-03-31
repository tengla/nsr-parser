package main

// KeyList
type KeyList struct {
	KeyValue []struct {
		Key   string
		Value string
	}
}

// Centroid
type Centroid struct {
	Location struct {
		Longitude float32
		Latitude  float32
	}
}

// PlaceEquipments
type PlaceEquipments struct {
	Id                   string `xml:"id,attr"`
	WaitingRoomEquipment struct {
		Id           string `xml:"id,attr"`
		Modification string `xml:"modification,attr"`
		Version      string `xml:"version,attr"`
		Sanitary     string
	}
	GeneralSign struct {
		PrivateCode     string
		SignContentType string
	}
}

// AccessibilityAssessment
type AccessibilityAssessment struct {
	Modification           string `xml:"modification,attr"`
	Version                string `xml:"version,attr"`
	Id                     string `xml:"id,attr"`
	MobilityImpairedAccess string
	Limitations            []struct {
		AccessibilityLimitation struct {
			Modification            string `xml:"modification,attr"`
			Version                 string `xml:"version,attr"`
			Id                      string `xml:"id,attr"`
			WheelchairAccess        string
			StepFreeAccess          string
			EscalatorFreeAccess     string
			LiftFreeAccess          string
			AudibleSignalsAvailable string
		}
	} `xml:"limitations"`
}

// Quay
type Quay struct {
	Id          string `xml:"id,attr"`
	Description struct {
		Lang    string `xml:"lang,attr"`
		Content string `xml:",chardata"`
	}
	PlaceEquipments     PlaceEquipments
	OtherTransportModes string
	KeyList             []KeyList `xml:"keyList"`
}

type StopPlace struct {
	Created      string `xml:"created,attr"`
	Id           string `xml:"id,attr"`
	Name         string
	ValidBetween struct {
		FromDate string
	}
	KeyList                 []KeyList `xml:"keyList"`
	Centroid                Centroid
	PlaceEquipments         PlaceEquipments `xml:"placeEquipments"`
	AccessibilityAssessment AccessibilityAssessment
	TopographicPlaceRef     struct {
		Ref     string `xml:"ref,attr"`
		Created string `xml:"created,attr"`
		Version string `xml:"version,attr"`
	}
	TransportMode       string
	OtherTransportModes string
	TariffZones         []struct {
		TariffZoneRef struct {
			Ref string `xml:"ref,attr"`
		}
	} `xml:"tariffZones"`
	StopPlaceType string
	Weighting     string
	ParentSiteRef struct {
		Ref     string `xml:"ref,attr"`
		Version string `xml:"version,attr"`
	}
	Quays []struct {
		Quay Quay
	} `xml:"quays"`
}

// Doc
type Doc struct {
	StopPlaces []StopPlace `xml:"StopPlace"`
}
