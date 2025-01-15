package user_request

type UserRequest struct {
	Preference string `json:"preference" validate:"required,oneof=CARDIO WEIGHT"`
	WeightUnit string `json:"weightUnit" validate:"required,oneof=KG LBS"`
	HeightUnit string `json:"heightUnit" validate:"required,oneof=CM INCH"`
	Weight     int    `json:"weight" validate:"required,min=10,max=1000"`
	Height     int    `json:"height" validate:"required,min=3,max=250"`
	Name       string `json:"name" validate:"omitempty,min=2,max=60"`
	ImageUri   string `json:"imageUri" validate:"omitempty,uri"`
}


