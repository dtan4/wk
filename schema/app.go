package schema

type AppOwner struct {
	ID           string `json:"_id"`
	GravatarHash string `json:"gravatarHash"`
	Username     string `json:"username"`
}

type AppProjectFollower struct {
	ID           string `json:"id"`
	UserId       string `json:"userId"`
	GravatarHash string `json:"gravatarHash"`
	ProfileKey   string `json:"profileKey"`
	Username     string `json:"username"`
}

type AppDeployTarget struct {
	ID                              string   `json:"id"`
	Name                            string   `json:"name"`
	LastDeploySuccessful            bool     `json:"lastDeploySuccessful"`
	CurrentDeployHasWarnings        string   `json:"currentDeployHasWarnings"`
	CurrentDeployId                 string   `json:"currentDeployId"`
	CurrentDeployCreationDate       string   `json:"currentDeployCreateionDate"`
	CurrentDeployMessage            string   `json:"currentDeploymessage"`
	CurrentDeployUrl                string   `json:"currentDeployUrl"`
	CurrentDeployBuildId            string   `json:"currentDeployBuildId"`
	CurrentDeployBuildBranch        string   `json:"currentDeployBuildBranch"`
	CurrentDeployBuildCommit        string   `json:"currentDeployBuildCommit"`
	CurrentDeployBuildCommitMessage string   `json:"currentDeployBuildCommitMessage"`
	LastDeployId                    string   `json:"lastDeployId"`
	LastDeployCreationDate          string   `json:"lastDeployCreationDate"`
	LastDeployMessage               string   `json:"lastDeploymessage"`
	LastDeployBuildId               string   `json:"lastDeployBuildId"`
	LastDeployBuildBranch           string   `json:"lastDeployBuildBranch"`
	LastDeployBuildCommit           string   `json:"lastDeployBuildCommit"`
	LastDeployBuildCommitMessage    string   `json:"lastDeployBuildCommitMessage"`
	Type                            string   `json:"type"`
	AutoDeploy                      bool     `json:"autoDeploy"`
	AutoDeployBranches              []string `json:"autoDeployBranches"`
}

type App struct {
	ID                   string               `json:"id"`
	UserId               string               `json:"userId"`
	CurrentUserId        string               `json:"currentUserId"`
	ProjectOwner         string               `json:"projectOwner"`
	Owner                AppOwner             `json:"owner"`
	Name                 string               `json:"name"`
	URL                  string               `json:"url"`
	Author               string               `json:"author"`
	Date                 string               `json:"date"`
	BadgeKey             string               `json:"badgeKey"`
	TotalChecks          integer              `json:"totalChecks"`
	TotalChecksDone      integer              `json:"totalChecksDone"`
	Status               string               `json:"status"`
	BaseUrl              string               `json:"baseUrl"`
	AllowedActions       []string             `json:"allowedActions"`
	ProjectFollowers     []AppProjectFollower `json:"projectFollowers"`
	TotalFollwers        integer              `json:"totalFollowers"`
	IsFollower           bool                 `json:"isFollower"`
	RealtimeBuilds       bool                 `json:"realtimeBuilds"`
	DailyBuilds          bool                 `json:"dailyBuilds"`
	WeeklyBuilds         bool                 `json:"weeklyBuilds"`
	RealtimeDeploys      bool                 `json:"realtimeDeploys"`
	DailyDeploys         bool                 `json:"dailyDeploys"`
	WeeklyDeploys        bool                 `json:"weeklyDeploys"`
	Privacy              string               `json:"privacy"`
	Version              integer              `json:"version"`
	Nudges               []string             `json:"nudges"`
	DeployTargets        []AppDeployTarget    `json:"deployTargets"`
	FailedDeployCount    integer              `json:"failedDeployCount"`
	DeployWarnings       integer              `json:"deployWarnings"`
	HasWebHook           bool                 `json:"hasWebHook"`
	Registry             string               `json:"registry"`
	PrivateAllowed       bool                 `json:"privateAllowed"`
	AllowedDeployTargets []string             `json:"allowedDeployTargets"`
	FilterType           []string             `json:filterType`
}
