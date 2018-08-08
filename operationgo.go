/******************************
** Mission: Secrets Revealed **
*******************************
**
** IMPORTANT: You only get one chance to run your code on this level
** before your score is recorded on the leaderboard.
**
** You download and aggregate the USB data into a single file. As you
** sit down to dive in, your initial enthusiasm turns to concern when
** you see that as soon as some files are read, other files lock.
** It's absolutely critical for the case against Epoch to get as much
** data as you can. Epoch's file system allows you to input up to ten
** switching commands on the file locks. What's worse is you can't see
** which files are locked until you attempt to read the file. Everyone
** is counting on you to deliver as many readable files as possible.
**
**/

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// commands := setCommands()
	files := createFiles()
	for _, f := range files {
		fmt.Printf("original: f.megabytes = %d, f.locked = %t \n", f.megabytes, f.locked)
	}
	fmt.Printf("hero = %+v \n", findHero(files))
	// files.run(commands)
	// for _, f := range files {
	// 	fmt.Printf("operated: f.megabytes = %d, f.locked = %t \n", f.megabytes, f.locked)
	// }
	// megabytesCollected := files.countMegabytes()
	// fmt.Println("Megabytes Collected:", megabytesCollected)
}

var (
	fileSizes []int
	fileLocks []bool
)

func init() {
	for i := 0; i < 25; i++ {
		r := rand.Intn(1000)
		fileSizes = append(fileSizes, r)
		fileLocks = append(fileLocks, r%2 == 0)
	}
}

// File ...
type File struct {
	megabytes int
	locked    bool
}

// Files ...
type Files []File

func (files Files) Len() int {
	return len(files)
}

func (files Files) Swap(i, j int) {
	files[i], files[j] = files[j], files[i]
}

// --------------------------------------------------------------------------------

var findHero = func(files Files) *File {
	var hero *File
	for _, f := range files {
		if hero == nil {
			hero = &f
		}
		fmt.Printf("f.megabytes = %d, f.locked = %t, hero.megabytes = %d \n", f.megabytes, f.locked, hero.megabytes)
		if f.megabytes > (*hero).megabytes {
			hero = &f
		}
	}
	return hero
}

var hero *File

func (files Files) Less(i, j int) bool {
	if hero == nil {
		hero = findHero(files)
	}
	// fmt.Printf("%d %d, hero = %+v \n", i, j, hero)
	if files[i] == *hero {
		return false
	}
	if files[i].locked == files[j].locked {
		return files[i].megabytes < files[j].megabytes
	}
	return files[j].locked
}

func setCommands() []string {
	commands := make([]string, 10)
	commands[0] = "sort"
	commands[1] = "chain"
	commands[2] = "sort"
	commands[3] = "chain"
	commands[4] = "sort"
	commands[5] = "chain"
	commands[6] = "sort"
	commands[7] = "chain"
	// commands[4] = "sort"
	// commands[5] = "flip_intervals"
	return commands
}

// --------------------------------------------------------------------------------

func (files Files) run(commands []string) {
	for i := 0; i < 10; i++ {
		switch commands[i] {
		case "sort":
			sort.Sort(files)
		case "chain":
			files.chainLocks()
		case "flip_all":
			files.flipAllLocks()
		case "flip_intervals":
			files.flipIntervals()
		}
	}
}

// 倒序 i%4==0 的锁状态求反
func (files Files) flipIntervals() {
	for i := len(files) - 1; i > 0; i-- {
		if i%4 == 0 && files[i].megabytes > files[i-1].megabytes {
			files[i].locked = !files[i].locked
		}
	}
}

// 所有文件的锁状态求反
func (files Files) flipAllLocks() {
	for i := range files {
		files[i].locked = !files[i].locked
	}
}

// 倒序 如果最后一个文件为 unlocked 收益会比较大
func (files Files) chainLocks() {
	chain_type := files[len(files)-1].locked
	for i := len(files) - 2; i > 0; i-- {
		if files[i].locked != chain_type && files[i].megabytes > files[i-1].megabytes {
			files[i].locked = chain_type
		} else {
			return
		}
	}
}

func createFiles() Files {
	files := make(Files, 25)
	megabytes := 0

	for i := range files {
		files[i] = File{
			megabytes: fileSizes[i],
			locked:    fileLocks[i],
		}
		megabytes += fileSizes[i]
	}
	fmt.Println("Megabytes Available:", megabytes)
	return files
}

func (files Files) countMegabytes() int {
	megabytes := 0
	for i := range files {
		if !files[i].locked {
			megabytes += files[i].megabytes
		}
	}
	return megabytes
}
