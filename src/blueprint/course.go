package blueprint

const (
	CourseDataEndpoint    = "/course/filter?format=json&location[]=FR&location[]=FR%2FSTG&course[]=bachelor%2Fclassic&scolaryear[]=2019"
	CourseDetailsEndpoint = "/module/%v/%v/%v/?format=json"
)

type CourseSummary struct {
	Id               int      `json:"id"`
	TitleCn          string   `json:"title_cn"`
	Semester         int      `json:"semester"`
	Num              string   `json:"num"`
	Begin            string   `json:"begin"`
	End              string   `json:"end"`
	EndRegister      string   `json:"end_register"`
	Scolaryear       int      `json:"scolaryear"`
	Code             string   `json:"code"`
	Codeinstance     string   `json:"codeinstance"`
	LocationTitle    string   `json:"location_title"`
	InstanceLocation string   `json:"instance_location"`
	Flags            string   `json:"flags"`
	Credits          string   `json:"credits"`
	Rights           []string `json:"rights"`
	Status           string   `json:"status"`
	WaitingGrades    string   `json:"waiting_grades"`
	ActivePromo      string   `json:"active_promo"`
	Open             string   `json:"open"`
	Title            string   `json:"title"`
}

type CourseActivity struct {
	ActivityCode      string      `json:"codeacti"`
	CallIhk           string      `json:"call_ihk"`
	Slug              string      `json:"slug"`
	InstanceLocation  string      `json:"instance_location"`
	ModuleTitle       string      `json:"module_title"`
	Title             string      `json:"title"`
	Description       string      `json:"description"`
	TypeTitle         string      `json:"type_title"`
	TypeCode          string      `json:"type_code"`
	Begin             string      `json:"begin"`
	Start             string      `json:"start"`
	EndRegister       string      `json:"end_register"`
	Deadline          string      `json:"deadline"`
	End               string      `json:"end"`
	NbHours           string      `json:"nb_hours"`
	NbGroup           int         `json:"nb_group"`
	Num               int         `json:"num"`
	Register          string      `json:"register"`
	RegisterByBloc    string      `json:"register_by_bloc"`
	RegisterProf      string      `json:"register_prof"`
	TitleLocationType string      `json:"title_location_type"`
	IsProjet          bool        `json:"is_projet"`
	ProjectId         string      `json:"id_projet"`
	ProjectTitle      string      `json:"project_title"`
	IsNote            bool        `json:"is_note"`
	NbNotes           int         `json:"nb_notes"`
	IsBlocins         bool        `json:"is_blocins"`
	RdvStatus         string      `json:"rdv_status"`
	IDBareme          string      `json:"id_bareme"`
	TitleBareme       string      `json:"title_bareme"`
	Archive           string      `json:"archive"`
	HashElearning     string      `json:"hash_elearning"`
	GedNodeAdm        string      `json:"ged_node_adm"`
	NbPlanified       int         `json:"nb_planified"`
	Hidden            bool        `json:"hidden"`
	Project           ProjectData `json:"project"`
	Events            []struct {
		Code                string      `json:"code"`
		NumEvent            string      `json:"num_event"`
		Seats               string      `json:"seats"`
		Title               string      `json:"title"`
		Description         string      `json:"description"`
		NbInscrits          string      `json:"nb_inscrits"`
		Begin               string      `json:"begin"`
		End                 string      `json:"end"`
		IDActivite          string      `json:"id_activite"`
		Location            string      `json:"location"`
		NbMaxStudentsProjet string      `json:"nb_max_students_projet"`
		AlreadyRegister     string      `json:"already_register"`
		UserStatus          string      `json:"user_status"`
		AllowToken          string      `json:"allow_token"`
		Assistants          []StaffData `json:"assistants"`
	} `json:"events"`
}

type CourseDetails struct {
	Scolaryear         string           `json:"scolaryear"`
	Codemodule         string           `json:"codemodule"`
	Codeinstance       string           `json:"codeinstance"`
	Semester           int              `json:"semester"`
	ScolaryearTemplate string           `json:"scolaryear_template"`
	Title              string           `json:"title"`
	Begin              string           `json:"begin"`
	EndRegister        string           `json:"end_register"`
	End                string           `json:"end"`
	Past               string           `json:"past"`
	Closed             string           `json:"closed"`
	Opened             string           `json:"opened"`
	UserCredits        string           `json:"user_credits"`
	Credits            int              `json:"credits"`
	Description        string           `json:"description"`
	Competence         string           `json:"competence"`
	Flags              string           `json:"flags"`
	InstanceFlags      string           `json:"instance_flags"`
	MaxIns             string           `json:"max_ins"`
	InstanceLocation   string           `json:"instance_location"`
	Hidden             string           `json:"hidden"`
	OldACLBackup       string           `json:"old_acl_backup"`
	Resp               []StaffData      `json:"resp"`
	Assistant          []StaffData      `json:"assistant"`
	Rights             string           `json:"rights"`
	TemplateResp       []StaffData      `json:"template_resp"`
	AllowRegister      int              `json:"allow_register"`
	DateIns            string           `json:"date_ins"`
	StudentRegistered  int              `json:"student_registered"`
	StudentGrade       string           `json:"student_grade"`
	StudentCredits     int              `json:"student_credits"`
	Color              string           `json:"color"`
	StudentFlags       string           `json:"student_flags"`
	CurrentResp        bool             `json:"current_resp"`
	Activites          []CourseActivity `json:"activites"`
}

type Course struct {
	Summary CourseSummary
	Details CourseDetails
}
