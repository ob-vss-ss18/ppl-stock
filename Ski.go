package main

type Einsatz string
const  (
	Langlauf Einsatz = "Langlauf"
	Piste Category = "Piste"
	Park  Category = "Park"
)

type Klasse string
const  (
	Skate Klasse = "Skate"
	Allround Klasse = "Allround"
	Twintip Klasse = "Twintip"
)

type Nutzer string
const  (
	Kinder Nutzer = "Kinder"
	KinderJugendliche Nutzer = "Kinder / Jugend"
	Jugend Nutzer = "Jugendliche"
	JugendErwachsen = "Jugend / Erw."
	Erwachsener = "Erwachsene"
)

type Ski struct {
	SkiNr uint32
	ShopNr uint32
	Einsatz Einsatz
	Klasse Klasse
	Nutzer Nutzer
	Gender Gender
	Brand string
	Model string
	Length uint8
	SetPrice float32
	Radius float32
	ArtNr string
	Color string
	SkiPrice float32
	Saison Saison
	BindNr uint32
	BindBrand string
	BindModel string
	BindColor string
	BindVorne uint8
	BindMitte uint8
	BindHinten uint8
	Status Status
	Lagerort Lagerort
	Notes Notes
}