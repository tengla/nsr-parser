package main

// NSR document structs

// StopPlace
type StopPlace struct {
	Text         string `xml:",chardata"`
	Created      string `xml:"created,attr"`
	Changed      string `xml:"changed,attr"`
	Modification string `xml:"modification,attr"`
	Version      string `xml:"version,attr"`
	ID           string `xml:"id,attr"`
	ValidBetween struct {
		Text     string `xml:",chardata"`
		FromDate string `xml:"FromDate"` // 2017-06-19T19:12:29.887, ...
		ToDate   string `xml:"ToDate"`   // 2021-05-02T01:00:09, 2021...
	} `xml:"ValidBetween"`
	KeyList struct {
		Text     string `xml:",chardata"`
		KeyValue []struct {
			Text  string `xml:",chardata"`
			Key   string `xml:"Key"`   // jbvCode, uicCode, iffCode...
			Value string `xml:"Value"` // GU, 7601213, 7601213, 121...
		} `xml:"KeyValue"`
	} `xml:"keyList"`
	Name struct {
		Text string `xml:",chardata"` // Gudå, Eikenes stasjon, H...
		Lang string `xml:"lang,attr"`
	} `xml:"Name"`
	Centroid struct {
		Text     string `xml:",chardata"`
		Location struct {
			Text      string `xml:",chardata"`
			Longitude string `xml:"Longitude"` // 11.620069, 9.895932, 11.1...
			Latitude  string `xml:"Latitude"`  // 63.444097, 59.146811, 62....
		} `xml:"Location"`
	} `xml:"Centroid"`
	AccessibilityAssessment struct {
		Text                   string `xml:",chardata"`
		Modification           string `xml:"modification,attr"`
		Version                string `xml:"version,attr"`
		ID                     string `xml:"id,attr"`
		MobilityImpairedAccess string `xml:"MobilityImpairedAccess"` // partial, partial, partial...
		Limitations            struct {
			Text                    string `xml:",chardata"`
			AccessibilityLimitation struct {
				Text                    string `xml:",chardata"`
				Modification            string `xml:"modification,attr"`
				Version                 string `xml:"version,attr"`
				ID                      string `xml:"id,attr"`
				WheelchairAccess        string `xml:"WheelchairAccess"`        // true, false, true, true, ...
				StepFreeAccess          string `xml:"StepFreeAccess"`          // true, false, true, true, ...
				EscalatorFreeAccess     string `xml:"EscalatorFreeAccess"`     // unknown, unknown, unknown...
				LiftFreeAccess          string `xml:"LiftFreeAccess"`          // unknown, unknown, unknown...
				AudibleSignalsAvailable string `xml:"AudibleSignalsAvailable"` // unknown, unknown, unknown...
			} `xml:"AccessibilityLimitation"`
		} `xml:"limitations"`
	} `xml:"AccessibilityAssessment"`
	TopographicPlaceRef struct {
		Text    string `xml:",chardata"`
		Ref     string `xml:"ref,attr"`
		Created string `xml:"created,attr"`
		Version string `xml:"version,attr"`
	} `xml:"TopographicPlaceRef"`
	TransportMode       string `xml:"TransportMode"` // bus, rail, bus, bus, bus,...
	BusSubmode          string `xml:"BusSubmode"`    // railReplacementBus, railR...
	OtherTransportModes string `xml:"OtherTransportModes"`
	TariffZones         struct {
		Text          string `xml:",chardata"`
		TariffZoneRef []struct {
			Text    string `xml:",chardata"`
			Ref     string `xml:"ref,attr"`
			Version string `xml:"version,attr"`
		} `xml:"TariffZoneRef"`
	} `xml:"tariffZones"`
	StopPlaceType string `xml:"StopPlaceType"` // onstreetBus, railStation,...
	Weighting     string `xml:"Weighting"`     // interchangeAllowed, inter...
	Quays         struct {
		Text string `xml:",chardata"`
		Quay []struct {
			Text         string `xml:",chardata"`
			Modification string `xml:"modification,attr"`
			Version      string `xml:"version,attr"`
			ID           string `xml:"id,attr"`
			Changed      string `xml:"changed,attr"`
			Created      string `xml:"created,attr"`
			KeyList      struct {
				Text     string `xml:",chardata"`
				KeyValue []struct {
					Text  string `xml:",chardata"`
					Key   string `xml:"Key"`   // uicCode, imported-id, gra...
					Value string `xml:"Value"` // 7601213, NSB:Quay:7601213...
				} `xml:"KeyValue"`
			} `xml:"keyList"`
			Centroid struct {
				Text     string `xml:",chardata"`
				Location struct {
					Text      string `xml:",chardata"`
					Longitude string `xml:"Longitude"` // 11.620069, 9.895932, 11.1...
					Latitude  string `xml:"Latitude"`  // 63.444097, 59.146811, 62....
				} `xml:"Location"`
			} `xml:"Centroid"`
			OtherTransportModes string `xml:"OtherTransportModes"`
			PrivateCode         struct {
				Text string `xml:",chardata"` // 1, 1, 1, 1, 1, 1, 01, 01,...
				Type string `xml:"type,attr"`
			} `xml:"PrivateCode"`
			PublicCode     string `xml:"PublicCode"`     // 1, 1, 1, 1, 1, 1, O, 2, 1...
			CompassBearing string `xml:"CompassBearing"` // 201.0, 59.0, 34.0, 9.0, 3...
			Name           struct {
				Text string `xml:",chardata"` // Bukkholmveien 1, Ise spor...
				Lang string `xml:"lang,attr"`
			} `xml:"Name"`
			Description struct {
				Text string `xml:",chardata"` // Retning Frognerseteren, R...
				Lang string `xml:"lang,attr"`
			} `xml:"Description"`
			PlaceEquipments struct {
				Text        string `xml:",chardata"`
				ID          string `xml:"id,attr"`
				GeneralSign []struct {
					Text            string `xml:",chardata"`
					Modification    string `xml:"modification,attr"`
					Version         string `xml:"version,attr"`
					ID              string `xml:"id,attr"`
					PrivateCode     string `xml:"PrivateCode"`     // 512, 512, 512, 512, 512, ...
					SignContentType string `xml:"SignContentType"` // transportMode, transportM...
					Content         struct {
						Text string `xml:",chardata"` // Timetable, Timetable, Tim...
						Lang string `xml:"lang,attr"`
					} `xml:"Content"`
				} `xml:"GeneralSign"`
				ShelterEquipment []struct {
					Text         string `xml:",chardata"`
					Modification string `xml:"modification,attr"`
					Version      string `xml:"version,attr"`
					ID           string `xml:"id,attr"`
					Enclosed     string `xml:"Enclosed"` // true, true, true, true, t...
					Seats        string `xml:"Seats"`    // 1, 1, 1, 1, 1, 1, 1, 1, 1...
					StepFree     string `xml:"StepFree"` // false, false, false, fals...
				} `xml:"ShelterEquipment"`
				WaitingRoomEquipment []struct {
					Text         string `xml:",chardata"`
					Modification string `xml:"modification,attr"`
					Version      string `xml:"version,attr"`
					ID           string `xml:"id,attr"`
					Sanitary     string `xml:"Sanitary"`
					Seats        string `xml:"Seats"`    // 1, 0
					StepFree     string `xml:"StepFree"` // false, false
					Heated       string `xml:"Heated"`   // false, false
				} `xml:"WaitingRoomEquipment"`
				SanitaryEquipment struct {
					Text                 string `xml:",chardata"`
					Modification         string `xml:"modification,attr"`
					Version              string `xml:"version,attr"`
					ID                   string `xml:"id,attr"`
					Gender               string `xml:"Gender"` // both, both, both, both, b...
					SanitaryFacilityList string `xml:"SanitaryFacilityList"`
					PaymentMethods       string `xml:"PaymentMethods"`
					NumberOfToilets      string `xml:"NumberOfToilets"` // 1
				} `xml:"SanitaryEquipment"`
				TicketingEquipment struct {
					Text                    string `xml:",chardata"`
					Modification            string `xml:"modification,attr"`
					Version                 string `xml:"version,attr"`
					ID                      string `xml:"id,attr"`
					VehicleModes            string `xml:"VehicleModes"`
					TicketMachines          string `xml:"TicketMachines"`   // true, true, true, true, t...
					NumberOfMachines        string `xml:"NumberOfMachines"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
					TicketingFacilityList   string `xml:"TicketingFacilityList"`
					TicketOffice            string `xml:"TicketOffice"` // true, true, true, true, t...
					PaymentMethods          string `xml:"PaymentMethods"`
					TicketTypesAvailable    string `xml:"TicketTypesAvailable"`
					ScopeOfTicketsAvailable string `xml:"ScopeOfTicketsAvailable"`
				} `xml:"TicketingEquipment"`
			} `xml:"placeEquipments"`
			AccessibilityAssessment struct {
				Text                   string `xml:",chardata"`
				Modification           string `xml:"modification,attr"`
				Version                string `xml:"version,attr"`
				ID                     string `xml:"id,attr"`
				MobilityImpairedAccess string `xml:"MobilityImpairedAccess"` // partial, partial, partial...
				Limitations            struct {
					Text                    string `xml:",chardata"`
					AccessibilityLimitation struct {
						Text                    string `xml:",chardata"`
						Modification            string `xml:"modification,attr"`
						Version                 string `xml:"version,attr"`
						ID                      string `xml:"id,attr"`
						WheelchairAccess        string `xml:"WheelchairAccess"`        // false, false, false, true...
						StepFreeAccess          string `xml:"StepFreeAccess"`          // false, true, false, true,...
						EscalatorFreeAccess     string `xml:"EscalatorFreeAccess"`     // unknown, unknown, unknown...
						LiftFreeAccess          string `xml:"LiftFreeAccess"`          // unknown, unknown, unknown...
						AudibleSignalsAvailable string `xml:"AudibleSignalsAvailable"` // unknown, unknown, unknown...
					} `xml:"AccessibilityLimitation"`
				} `xml:"limitations"`
			} `xml:"AccessibilityAssessment"`
			ValidBetween struct {
				Text     string `xml:",chardata"`
				FromDate string `xml:"FromDate"` // 2017-06-21T17:01:09.146, ...
			} `xml:"ValidBetween"`
		} `xml:"Quay"`
	} `xml:"quays"`
	PlaceEquipments struct {
		Text             string `xml:",chardata"`
		ID               string `xml:"id,attr"`
		ShelterEquipment struct {
			Text         string `xml:",chardata"`
			Modification string `xml:"modification,attr"`
			Version      string `xml:"version,attr"`
			ID           string `xml:"id,attr"`
			Enclosed     string `xml:"Enclosed"` // true, true, false, false,...
			Seats        string `xml:"Seats"`    // 1, 1, 1, 1, 1, 1, 1, 1, 1...
			StepFree     string `xml:"StepFree"` // false, false, false, fals...
		} `xml:"ShelterEquipment"`
		WaitingRoomEquipment struct {
			Text         string `xml:",chardata"`
			Modification string `xml:"modification,attr"`
			Version      string `xml:"version,attr"`
			ID           string `xml:"id,attr"`
			Sanitary     string `xml:"Sanitary"`
			Seats        string `xml:"Seats"`    // 1, 1, 1, 1, 1, 1, 1, 1, 1...
			StepFree     string `xml:"StepFree"` // false, false, false, fals...
			Heated       string `xml:"Heated"`   // false, false, false, fals...
		} `xml:"WaitingRoomEquipment"`
		TicketingEquipment struct {
			Text                    string `xml:",chardata"`
			Modification            string `xml:"modification,attr"`
			Version                 string `xml:"version,attr"`
			ID                      string `xml:"id,attr"`
			VehicleModes            string `xml:"VehicleModes"`
			TicketMachines          string `xml:"TicketMachines"`   // true, true, true, true, t...
			NumberOfMachines        string `xml:"NumberOfMachines"` // 1, 1, 1, 1, 1, 1, 2, 2, 1...
			TicketingFacilityList   string `xml:"TicketingFacilityList"`
			TicketOffice            string `xml:"TicketOffice"` // false, false, false, fals...
			PaymentMethods          string `xml:"PaymentMethods"`
			TicketTypesAvailable    string `xml:"TicketTypesAvailable"`
			ScopeOfTicketsAvailable string `xml:"ScopeOfTicketsAvailable"`
		} `xml:"TicketingEquipment"`
		GeneralSign struct {
			Text            string `xml:",chardata"`
			Modification    string `xml:"modification,attr"`
			Version         string `xml:"version,attr"`
			ID              string `xml:"id,attr"`
			PrivateCode     string `xml:"PrivateCode"`     // 512, 512, 512, 512, 512, ...
			SignContentType string `xml:"SignContentType"` // transportMode, transportM...
		} `xml:"GeneralSign"`
		SanitaryEquipment struct {
			Text                 string `xml:",chardata"`
			Modification         string `xml:"modification,attr"`
			Version              string `xml:"version,attr"`
			ID                   string `xml:"id,attr"`
			Gender               string `xml:"Gender"` // both, both, both, both, b...
			SanitaryFacilityList string `xml:"SanitaryFacilityList"`
			PaymentMethods       string `xml:"PaymentMethods"`
			NumberOfToilets      string `xml:"NumberOfToilets"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
		} `xml:"SanitaryEquipment"`
	} `xml:"placeEquipments"`
	Description struct {
		Text string `xml:",chardata"` // skole, Oppøyen, i Sognad...
		Lang string `xml:"lang,attr"`
	} `xml:"Description"`
	AlternativeNames struct {
		Text            string `xml:",chardata"`
		AlternativeName []struct {
			Text         string `xml:",chardata"`
			Modification string `xml:"modification,attr"`
			Version      string `xml:"version,attr"`
			ID           string `xml:"id,attr"`
			NameType     string `xml:"NameType"` // alias, alias, alias, alia...
			Name         struct {
				Text string `xml:",chardata"` // Hoelsgård hotel, Nysvemo...
				Lang string `xml:"lang,attr"`
			} `xml:"Name"`
		} `xml:"AlternativeName"`
	} `xml:"alternativeNames"`
	RailSubmode   string `xml:"RailSubmode"`  // touristRailway, touristRa...
	WaterSubmode  string `xml:"WaterSubmode"` // localCarFerry, localPasse...
	MetroSubmode  string `xml:"MetroSubmode"` // metro, metro, metro, metr...
	ParentSiteRef struct {
		Text    string `xml:",chardata"`
		Ref     string `xml:"ref,attr"`
		Version string `xml:"version,attr"`
	} `xml:"ParentSiteRef"`
	AirSubmode       string `xml:"AirSubmode"`       // helicopterService, unknow...
	TelecabinSubmode string `xml:"TelecabinSubmode"` // telecabin, telecabin
	AdjacentSites    struct {
		Text    string `xml:",chardata"`
		SiteRef struct {
			Text string `xml:",chardata"`
			Ref  string `xml:"ref,attr"`
		} `xml:"SiteRef"`
	} `xml:"adjacentSites"`
	TramSubmode string `xml:"TramSubmode"` // localTram, localTram, loc...
	PrivateCode string `xml:"PrivateCode"` // 1
}

// Doc
type Doc struct {
	StopPlaces []StopPlace `xml:"StopPlace"`
}
