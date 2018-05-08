package main

type Category string
const  (
	Children Category = "Kinder"
	Teenager Category = "Jugendliche"
	Adults   Category = "Erwachsene"
	TeenAndAdults Category = "Jugend/Erw."
)

type Art string
const  (
	Alpin Art = "Aplin"
)

type Gender string
const  (
	Male Gender = "Männlich"
	Female Gender = "Weiblich"
	Uni Gender = "Uni"
)

type Saison string
const  (
	None Saison = "o. A."
)

type Status string
const  (
	Used Status = "gebr."
	New Status = "neu"
)

type Lagerort string
const  (
	Rent Lagerort = "Verleih"
	Shop Lagerort = "Shop"
)

type Notes string
const  (
	Sold Notes = "verkauft"
	Broken Notes = "defekt"
)

type Skistock struct {
	StockNr uint32
	Category Category
	Art Art
	Gender Gender
	Brand string
	Modell string
	Length uint8
	Price float32
	Color string
	Saison Saison
	Status Status
	Lagerort Lagerort
	Notes Notes
}