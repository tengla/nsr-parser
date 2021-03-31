package main

import "encoding/xml"

// PublicationDelivery was generated 2021-03-31 15:03:22 by thomasen on Thomas-MacBook.
type PublicationDelivery struct {
	XMLName              xml.Name `xml:"PublicationDelivery"`
	Text                 string   `xml:",chardata"`
	Xmlns                string   `xml:"xmlns,attr"`
	Ns2                  string   `xml:"ns2,attr"`
	Ns3                  string   `xml:"ns3,attr"`
	Xsi                  string   `xml:"xsi,attr"`
	Version              string   `xml:"version,attr"`
	SchemaLocation       string   `xml:"schemaLocation,attr"`
	PublicationTimestamp string   `xml:"PublicationTimestamp"` // 2021-03-30T01:52:55.548
	ParticipantRef       string   `xml:"ParticipantRef"`       // NSR
	DataObjects          struct {
		Text      string `xml:",chardata"`
		SiteFrame struct {
			Text          string `xml:",chardata"`
			Modification  string `xml:"modification,attr"`
			Version       string `xml:"version,attr"`
			ID            string `xml:"id,attr"`
			Description   string `xml:"Description"` // Site frame ExportParams{t...
			FrameDefaults struct {
				Text          string `xml:",chardata"`
				DefaultLocale struct {
					Text     string `xml:",chardata"`
					TimeZone string `xml:"TimeZone"` // Europe/Oslo
				} `xml:"DefaultLocale"`
			} `xml:"FrameDefaults"`
			TopographicPlaces struct {
				Text             string `xml:",chardata"`
				TopographicPlace []struct {
					Text         string `xml:",chardata"`
					Created      string `xml:"created,attr"`
					Modification string `xml:"modification,attr"`
					Version      string `xml:"version,attr"`
					ID           string `xml:"id,attr"`
					Changed      string `xml:"changed,attr"`
					ValidBetween struct {
						Text     string `xml:",chardata"`
						FromDate string `xml:"FromDate"` // 2020-01-01T00:00:00, 2020...
						ToDate   string `xml:"ToDate"`   // 2020-01-19T13:15:33.05
					} `xml:"ValidBetween"`
					KeyList struct {
						Text     string `xml:",chardata"`
						KeyValue struct {
							Text  string `xml:",chardata"`
							Key   string `xml:"Key"`   // CHANGED_BY, CHANGED_BY, C...
							Value string `xml:"Value"` // assad.riaz, assad.riaz, a...
						} `xml:"KeyValue"`
					} `xml:"keyList"`
					Polygon struct {
						Text     string `xml:",chardata"`
						ID       string `xml:"id,attr"`
						Exterior struct {
							Text       string `xml:",chardata"`
							LinearRing struct {
								Text    string `xml:",chardata"`
								PosList string `xml:"posList"` // 60.10712230902629 8.93025...
							} `xml:"LinearRing"`
						} `xml:"exterior"`
					} `xml:"Polygon"`
					IsoCode    string `xml:"IsoCode"` // NO-30, SWE, SE-19, NO-46,...
					Descriptor struct {
						Text string `xml:",chardata"`
						Name struct {
							Text string `xml:",chardata"` // Viken, Aremark, Sverige, ...
							Lang string `xml:"lang,attr"`
						} `xml:"Name"`
					} `xml:"Descriptor"`
					TopographicPlaceType string `xml:"TopographicPlaceType"` // county, municipality, cou...
					CountryRef           struct {
						Text string `xml:",chardata"`
						Ref  string `xml:"ref,attr"`
					} `xml:"CountryRef"`
					ParentTopographicPlaceRef struct {
						Text    string `xml:",chardata"`
						Ref     string `xml:"ref,attr"`
						Version string `xml:"version,attr"`
					} `xml:"ParentTopographicPlaceRef"`
				} `xml:"TopographicPlace"`
			} `xml:"topographicPlaces"`
			GroupsOfStopPlaces struct {
				Text              string `xml:",chardata"`
				GroupOfStopPlaces []struct {
					Text         string `xml:",chardata"`
					Created      string `xml:"created,attr"`
					Changed      string `xml:"changed,attr"`
					Modification string `xml:"modification,attr"`
					Version      string `xml:"version,attr"`
					ID           string `xml:"id,attr"`
					KeyList      struct {
						Text     string `xml:",chardata"`
						KeyValue []struct {
							Text  string `xml:",chardata"`
							Key   string `xml:"Key"`   // imported-id, CHANGED_BY, ...
							Value string `xml:"Value"` // johan, johan, johan, joha...
						} `xml:"KeyValue"`
					} `xml:"keyList"`
					Name struct {
						Text string `xml:",chardata"` // Oslo, Asker, Moss, Bergen...
						Lang string `xml:"lang,attr"`
					} `xml:"Name"`
					Description struct {
						Text string `xml:",chardata"` // Dale i Sunnfjord, Sande i...
						Lang string `xml:"lang,attr"`
					} `xml:"Description"`
					Members struct {
						Text         string `xml:",chardata"`
						StopPlaceRef []struct {
							Text string `xml:",chardata"`
							Ref  string `xml:"ref,attr"`
						} `xml:"StopPlaceRef"`
					} `xml:"members"`
					Centroid struct {
						Text     string `xml:",chardata"`
						Location struct {
							Text      string `xml:",chardata"`
							Longitude string `xml:"Longitude"` // 10.747898, 10.434355, 10....
							Latitude  string `xml:"Latitude"`  // 59.911379, 59.833202, 59....
						} `xml:"Location"`
					} `xml:"Centroid"`
				} `xml:"GroupOfStopPlaces"`
			} `xml:"groupsOfStopPlaces"`
			StopPlaces struct {
				Text      string `xml:",chardata"`
				StopPlace []struct {
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
				} `xml:"StopPlace"`
			} `xml:"stopPlaces"`
			Parkings struct {
				Text    string `xml:",chardata"`
				Parking []struct {
					Text         string `xml:",chardata"`
					Created      string `xml:"created,attr"`
					Changed      string `xml:"changed,attr"`
					Modification string `xml:"modification,attr"`
					Version      string `xml:"version,attr"`
					ID           string `xml:"id,attr"`
					KeyList      struct {
						Text     string `xml:",chardata"`
						KeyValue []struct {
							Text  string `xml:",chardata"`
							Key   string `xml:"Key"`   // imported-id, CHANGED_BY, ...
							Value string `xml:"Value"` // NSB:Parking:107bike, moni...
						} `xml:"KeyValue"`
					} `xml:"keyList"`
					Name struct {
						Text string `xml:",chardata"` // Drammen, Leangen, Støren...
						Lang string `xml:"lang,attr"`
					} `xml:"Name"`
					Description string `xml:"Description"` // Med tak., Parkering innef...
					Centroid    struct {
						Text     string `xml:",chardata"`
						Location struct {
							Text      string `xml:",chardata"`
							Longitude string `xml:"Longitude"` // 10.203112, 10.465357, 10....
							Latitude  string `xml:"Latitude"`  // 59.740488, 63.436234, 63....
						} `xml:"Location"`
					} `xml:"Centroid"`
					Covered       string `xml:"Covered"` // covered, outdoors, outdoo...
					ParentSiteRef struct {
						Text    string `xml:",chardata"`
						Ref     string `xml:"ref,attr"`
						Version string `xml:"version,attr"`
					} `xml:"ParentSiteRef"`
					ParkingVehicleTypes   string `xml:"ParkingVehicleTypes"`   // pedalCycle, car, pedalCyc...
					TotalCapacity         string `xml:"TotalCapacity"`         // 0, 21, 0, 75, 0, 23, 0, 0...
					ParkingPaymentProcess string `xml:"ParkingPaymentProcess"` // payByPrepaidToken payByMo...
					ParkingLayout         string `xml:"ParkingLayout"`         // openSpace, openSpace, ope...
					ParkingProperties     struct {
						Text              string `xml:",chardata"`
						ParkingProperties struct {
							Text             string `xml:",chardata"`
							ParkingUserTypes string `xml:"ParkingUserTypes"`
							Spaces           struct {
								Text            string `xml:",chardata"`
								ParkingCapacity []struct {
									Text                            string `xml:",chardata"`
									Version                         string `xml:"version,attr"`
									ID                              string `xml:"id,attr"`
									ParkingUserType                 string `xml:"ParkingUserType"`                 // allUsers, registeredDisab...
									NumberOfSpaces                  string `xml:"NumberOfSpaces"`                  // 20, 1, 46, 4, 21, 2, 10, ...
									NumberOfSpacesWithRechargePoint string `xml:"NumberOfSpacesWithRechargePoint"` // 0, 0, 0, 0, 0, 0, 0, 0, 0...
								} `xml:"ParkingCapacity"`
							} `xml:"spaces"`
						} `xml:"ParkingProperties"`
					} `xml:"parkingProperties"`
					ValidBetween struct {
						Text     string `xml:",chardata"`
						FromDate string `xml:"FromDate"` // 2017-09-18T14:00:15.414, ...
					} `xml:"ValidBetween"`
					RechargingAvailable string `xml:"RechargingAvailable"` // true, true, false, false,...
				} `xml:"Parking"`
			} `xml:"parkings"`
			TariffZones struct {
				Text       string `xml:",chardata"`
				TariffZone []struct {
					Text         string `xml:",chardata"`
					Created      string `xml:"created,attr"`
					Modification string `xml:"modification,attr"`
					Version      string `xml:"version,attr"`
					ID           string `xml:"id,attr"`
					Changed      string `xml:"changed,attr"`
					ValidBetween struct {
						Text     string `xml:",chardata"`
						FromDate string `xml:"FromDate"` // 2019-04-11T10:44:51.17, 2...
					} `xml:"ValidBetween"`
					KeyList struct {
						Text     string `xml:",chardata"`
						KeyValue []struct {
							Text  string `xml:",chardata"`
							Key   string `xml:"Key"`   // imported-id, CHANGED_BY, ...
							Value string `xml:"Value"` // assad.riaz, assad.riaz, a...
						} `xml:"KeyValue"`
					} `xml:"keyList"`
					Name struct {
						Text string `xml:",chardata"` // Snelandia, E2, E1, D, C6,...
						Lang string `xml:"lang,attr"`
					} `xml:"Name"`
					Polygon struct {
						Text     string `xml:",chardata"`
						ID       string `xml:"id,attr"`
						Exterior struct {
							Text       string `xml:",chardata"`
							LinearRing struct {
								Text    string `xml:",chardata"`
								PosList string `xml:"posList"` // 69.0519407 28.929451 69.0...
							} `xml:"LinearRing"`
						} `xml:"exterior"`
					} `xml:"Polygon"`
				} `xml:"TariffZone"`
			} `xml:"tariffZones"`
		} `xml:"SiteFrame"`
	} `xml:"dataObjects"`
}
