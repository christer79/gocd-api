package gocdapi

type Coordinate struct {
	PipelineGroup string
	PipelineName  string
	PipelineCount int
	StageName     string
	StageCount    int
	JobName       string
	FilePath      string
}

type MaterialModification struct {
	EmailAddress     string `json:"email_address"`
	Id               int    `json:"id"`
	ModificationTime int    `json:"modified_time"`
	UserName         string `json:"user_name"`
	Comment          string `json:"comment"`
	Revision         string `json:"revision"`
}

type MaterialRevision struct {
	MaterialModifications []MaterialModification `json:"modifications"`
	Material              Material               `json:"material"`
	Changed               bool                   `json:"changed"`
}

type BuildCause struct {
	Approver          string             `json:"approver"`
	MaterialRevisions []MaterialRevision `json:"material_revisions"`

	TriggerForced  bool   `json:"trigger_forced"`
	TriggerMessage string `json:"trigger_message"`
}

type PipelineInstance struct {
	BuildCause          BuildCause `json:"build_cause"`
	Name                string     `json:"name"`
	NaturalOrder        string     `json:natural_order`
	CanRun              bool       `json:"can_run"`
	Comment             string     `json:"comment"`
	Stages              []Stage    `json:"stages"`
	Counter             int        `json:"counter"`
	Id                  int        `json:"id"`
	PreparingToSchedule bool       `json:"reparing_to_schedule"`
	Label               string     `json:"label"`
}

type Pagination struct {
	Offsett  int `json:"offset"`
	Total    int `json:"total"`
	PageSize int `json:"page_size"`
}

type PipelineHistory struct {
	Pipelines  []PipelineInstance `json:"pipelines"`
	Pagination Pagination         `json:"pagination"`
}
type PipelineStatus struct {
	Locked      bool `json:"locked"`
	Paosed      bool `json:"paused"`
	Schedulable bool `json:"schedulable"`
}

//Agent represents a Agent index as in the response when listing agents
type Agent struct {
	Os           string   `json:"operating_system"`
	Environments []string `json:"environments"`
	UUID         string   `json:"uuid"`
	AgentName    string   `json:"agent_name"`
	Resources    []string `json:"resources"`
	Sandbox      string   `json:"sandbox"`
	Status       string   `json:"status"`
	BuildLocator string   `json:"build_locator"`
	IPAddress    string   `json:"ip_address"`
	Enabled      bool     `json:"enabled"`
}

type Embedded struct {
	Agents []Agent `json:"agents"`
}

type AgentList struct {
	Embedded Embedded `json:"_embedded"`
}

type Message struct {
	Message string `json:"message"`
}

type File struct {
	Name string `json:"name"`
	Url  string `json:"url"`
	Type string `json:"type"`
}

type Artifact struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Type  string `json:"type"`
	Files []File `json:"files"`
}

type Job struct {
	Name          string `json:"name"`
	Result        string `json:"result"`
	State         string `json:"state"`
	Id            int    `json:"id"`
	ScheduledDate int    `json:"scheduled_date"`
}

//Stage holds informaton about a Stage of a Pipeline
type Stage struct {
	Name             string `json:"name"`
	ApprovedBy       string `json:""`
	Jobs             []Job  `json:"job"`
	CanRun           bool   `json:"can_run"`
	Result           string `json:"result"`
	ApprovalType     string `json:"approval_type"`
	Counter          string `json:"counter"`
	Id               int    `json:"id"`
	OperatePermisson bool   `json:"operate_permission"`
	RerunOfCounter   string `json:"rerun_of_counter"`
	Scheduled        bool   `json:"scheduled"`
}

//Material represents information about material such as Git or Mercurial repo
type Material struct {
	Fingerprint string `json:"fingerprint"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Id          int    `json:"id"`
}

//Pipeline represents information about a single pipeline
type Pipeline struct {
	Name      string     `json:"name"`
	Label     string     `json:"label"`
	Materials []Material `json:"materials"`
	Stages    []Stage    `json:"stages"`
}

//PipelineGroup holds informationa bout a single pipeline group
type PipelineGroup struct {
	Name      string     `json:"name"`
	Pipelines []Pipeline `json:"pipelines"`
}

type User struct {
	LoginName      string   `json:"login_name"`
	DisplayName    string   `json:"display_name"`
	Enabled        bool     `json:"enabled"`
	Email          string   `json:"email"`
	EmailMe        bool     `json:"email_me"`
	CheckinAliases []string `json:"checkin_aliases"`
}

type UEmbedded struct {
	Users []User `json:"users"`
}

type UserList struct {
	UEmbedded UEmbedded `json:"_embedded"`
}
