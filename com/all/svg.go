package all

import "fmt"

func svg(name string) string {
	return fmt.Sprintf("'svg/%s.svg'", name)
}

var (
	SvgProject             = svg("project")
	SvgCommit              = svg("commit")
	SvgPullRequests        = svg("pullRequests")
	SvgStructure           = svg("structure")
	SvgMoreHorizontal      = svg("moreHorizontal")
	SvgMoreVertical        = svg("moreVertical")
	SvgVCS                 = svg("vcs")
	SvgProblems            = svg("problems")
	SvgTerminal            = svg("terminal")
	SvgServices            = svg("services")
	SvgPythonPackages      = svg("pythonPackages")
	SvgSourceRootFileLayer = svg("sourceRootFileLayer")
)
