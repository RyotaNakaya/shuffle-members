package controller

// Project は シャッフルできる単位を表します
type Project struct {
}

// Create はプロジェクトの作成を行います
func (p *Project) Create() string {
	return "hoge"
}
