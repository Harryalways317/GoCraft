package gocraft

const version = "1.0.0"

type GoCraft struct {
	AppName string
	Debug   bool
	Version string
}

func (g *GoCraft) New(rootPath string) error {

	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware", "scripts"},
	}
	err := g.Init(pathConfig)
	if err != nil {
		return err
	}
	return nil
}

func (g *GoCraft) Delete(rootPath string) error {

	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware", "scripts"},
	}

	err1 := g.Destroy(pathConfig)
	if err1 != nil {
		return err1
	}

	return nil
}

func (g *GoCraft) Init(p initPaths) error {

	root := p.rootPath
	for _, path := range p.folderNames {
		//create dir if not exists
		err := g.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}

	}
	return nil
}

func (g *GoCraft) Destroy(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		//delete dir if  exists
		err := g.DeleteDirIfExists(root + "/" + path)
		if err != nil {
			return err
		}

	}
	return nil
}
