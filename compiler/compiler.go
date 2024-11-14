package compiler

type Compiler struct {
	programName string
}

func NewCompiler(programName string) *Compiler {
	return &Compiler{
		programName: programName,
	}
}

func (c *Compiler) CommandMain(args []string) {

}
