package shom

type Wfs struct {
	Type     string       `json:"type"`
	Features []WfsFeature `json:"features"`
}

type WfsFeature struct {
	Type         string               `json:"type"`
	Id           string               `json:"id"`
	Geometry     WfsFeatureGeometry   `json:"geometry"`
	GeometryName string               `json:"geometry_name"`
	Properties   WfsFeatureProperties `json:"properties"`
}

type WfsFeatureGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type WfsFeatureProperties struct {
	Cst      string  `json:"cst"`
	Toponyme string  `json:"toponyme"`
	Ut       int     `json:"ut"`
	UtSup    int     `json:"ut_sup"`
	BaieDeSe int     `json:"baie_de_se"`
	Pays     string  `json:"pays"`
	Coeff    int     `json:"coeff"`
	Official int     `json:"official"`
	Nota     int     `json:"nota"`
	Cache    int     `json:"cache"`
	HLegale  int     `json:"h_legale"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

// Annuaire de marees
// {"2022-03-22":[["tide.low","04:00","1.37","---"],["tide.high","09:39","11.44","91"],["tide.low","16:19","1.60","---"],["tide.none","--:--","---","---"]]}
type AnnuaireMarees map[string][][]string

// Hauteur d'eau par heure
// {"2022-03-22":[["00:00:00",7.67],["00:05:00",7.47],...]}
type HauteurEauParHeure map[string][][]string

// Grandes marees
// [[["82","88"],["94","98"],...]}
type Coeff [][][]string
