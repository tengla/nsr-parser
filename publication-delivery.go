package main

import "encoding/xml"

// PublicationDelivery was generated 2021-03-31 15:41:46 by thomasen on Thomas-MacBook.
type PublicationDelivery struct {
	XMLName              xml.Name `xml:"PublicationDelivery" json:"publicationdelivery,omitempty"`
	Xmlns                string   `xml:"xmlns,attr" json:"xmlns,omitempty"`
	Ns2                  string   `xml:"ns2,attr" json:"ns2,omitempty"`
	Ns3                  string   `xml:"ns3,attr" json:"ns3,omitempty"`
	Xsi                  string   `xml:"xsi,attr" json:"xsi,omitempty"`
	Version              string   `xml:"version,attr" json:"version,omitempty"`
	SchemaLocation       string   `xml:"schemaLocation,attr" json:"schemalocation,omitempty"`
	PublicationTimestamp string   `xml:"PublicationTimestamp"` // 2021-03-30T01:52:55.548
	ParticipantRef       string   `xml:"ParticipantRef"`       // NSR
	DataObjects          struct {
		SiteFrame struct {
			Modification  string `xml:"modification,attr" json:"modification,omitempty"`
			Version       string `xml:"version,attr" json:"version,omitempty"`
			ID            string `xml:"id,attr" json:"id,omitempty"`
			Description   string `xml:"Description"` // Site frame ExportParams{t...
			FrameDefaults struct {
				DefaultLocale struct {
					TimeZone string `xml:"TimeZone"` // Europe/Oslo
				} `xml:"DefaultLocale" json:"defaultlocale,omitempty"`
			} `xml:"FrameDefaults" json:"framedefaults,omitempty"`
			TopographicPlaces struct {
				TopographicPlace []struct {
					Created      string `xml:"created,attr" json:"created,omitempty"`
					Modification string `xml:"modification,attr" json:"modification,omitempty"`
					Version      string `xml:"version,attr" json:"version,omitempty"`
					ID           string `xml:"id,attr" json:"id,omitempty"`
					Changed      string `xml:"changed,attr" json:"changed,omitempty"`
					ValidBetween struct {
						FromDate string `xml:"FromDate"` // 2020-01-01T00:00:00, 2020...
						ToDate   string `xml:"ToDate"`   // 2020-01-19T13:15:33.05
					} `xml:"ValidBetween" json:"validbetween,omitempty"`
					KeyList struct {
						KeyValue struct {
							Key   string `xml:"Key"`   // CHANGED_BY, CHANGED_BY, C...
							Value string `xml:"Value"` // assad.riaz, assad.riaz, a...
						} `xml:"KeyValue" json:"keyvalue,omitempty"`
					} `xml:"keyList" json:"keylist,omitempty"`
					Polygon struct {
						ID       string `xml:"id,attr" json:"id,omitempty"`
						Exterior struct {
							LinearRing struct {
								PosList string `xml:"posList"` // 60.10712230902629 8.93025...
							} `xml:"LinearRing" json:"linearring,omitempty"`
						} `xml:"exterior" json:"exterior,omitempty"`
					} `xml:"Polygon" json:"polygon,omitempty"`
					IsoCode    string `xml:"IsoCode"` // NO-30, SWE, SE-19, NO-46,...
					Descriptor struct {
						Name struct {
							Text string `xml:",chardata" json:"text,omitempty"` // Viken, Aremark, Sverige, ...
							Lang string `xml:"lang,attr" json:"lang,omitempty"`
						} `xml:"Name" json:"name,omitempty"`
					} `xml:"Descriptor" json:"descriptor,omitempty"`
					TopographicPlaceType string `xml:"TopographicPlaceType"` // county, municipality, cou...
					CountryRef           struct {
						Ref string `xml:"ref,attr" json:"ref,omitempty"`
					} `xml:"CountryRef" json:"countryref,omitempty"`
					ParentTopographicPlaceRef struct {
						Ref     string `xml:"ref,attr" json:"ref,omitempty"`
						Version string `xml:"version,attr" json:"version,omitempty"`
					} `xml:"ParentTopographicPlaceRef" json:"parenttopographicplaceref,omitempty"`
				} `xml:"TopographicPlace" json:"topographicplace,omitempty"`
			} `xml:"topographicPlaces" json:"topographicplaces,omitempty"`
			GroupsOfStopPlaces struct {
				GroupOfStopPlaces []struct {
					Created      string `xml:"created,attr" json:"created,omitempty"`
					Changed      string `xml:"changed,attr" json:"changed,omitempty"`
					Modification string `xml:"modification,attr" json:"modification,omitempty"`
					Version      string `xml:"version,attr" json:"version,omitempty"`
					ID           string `xml:"id,attr" json:"id,omitempty"`
					KeyList      struct {
						KeyValue []struct {
							Key   string `xml:"Key"`   // imported-id, CHANGED_BY, ...
							Value string `xml:"Value"` // johan, johan, johan, joha...
						} `xml:"KeyValue" json:"keyvalue,omitempty"`
					} `xml:"keyList" json:"keylist,omitempty"`
					Name struct {
						Text string `xml:",chardata" json:"text,omitempty"` // Oslo, Asker, Moss, Bergen...
						Lang string `xml:"lang,attr" json:"lang,omitempty"`
					} `xml:"Name" json:"name,omitempty"`
					Description struct {
						Text string `xml:",chardata" json:"text,omitempty"` // Dale i Sunnfjord, Sande i...
						Lang string `xml:"lang,attr" json:"lang,omitempty"`
					} `xml:"Description" json:"description,omitempty"`
					Members struct {
						StopPlaceRef []struct {
							Ref string `xml:"ref,attr" json:"ref,omitempty"`
						} `xml:"StopPlaceRef" json:"stopplaceref,omitempty"`
					} `xml:"members" json:"members,omitempty"`
					Centroid struct {
						Location struct {
							Longitude string `xml:"Longitude"` // 10.747898, 10.434355, 10....
							Latitude  string `xml:"Latitude"`  // 59.911379, 59.833202, 59....
						} `xml:"Location" json:"location,omitempty"`
					} `xml:"Centroid" json:"centroid,omitempty"`
				} `xml:"GroupOfStopPlaces" json:"groupofstopplaces,omitempty"`
			} `xml:"groupsOfStopPlaces" json:"groupsofstopplaces,omitempty"`
			StopPlaces struct {
				StopPlace []struct {
					Created      string `xml:"created,attr" json:"created,omitempty"`
					Changed      string `xml:"changed,attr" json:"changed,omitempty"`
					Modification string `xml:"modification,attr" json:"modification,omitempty"`
					Version      string `xml:"version,attr" json:"version,omitempty"`
					ID           string `xml:"id,attr" json:"id,omitempty"`
					ValidBetween struct {
						FromDate string `xml:"FromDate"` // 2017-06-19T19:12:29.887, ...
						ToDate   string `xml:"ToDate"`   // 2021-05-02T01:00:09, 2021...
					} `xml:"ValidBetween" json:"validbetween,omitempty"`
					KeyList struct {
						KeyValue []struct {
							Key   string `xml:"Key"`   // jbvCode, uicCode, iffCode...
							Value string `xml:"Value"` // GU, 7601213, 7601213, 121...
						} `xml:"KeyValue" json:"keyvalue,omitempty"`
					} `xml:"keyList" json:"keylist,omitempty"`
					Name struct {
						Text string `xml:",chardata" json:"text,omitempty"` // Gudå, Eikenes stasjon, H...
						Lang string `xml:"lang,attr" json:"lang,omitempty"`
					} `xml:"Name" json:"name,omitempty"`
					Centroid struct {
						Location struct {
							Longitude string `xml:"Longitude"` // 11.620069, 9.895932, 11.1...
							Latitude  string `xml:"Latitude"`  // 63.444097, 59.146811, 62....
						} `xml:"Location" json:"location,omitempty"`
					} `xml:"Centroid" json:"centroid,omitempty"`
					AccessibilityAssessment struct {
						Modification           string `xml:"modification,attr" json:"modification,omitempty"`
						Version                string `xml:"version,attr" json:"version,omitempty"`
						ID                     string `xml:"id,attr" json:"id,omitempty"`
						MobilityImpairedAccess string `xml:"MobilityImpairedAccess"` // partial, partial, partial...
						Limitations            struct {
							AccessibilityLimitation struct {
								Modification            string `xml:"modification,attr" json:"modification,omitempty"`
								Version                 string `xml:"version,attr" json:"version,omitempty"`
								ID                      string `xml:"id,attr" json:"id,omitempty"`
								WheelchairAccess        string `xml:"WheelchairAccess"`        // true, false, true, true, ...
								StepFreeAccess          string `xml:"StepFreeAccess"`          // true, false, true, true, ...
								EscalatorFreeAccess     string `xml:"EscalatorFreeAccess"`     // unknown, unknown, unknown...
								LiftFreeAccess          string `xml:"LiftFreeAccess"`          // unknown, unknown, unknown...
								AudibleSignalsAvailable string `xml:"AudibleSignalsAvailable"` // unknown, unknown, unknown...
							} `xml:"AccessibilityLimitation" json:"accessibilitylimitation,omitempty"`
						} `xml:"limitations" json:"limitations,omitempty"`
					} `xml:"AccessibilityAssessment" json:"accessibilityassessment,omitempty"`
					TopographicPlaceRef struct {
						Ref     string `xml:"ref,attr" json:"ref,omitempty"`
						Created string `xml:"created,attr" json:"created,omitempty"`
						Version string `xml:"version,attr" json:"version,omitempty"`
					} `xml:"TopographicPlaceRef" json:"topographicplaceref,omitempty"`
					TransportMode       string `xml:"TransportMode"` // bus, rail, bus, bus, bus,...
					BusSubmode          string `xml:"BusSubmode"`    // railReplacementBus, railR...
					OtherTransportModes string `xml:"OtherTransportModes"`
					TariffZones         struct {
						TariffZoneRef []struct {
							Ref     string `xml:"ref,attr" json:"ref,omitempty"`
							Version string `xml:"version,attr" json:"version,omitempty"`
						} `xml:"TariffZoneRef" json:"tariffzoneref,omitempty"`
					} `xml:"tariffZones" json:"tariffzones,omitempty"`
					StopPlaceType string `xml:"StopPlaceType"` // onstreetBus, railStation,...
					Weighting     string `xml:"Weighting"`     // interchangeAllowed, inter...
					Quays         struct {
						Quay []struct {
							Modification string `xml:"modification,attr" json:"modification,omitempty"`
							Version      string `xml:"version,attr" json:"version,omitempty"`
							ID           string `xml:"id,attr" json:"id,omitempty"`
							Changed      string `xml:"changed,attr" json:"changed,omitempty"`
							Created      string `xml:"created,attr" json:"created,omitempty"`
							KeyList      struct {
								KeyValue []struct {
									Key   string `xml:"Key"`   // uicCode, imported-id, gra...
									Value string `xml:"Value"` // 7601213, NSB:Quay:7601213...
								} `xml:"KeyValue" json:"keyvalue,omitempty"`
							} `xml:"keyList" json:"keylist,omitempty"`
							Centroid struct {
								Location struct {
									Longitude string `xml:"Longitude"` // 11.620069, 9.895932, 11.1...
									Latitude  string `xml:"Latitude"`  // 63.444097, 59.146811, 62....
								} `xml:"Location" json:"location,omitempty"`
							} `xml:"Centroid" json:"centroid,omitempty"`
							OtherTransportModes string `xml:"OtherTransportModes"`
							PrivateCode         struct {
								Text string `xml:",chardata" json:"text,omitempty"` // 1, 1, 1, 1, 1, 1, 01, 01,...
								Type string `xml:"type,attr" json:"type,omitempty"`
							} `xml:"PrivateCode" json:"privatecode,omitempty"`
							PublicCode     string `xml:"PublicCode"`     // 1, 1, 1, 1, 1, 1, O, 2, 1...
							CompassBearing string `xml:"CompassBearing"` // 201.0, 59.0, 34.0, 9.0, 3...
							Name           struct {
								Text string `xml:",chardata" json:"text,omitempty"` // Bukkholmveien 1, Ise spor...
								Lang string `xml:"lang,attr" json:"lang,omitempty"`
							} `xml:"Name" json:"name,omitempty"`
							Description struct {
								Text string `xml:",chardata" json:"text,omitempty"` // Retning Frognerseteren, R...
								Lang string `xml:"lang,attr" json:"lang,omitempty"`
							} `xml:"Description" json:"description,omitempty"`
							PlaceEquipments struct {
								ID          string `xml:"id,attr" json:"id,omitempty"`
								GeneralSign []struct {
									Modification    string `xml:"modification,attr" json:"modification,omitempty"`
									Version         string `xml:"version,attr" json:"version,omitempty"`
									ID              string `xml:"id,attr" json:"id,omitempty"`
									PrivateCode     string `xml:"PrivateCode"`     // 512, 512, 512, 512, 512, ...
									SignContentType string `xml:"SignContentType"` // transportMode, transportM...
									Content         struct {
										Text string `xml:",chardata" json:"text,omitempty"` // Timetable, Timetable, Tim...
										Lang string `xml:"lang,attr" json:"lang,omitempty"`
									} `xml:"Content" json:"content,omitempty"`
								} `xml:"GeneralSign" json:"generalsign,omitempty"`
								ShelterEquipment []struct {
									Modification string `xml:"modification,attr" json:"modification,omitempty"`
									Version      string `xml:"version,attr" json:"version,omitempty"`
									ID           string `xml:"id,attr" json:"id,omitempty"`
									Enclosed     string `xml:"Enclosed"` // true, true, true, true, t...
									Seats        string `xml:"Seats"`    // 1, 1, 1, 1, 1, 1, 1, 1, 1...
									StepFree     string `xml:"StepFree"` // false, false, false, fals...
								} `xml:"ShelterEquipment" json:"shelterequipment,omitempty"`
								WaitingRoomEquipment []struct {
									Modification string `xml:"modification,attr" json:"modification,omitempty"`
									Version      string `xml:"version,attr" json:"version,omitempty"`
									ID           string `xml:"id,attr" json:"id,omitempty"`
									Sanitary     string `xml:"Sanitary"`
									Seats        string `xml:"Seats"`    // 1, 0
									StepFree     string `xml:"StepFree"` // false, false
									Heated       string `xml:"Heated"`   // false, false
								} `xml:"WaitingRoomEquipment" json:"waitingroomequipment,omitempty"`
								SanitaryEquipment struct {
									Modification         string `xml:"modification,attr" json:"modification,omitempty"`
									Version              string `xml:"version,attr" json:"version,omitempty"`
									ID                   string `xml:"id,attr" json:"id,omitempty"`
									Gender               string `xml:"Gender"` // both, both, both, both, b...
									SanitaryFacilityList string `xml:"SanitaryFacilityList"`
									PaymentMethods       string `xml:"PaymentMethods"`
									NumberOfToilets      string `xml:"NumberOfToilets"` // 1
								} `xml:"SanitaryEquipment" json:"sanitaryequipment,omitempty"`
								TicketingEquipment struct {
									Modification            string `xml:"modification,attr" json:"modification,omitempty"`
									Version                 string `xml:"version,attr" json:"version,omitempty"`
									ID                      string `xml:"id,attr" json:"id,omitempty"`
									VehicleModes            string `xml:"VehicleModes"`
									TicketMachines          string `xml:"TicketMachines"`   // true, true, true, true, t...
									NumberOfMachines        string `xml:"NumberOfMachines"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
									TicketingFacilityList   string `xml:"TicketingFacilityList"`
									TicketOffice            string `xml:"TicketOffice"` // true, true, true, true, t...
									PaymentMethods          string `xml:"PaymentMethods"`
									TicketTypesAvailable    string `xml:"TicketTypesAvailable"`
									ScopeOfTicketsAvailable string `xml:"ScopeOfTicketsAvailable"`
								} `xml:"TicketingEquipment" json:"ticketingequipment,omitempty"`
							} `xml:"placeEquipments" json:"placeequipments,omitempty"`
							AccessibilityAssessment struct {
								Modification           string `xml:"modification,attr" json:"modification,omitempty"`
								Version                string `xml:"version,attr" json:"version,omitempty"`
								ID                     string `xml:"id,attr" json:"id,omitempty"`
								MobilityImpairedAccess string `xml:"MobilityImpairedAccess"` // partial, partial, partial...
								Limitations            struct {
									AccessibilityLimitation struct {
										Modification            string `xml:"modification,attr" json:"modification,omitempty"`
										Version                 string `xml:"version,attr" json:"version,omitempty"`
										ID                      string `xml:"id,attr" json:"id,omitempty"`
										WheelchairAccess        string `xml:"WheelchairAccess"`        // false, false, false, true...
										StepFreeAccess          string `xml:"StepFreeAccess"`          // false, true, false, true,...
										EscalatorFreeAccess     string `xml:"EscalatorFreeAccess"`     // unknown, unknown, unknown...
										LiftFreeAccess          string `xml:"LiftFreeAccess"`          // unknown, unknown, unknown...
										AudibleSignalsAvailable string `xml:"AudibleSignalsAvailable"` // unknown, unknown, unknown...
									} `xml:"AccessibilityLimitation" json:"accessibilitylimitation,omitempty"`
								} `xml:"limitations" json:"limitations,omitempty"`
							} `xml:"AccessibilityAssessment" json:"accessibilityassessment,omitempty"`
							ValidBetween struct {
								FromDate string `xml:"FromDate"` // 2017-06-21T17:01:09.146, ...
							} `xml:"ValidBetween" json:"validbetween,omitempty"`
						} `xml:"Quay" json:"quay,omitempty"`
					} `xml:"quays" json:"quays,omitempty"`
					PlaceEquipments struct {
						ID               string `xml:"id,attr" json:"id,omitempty"`
						ShelterEquipment struct {
							Modification string `xml:"modification,attr" json:"modification,omitempty"`
							Version      string `xml:"version,attr" json:"version,omitempty"`
							ID           string `xml:"id,attr" json:"id,omitempty"`
							Enclosed     string `xml:"Enclosed"` // true, true, false, false,...
							Seats        string `xml:"Seats"`    // 1, 1, 1, 1, 1, 1, 1, 1, 1...
							StepFree     string `xml:"StepFree"` // false, false, false, fals...
						} `xml:"ShelterEquipment" json:"shelterequipment,omitempty"`
						WaitingRoomEquipment struct {
							Modification string `xml:"modification,attr" json:"modification,omitempty"`
							Version      string `xml:"version,attr" json:"version,omitempty"`
							ID           string `xml:"id,attr" json:"id,omitempty"`
							Sanitary     string `xml:"Sanitary"`
							Seats        string `xml:"Seats"`    // 1, 1, 1, 1, 1, 1, 1, 1, 1...
							StepFree     string `xml:"StepFree"` // false, false, false, fals...
							Heated       string `xml:"Heated"`   // false, false, false, fals...
						} `xml:"WaitingRoomEquipment" json:"waitingroomequipment,omitempty"`
						TicketingEquipment struct {
							Modification            string `xml:"modification,attr" json:"modification,omitempty"`
							Version                 string `xml:"version,attr" json:"version,omitempty"`
							ID                      string `xml:"id,attr" json:"id,omitempty"`
							VehicleModes            string `xml:"VehicleModes"`
							TicketMachines          string `xml:"TicketMachines"`   // true, true, true, true, t...
							NumberOfMachines        string `xml:"NumberOfMachines"` // 1, 1, 1, 1, 1, 1, 2, 2, 1...
							TicketingFacilityList   string `xml:"TicketingFacilityList"`
							TicketOffice            string `xml:"TicketOffice"` // false, false, false, fals...
							PaymentMethods          string `xml:"PaymentMethods"`
							TicketTypesAvailable    string `xml:"TicketTypesAvailable"`
							ScopeOfTicketsAvailable string `xml:"ScopeOfTicketsAvailable"`
						} `xml:"TicketingEquipment" json:"ticketingequipment,omitempty"`
						GeneralSign struct {
							Modification    string `xml:"modification,attr" json:"modification,omitempty"`
							Version         string `xml:"version,attr" json:"version,omitempty"`
							ID              string `xml:"id,attr" json:"id,omitempty"`
							PrivateCode     string `xml:"PrivateCode"`     // 512, 512, 512, 512, 512, ...
							SignContentType string `xml:"SignContentType"` // transportMode, transportM...
						} `xml:"GeneralSign" json:"generalsign,omitempty"`
						SanitaryEquipment struct {
							Modification         string `xml:"modification,attr" json:"modification,omitempty"`
							Version              string `xml:"version,attr" json:"version,omitempty"`
							ID                   string `xml:"id,attr" json:"id,omitempty"`
							Gender               string `xml:"Gender"` // both, both, both, both, b...
							SanitaryFacilityList string `xml:"SanitaryFacilityList"`
							PaymentMethods       string `xml:"PaymentMethods"`
							NumberOfToilets      string `xml:"NumberOfToilets"` // 1, 1, 1, 1, 1, 1, 1, 1, 1...
						} `xml:"SanitaryEquipment" json:"sanitaryequipment,omitempty"`
					} `xml:"placeEquipments" json:"placeequipments,omitempty"`
					Description struct {
						Text string `xml:",chardata" json:"text,omitempty"` // skole, Oppøyen, i Sognad...
						Lang string `xml:"lang,attr" json:"lang,omitempty"`
					} `xml:"Description" json:"description,omitempty"`
					AlternativeNames struct {
						AlternativeName []struct {
							Modification string `xml:"modification,attr" json:"modification,omitempty"`
							Version      string `xml:"version,attr" json:"version,omitempty"`
							ID           string `xml:"id,attr" json:"id,omitempty"`
							NameType     string `xml:"NameType"` // alias, alias, alias, alia...
							Name         struct {
								Text string `xml:",chardata" json:"text,omitempty"` // Hoelsgård hotel, Nysvemo...
								Lang string `xml:"lang,attr" json:"lang,omitempty"`
							} `xml:"Name" json:"name,omitempty"`
						} `xml:"AlternativeName" json:"alternativename,omitempty"`
					} `xml:"alternativeNames" json:"alternativenames,omitempty"`
					RailSubmode   string `xml:"RailSubmode"`  // touristRailway, touristRa...
					WaterSubmode  string `xml:"WaterSubmode"` // localCarFerry, localPasse...
					MetroSubmode  string `xml:"MetroSubmode"` // metro, metro, metro, metr...
					ParentSiteRef struct {
						Ref     string `xml:"ref,attr" json:"ref,omitempty"`
						Version string `xml:"version,attr" json:"version,omitempty"`
					} `xml:"ParentSiteRef" json:"parentsiteref,omitempty"`
					AirSubmode       string `xml:"AirSubmode"`       // helicopterService, unknow...
					TelecabinSubmode string `xml:"TelecabinSubmode"` // telecabin, telecabin
					AdjacentSites    struct {
						SiteRef struct {
							Ref string `xml:"ref,attr" json:"ref,omitempty"`
						} `xml:"SiteRef" json:"siteref,omitempty"`
					} `xml:"adjacentSites" json:"adjacentsites,omitempty"`
					TramSubmode string `xml:"TramSubmode"` // localTram, localTram, loc...
					PrivateCode string `xml:"PrivateCode"` // 1
				} `xml:"StopPlace" json:"stopplace,omitempty"`
			} `xml:"stopPlaces" json:"stopplaces,omitempty"`
			Parkings struct {
				Parking []struct {
					Created      string `xml:"created,attr" json:"created,omitempty"`
					Changed      string `xml:"changed,attr" json:"changed,omitempty"`
					Modification string `xml:"modification,attr" json:"modification,omitempty"`
					Version      string `xml:"version,attr" json:"version,omitempty"`
					ID           string `xml:"id,attr" json:"id,omitempty"`
					KeyList      struct {
						KeyValue []struct {
							Key   string `xml:"Key"`   // imported-id, CHANGED_BY, ...
							Value string `xml:"Value"` // NSB:Parking:107bike, moni...
						} `xml:"KeyValue" json:"keyvalue,omitempty"`
					} `xml:"keyList" json:"keylist,omitempty"`
					Name struct {
						Text string `xml:",chardata" json:"text,omitempty"` // Drammen, Leangen, Støren...
						Lang string `xml:"lang,attr" json:"lang,omitempty"`
					} `xml:"Name" json:"name,omitempty"`
					Description string `xml:"Description"` // Med tak., Parkering innef...
					Centroid    struct {
						Location struct {
							Longitude string `xml:"Longitude"` // 10.203112, 10.465357, 10....
							Latitude  string `xml:"Latitude"`  // 59.740488, 63.436234, 63....
						} `xml:"Location" json:"location,omitempty"`
					} `xml:"Centroid" json:"centroid,omitempty"`
					Covered       string `xml:"Covered"` // covered, outdoors, outdoo...
					ParentSiteRef struct {
						Ref     string `xml:"ref,attr" json:"ref,omitempty"`
						Version string `xml:"version,attr" json:"version,omitempty"`
					} `xml:"ParentSiteRef" json:"parentsiteref,omitempty"`
					ParkingVehicleTypes   string `xml:"ParkingVehicleTypes"`   // pedalCycle, car, pedalCyc...
					TotalCapacity         string `xml:"TotalCapacity"`         // 0, 21, 0, 75, 0, 23, 0, 0...
					ParkingPaymentProcess string `xml:"ParkingPaymentProcess"` // payByPrepaidToken payByMo...
					ParkingLayout         string `xml:"ParkingLayout"`         // openSpace, openSpace, ope...
					ParkingProperties     struct {
						ParkingProperties struct {
							ParkingUserTypes string `xml:"ParkingUserTypes"`
							Spaces           struct {
								ParkingCapacity []struct {
									Version                         string `xml:"version,attr" json:"version,omitempty"`
									ID                              string `xml:"id,attr" json:"id,omitempty"`
									ParkingUserType                 string `xml:"ParkingUserType"`                 // allUsers, registeredDisab...
									NumberOfSpaces                  string `xml:"NumberOfSpaces"`                  // 20, 1, 46, 4, 21, 2, 10, ...
									NumberOfSpacesWithRechargePoint string `xml:"NumberOfSpacesWithRechargePoint"` // 0, 0, 0, 0, 0, 0, 0, 0, 0...
								} `xml:"ParkingCapacity" json:"parkingcapacity,omitempty"`
							} `xml:"spaces" json:"spaces,omitempty"`
						} `xml:"ParkingProperties" json:"parkingproperties,omitempty"`
					} `xml:"parkingProperties" json:"parkingproperties,omitempty"`
					ValidBetween struct {
						FromDate string `xml:"FromDate"` // 2017-09-18T14:00:15.414, ...
					} `xml:"ValidBetween" json:"validbetween,omitempty"`
					RechargingAvailable string `xml:"RechargingAvailable"` // true, true, false, false,...
				} `xml:"Parking" json:"parking,omitempty"`
			} `xml:"parkings" json:"parkings,omitempty"`
			TariffZones struct {
				TariffZone []struct {
					Created      string `xml:"created,attr" json:"created,omitempty"`
					Modification string `xml:"modification,attr" json:"modification,omitempty"`
					Version      string `xml:"version,attr" json:"version,omitempty"`
					ID           string `xml:"id,attr" json:"id,omitempty"`
					Changed      string `xml:"changed,attr" json:"changed,omitempty"`
					ValidBetween struct {
						FromDate string `xml:"FromDate"` // 2019-04-11T10:44:51.17, 2...
					} `xml:"ValidBetween" json:"validbetween,omitempty"`
					KeyList struct {
						KeyValue []struct {
							Key   string `xml:"Key"`   // imported-id, CHANGED_BY, ...
							Value string `xml:"Value"` // assad.riaz, assad.riaz, a...
						} `xml:"KeyValue" json:"keyvalue,omitempty"`
					} `xml:"keyList" json:"keylist,omitempty"`
					Name struct {
						Text string `xml:",chardata" json:"text,omitempty"` // Snelandia, E2, E1, D, C6,...
						Lang string `xml:"lang,attr" json:"lang,omitempty"`
					} `xml:"Name" json:"name,omitempty"`
					Polygon struct {
						ID       string `xml:"id,attr" json:"id,omitempty"`
						Exterior struct {
							LinearRing struct {
								PosList string `xml:"posList"` // 69.0519407 28.929451 69.0...
							} `xml:"LinearRing" json:"linearring,omitempty"`
						} `xml:"exterior" json:"exterior,omitempty"`
					} `xml:"Polygon" json:"polygon,omitempty"`
				} `xml:"TariffZone" json:"tariffzone,omitempty"`
			} `xml:"tariffZones" json:"tariffzones,omitempty"`
		} `xml:"SiteFrame" json:"siteframe,omitempty"`
	} `xml:"dataObjects" json:"dataobjects,omitempty"`
}
