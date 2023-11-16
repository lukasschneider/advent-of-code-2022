package filetree

import ()

type Filetree struct {
    Root *Dir
}

func NewTree() *Filetree {
    root := &Dir{
        name: "/",
        Depth: 0,
        Files: []int64{},
        Subdirs: []*Dir{},
        PrevDir: nil,
    }
    tree := &Filetree{
        Root: root,
    }
    
    return tree
}

func (t *Filetree) GetRoot() *Dir {
    return t.Root
}

type Dir struct {
    name string
    Depth int
    Files []int64 
    Subdirs []*Dir
    PrevDir *Dir 
}

func (t *Filetree) NewDir(prevDir *Dir, name string) *Dir {
    dir := &Dir{
        name: name,
        Depth: prevDir.Depth + 1,
        PrevDir: prevDir,
    }
    
    return dir
}

func (d *Dir) AddFile(file int64) { 
    d.Files = append(d.Files, file)
}

func (d *Dir) AddSubdir(subdir *Dir) {
    d.Subdirs = append(d.Subdirs, subdir)
}

func (d *Dir) GetSubdirs() []*Dir {
    return d.Subdirs
}

func (d *Dir) GetPrevDir(dir *Dir) *Dir {
    return dir.PrevDir
}

func (d *Dir) GetDepth(dir *Dir) int {
    return dir.Depth
}

func (d *Dir) GetFiles() []int64 {
    return d.Files
}

func (d *Dir) FindDir(name string) *Dir {
    for _, dir := range d.Subdirs {
        if dir.name == name {
            return dir
        }
    }
    return nil
}





