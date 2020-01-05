package module

const (
	ModuleEndPoint = "/module/board/?format=json&start=%v&end=%v"
)

type ModuleStruct struct {
	ModuleTitle    string   `json:"title_module"`
	ModuleCode     string   `json:"codemodule"`
	ScholarYear    string   `json:"scolaryear"`
	CodeInstance   string   `json:"codeinstance"`
	CodeLocation   string   `json:"code_location"`
	BeginEvent     string   `json:"begin_event"`
	EndEvent       string   `json:"end_event"`
	Seats          string   `json:"seats"`
	NumEvent       string   `json:"num_event"`
	TypeActivity   string   `json:"type_acti"`
	TypeActiveCode string   `json:"type_acti_code"`
	CodeActive     string   `json:"codeacti"`
	ActivityTitle  string   `json:"acti_title"`
	Num            string   `json:"num"`
	BeginActive    string   `json:"begin_acti"`
	EndActive      string   `json:"end_acti"`
	Registered     int      `json:"registered"`
	SlotInfo       string   `json:"info_creneau"`
	Project        string   `json:"project"`
	Rights         []string `json:"rights"`
}
