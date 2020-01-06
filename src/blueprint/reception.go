package blueprint

const (
	ReceptionEndpoint = "/?format=json"
)

type Reception struct {
	IP    string `json:"ip"`
	Board struct {
		Projets []struct {
			Title           string `json:"title"`
			TitleLink       string `json:"title_link"`
			TimelineStart   string `json:"timeline_start"`
			TimelineEnd     string `json:"timeline_end"`
			TimelineBarre   string `json:"timeline_barre"`
			DateInscription bool   `json:"date_inscription"`
			IDActivite      string `json:"id_activite"`
			SoutenanceName  bool   `json:"soutenance_name"`
			SoutenanceLink  bool   `json:"soutenance_link"`
			SoutenanceDate  bool   `json:"soutenance_date"`
			SoutenanceSalle bool   `json:"soutenance_salle"`
		} `json:"projets"`
		Notes     []string `json:"notes"`
		Susies    []string `json:"susies"`
		Activites []struct {
			Title           string      `json:"title"`
			Module          string      `json:"module"`
			ModuleLink      string      `json:"module_link"`
			ModuleCode      string      `json:"module_code"`
			TitleLink       string      `json:"title_link"`
			TimelineStart   string      `json:"timeline_start"`
			TimelineEnd     string      `json:"timeline_end"`
			TimelineBarre   string      `json:"timeline_barre"`
			DateInscription interface{} `json:"date_inscription"`
			Salle           string      `json:"salle"`
			Intervenant     string      `json:"intervenant"`
			Token           string      `json:"token"`
			TokenLink       string      `json:"token_link"`
			RegisterLink    string      `json:"register_link"`
		} `json:"activites"`
		Modules []struct {
			Title           string `json:"title"`
			TitleLink       string `json:"title_link"`
			TimelineStart   string `json:"timeline_start"`
			TimelineEnd     string `json:"timeline_end"`
			TimelineBarre   string `json:"timeline_barre"`
			DateInscription string `json:"date_inscription"`
		} `json:"modules"`
		Stages []struct {
			Company        string `json:"company"`
			Link           string `json:"link"`
			TimelineStart  string `json:"timeline_start"`
			TimelineEnd    string `json:"timeline_end"`
			TimelineBarre  string `json:"timeline_barre"`
			CanNote        bool   `json:"can_note"`
			CompanyCanNote string `json:"company_can_note"`
			Status         string `json:"status"`
			Mandatory      bool   `json:"mandatory"`
		} `json:"stages"`
		Tickets []string `json:"tickets"`
	} `json:"board"`
	History []struct {
		Title string `json:"title"`
		User  struct {
			Picture string `json:"picture"`
			Title   string `json:"title"`
			URL     string `json:"url"`
		} `json:"user"`
		Content    string `json:"content"`
		Date       string `json:"date"`
		ID         string `json:"id"`
		Visible    string `json:"visible"`
		IDActivite string `json:"id_activite"`
		Class      string `json:"class"`
	} `json:"history"`
	Infos struct {
		Location string `json:"location"`
	} `json:"infos"`
	Current []struct {
		CreditsMin   string `json:"credits_min"`
		CreditsNorm  string `json:"credits_norm"`
		CreditsObj   string `json:"credits_obj"`
		NslogMin     string `json:"nslog_min"`
		NslogNorm    string `json:"nslog_norm"`
		Credits      string `json:"credits"`
		Grade        string `json:"grade"`
		Cycle        string `json:"cycle"`
		CodeModule   string `json:"code_module"`
		CurrentCycle string `json:"current_cycle"`
		SemesterCode string `json:"semester_code"`
		SemesterNum  string `json:"semester_num"`
		ActiveLog    string `json:"active_log"`
	} `json:"current"`
}

type Credits struct {
	Minimum   int
	Aimed     int
	Objective int
}
