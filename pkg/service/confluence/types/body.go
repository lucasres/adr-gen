package types

type Body struct {
	Storage BodyStorage `json:"storage"`
}

type BodyStorage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}
