package superclog

var (
	CategoryBuild    = Category{"build", "Build"}
	CategoryChore    = Category{"chore", "Chores"}
	CategoryCI       = Category{"ci", "CI"}
	CategoryDocs     = Category{"docs", "Documentation"}
	CategoryFeat     = Category{"feat", "Features"}
	CategoryFix      = Category{"fix", "Fixes"}
	CategoryPerf     = Category{"perf", "Performance Improvements"}
	CategoryRefactor = Category{"refactor", "Refactoring"}
	CategoryRevert   = Category{"revert", "Reverting"}
	CategoryStyle    = Category{"style", "Style"}
	CategoryTest     = Category{"test", "Test"}
	CategoryUnknown  = Category{"", "Miscellaneous"}

	CategoryList = []Category{
		CategoryBuild,
		CategoryChore,
		CategoryCI,
		CategoryDocs,
		CategoryFeat,
		CategoryFix,
		CategoryPerf,
		CategoryRefactor,
		CategoryRevert,
		CategoryStyle,
		CategoryTest,
		CategoryUnknown,
	}
)

type Category struct {
	Value        string
	FriendlyName string
}

// CategoryOf gets the category of the change
func CategoryOf(v string) Category {
	for i := range CategoryList {
		if CategoryList[i].Value == v {
			return CategoryList[i]
		}
	}
	return CategoryUnknown
}

// Categories holds which category the commit is.
type Categories map[Category]Component

// CalculateCategories calculates categories for the
func CalculateCategories(commits []Commit) Categories {
	v := make(Categories, 10)
	for i := range commits {
		v.AddCommit(&commits[i])
	}
	return v
}

func (c Categories) AddCommit(commit *Commit) {
	typ := commit.Conventional.Type
	component := commit.Conventional.Component

	if _, ok := c[typ][component]; !ok {
		c[typ] = make(Component, 5)
	}

	c[typ][component] = append(c[typ][component], commit)

}
