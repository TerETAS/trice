// Copyright 2020 Thomas.Hoehenleitner [at] seerose.net
// Use of this source code is governed by a license that can be found in the LICENSE file.

package id

// source tree management

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/rokath/trice/pkg/msg"
)

// todo: remove static values from parameter lists (Verbose)

const (
	// patSourceFile is a regex pattern matching any source file for patching
	patSourceFile = "(\\.c|\\.h|\\.cc|\\.cpp|\\.hpp)$"

	// patNbTRICE is a regex pattern matching any "TRICE*(Id(n), "", ... )". - see https://regex101.com/r/LNlAwY/9, The (?U) says non-greedy
	// patNbTRICE is a regex pattern matching any "TRICE*(Id(n), "", ... )". - see https://regex101.com/r/id0uSF/1
	//patNbTRICE = `(?U)(\bTRICE_S|TRICE0|TRICE8_[1-8]|TRICE16_[1-4]|TRICE32_[1-4]|TRICE64_[1-4]\b)\s*\(\s*\bId\b\s*\(\s*.*[0-9]\s*\)\s*,\s*".*"\s*.*\)`
	//patNbTRICE = `(?U)(\bTRICE[_S|0|8_1|8_2|8_3|8_4|8_5|8_6|8_7|8_8|16_1|16_2|16_3|16_4|32_1|32_2|32_3|32_4|64_1|64_2|64_3|64_4]\b)\s*\(\s*\bId\b\s*\(\s*.*[0-9]\s*\)\s*,\s*".*"\s*.*\)`
	patNbTRICE = `(?U)(\bTRICE_S|trice_s|TRICE0|trice0|TRICE8_[1-8]|trice8_[1-8]|TRICE16_[1-4]|trice16_[1-4]|TRICE32_[1-4]|trice32_[1-4]|TRICE64_[1-4]|trice64_[1-4]\b)\s*\(\s*\bId\b\s*\(\s*.*[0-9]\s*\)\s*,\s*".*"\s*.*\)`
	//patNbTRICE = `(?U)(\bTRICE_S|TRICE0|TRICE8_[1-8]|TRICE16_[1-4]|TRICE32_[1-4]|TRICE64_[1-4]\b)\s*\(\s*\bId\b\s*\(\s*.*[0-9]\s*\)\s*,\s*".*"\s*.*\)`
	// patNbID is a regex pattern matching any (first in string) "Id(n)" and usable in matches of matchNbTRICE
	patNbID = `\bId\s*\(\s*[0-9]*\s*\)`

	// patTypNameTRICE is a regex pattern matching "TRICE*" inside trice
	//patTypNameTRICE = `(\bTRICE_S\b|\bTRICE0\b|\bTRICE8_[1-8]\b|\bTRICE16_[1-4]\b|\bTRICE32_[1-4]\b|\bTRICE64_[1-4]\b)`
	patTypNameTRICE = `(\bTRICE_S\b|\bTRICE0\b|\bTRICE8_[1-8]\b|\bTRICE16_[1-4]\b|\bTRICE32_[1-4]\b|\bTRICE64_[1-4]\b|\btrice_S\b|\btrice0\b|\btrice8_[1-8]\b|\btrice16_[1-4]\b|\btrice32_[1-4]\b|\btrice64_[1-4]\b)`

	// patFmtString is a regex matching the first format string inside trice
	patFmtString = `"(.*)"`

	// patFullTriceWithoutID is a regex find a TRICE* line without Id, The (?U) says non-greedy
	//patFullTriceWithoutID = `(?U)(\bTRICE64|TRICE32|TRICE16|TRICE8|TRICE0|TRICE_S\b)\s*\(\s*".*"\s*.*\)`
	patFullTriceWithoutID = `(?U)(\bTRICE64|TRICE32|TRICE16|TRICE8|TRICE0|TRICE_S|trice64|trice32|trice16|trice8|trice0|trice_s\b)\s*\(\s*".*"\s*.*\)`

	// patTriceStartWithoutIDo is a regex
	//patTriceStartWithoutIDo = `(\bTRICE64|TRICE32|TRICE16|TRICE8|TRICE0|TRICE_S\b)\s*\(`
	patTriceStartWithoutIDo = `(\bTRICE64|TRICE32|TRICE16|TRICE8|TRICE0|TRICE_S|trice64|trice32|trice16|trice8|trice0|trice_s\b)\s*\(`

	// patTriceStartWithoutID is a regex
	//patTriceStartWithoutID = `(\bTRICE64|TRICE32|TRICE16|TRICE8|TRICE0|TRICE_S\b)\s*`
	patTriceStartWithoutID = `(\bTRICE64|TRICE32|TRICE16|TRICE8|TRICE0|TRICE_S|trice64|trice32|trice16|trice8|trice0|trice_s\b)\s*`

	// patNextFormatSpezifier is a regex find next format specifier in a string (exclude %%*)
	patNextFormatSpezifier = `(?:^|[^%])(%[0-9\.#]*(b|d|u|x|X|o|f))`
)

var (
	matchSourceFile           = regexp.MustCompile(patSourceFile)
	matchNbTRICE              = regexp.MustCompile(patNbTRICE)
	matchNbID                 = regexp.MustCompile(patNbID)
	matchTypNameTRICE         = regexp.MustCompile(patTypNameTRICE)
	matchFmtString            = regexp.MustCompile(patFmtString)
	matchFullTriceWithoutID   = regexp.MustCompile(patFullTriceWithoutID)
	matchTriceStartWithoutIDo = regexp.MustCompile(patTriceStartWithoutIDo)
	matchTriceStartWithoutID  = regexp.MustCompile(patTriceStartWithoutID)
	matchNextFormatSpezifier  = regexp.MustCompile(patNextFormatSpezifier)
)

func isSourceFile(fi os.FileInfo) bool {
	return matchSourceFile.MatchString(fi.Name())
}

func separatedIDsUpdate(root string, lu TriceIDLookUp, tflu TriceFmtLookUp) (modified bool) {
	// to do
	return false
}

// Additional actions needed: (Option -dry-run lets do a check in advance.)
// - Insert in all TRICE messages without ID an `Id(0),`
// - Check if all TRICE messages have correct parameter count and adapt the count without touching the Id(n),
// - Check if ID list has same ID more than one time and remove younger items with message.
//  - Read til.json in a map and write til.json after the map was finally manipulated back to til.json.
// - Check if Source tree has same ID with different TRICE strings.
//   - Keep the ID which is in til.json and set others to 0 with message.
//   - If none of the more than 1 time used ID is in til.json set all to 0 with message.
// - Check if in source code exist IDs not in til.json so far and extend til.json if there is no conflict.
//  - If the ID in soure code is already used in til.json differently set the ID in source code to 0 with message.
// NOT NEEDED: - Check if in til.json ID's not in source tree and mark them with `removed` timestamp.
// NOT NEEDED:   - If several source trees use same til.json, the `removed` timestamp is without sense.
// NOT NEEDED:   - If a `removed` timestamp is set, but the ID is in the source tree the `removed` timestamp is set to 0.

// Update is parsing source tree root and performing these actions:
// - replace.Type( Id(0), ...) with.Type( Id(n), ...)
// - find duplicate.Type( Id(n), ...) and replace one of them if trices are not identical
// - extend file fnIDList
func sharedIDsUpdate(root string, lu TriceIDLookUp, tflu TriceFmtLookUp) (modified bool) {
	if Verbose {
		fmt.Println("dir=", root)
		fmt.Println("List=", FnJSON)
	}
	msg.FatalInfoOnErr(filepath.Walk(root, visitUpdate(lu, tflu, &modified)), "failed to walk tree")
	return
}

func visitUpdate(lu TriceIDLookUp, tflu TriceFmtLookUp, pModified *bool) filepath.WalkFunc {
	// WalkFunc is the type of the function called for each file or directory
	// visited by Walk. The path argument contains the argument to Walk as a
	// prefix; that is, if Walk is called with "dir", which is a directory
	// containing the file "a", the walk function will be called with argument
	// "dir/a". The info argument is the os.FileInfo for the named path.
	//
	// If there was a problem walking to the file or directory named by path, the
	// incoming error will describe the problem and the function can decide how
	// to handle that error (and Walk will not descend into that directory). In the
	// case of an error, the info argument will be nil. If an error is returned,
	// processing stops. The sole exception is when the function returns the special
	// value SkipDir. If the function returns SkipDir when invoked on a directory,
	// Walk skips the directory's contents entirely. If the function returns SkipDir
	// when invoked on a non-directory file, Walk skips the remaining files in the
	// containing directory.
	return func(path string, fi os.FileInfo, err error) error {
		if err != nil || fi.IsDir() || !isSourceFile(fi) {
			return err // forward any error and do nothing
		}
		if Verbose {
			fmt.Println(path)
		}
		read, err := ioutil.ReadFile(path)
		if nil != err {
			return err
		}
		text := string(read)
		textN := updateParamCount(text)           // update parameter count: TRICE* to TRICE*_n
		textU := updateIDsShared(textN, lu, tflu) // update IDs: Id(0) -> Id(M)

		if 0 != strings.Compare(text, textU) {
			*pModified = true
		}

		// write out
		if *pModified && !DryRun {
			err = ioutil.WriteFile(path, []byte(textU), fi.Mode())
			if nil != err {
				return fmt.Errorf("failed to change %s: %v", path, err)
			}
		}
		return nil
	}
}

// tricePattern expects a string t containing a trice macro in the form `TRICE*(Id(n), "...", ...);`
// Returned id is the scanned n inside Id(n), only and only if n is a single decimal number.
// tf is the recognized trice.
// Only on success flag is true.
func tricePattern(t string) (id TriceID, tf TriceFmt, flag bool) {
	nbID := matchNbID.FindString(t)
	if "" == nbID {
		msg.InfoOnTrue(Verbose, fmt.Sprintln("No 'Id(n)' found inside "+t))
		return
	}
	n, e := strconv.Atoi(nbID)
	msg.FatalOnErr(e) // todo: error is here not possible
	id = TriceID(n)
	var tf TriceFmt
	tf.Type = matchTypNameTRICE.FindString(t)
	if "" == tf.Type {
		msg.Info(fmt.Sprintln("no 'TRICE*' found inside " + t))
		return
	}
}

// updateNextID is getting these parameters:
//    - p = pointer to ID List
//    - pListModified = pointer to the 'ID List modified flag', which is set true if s.th. changed in the List
//    - modified = the 'file modified flag', which is returned set true if s.th. changed in the file
//    - subs = the remaining file contents
//    - s = the full filecontents, which could be modified
//    - verbose flag
// updateNextID is returning these values (left to right):
//    - id flag is true if an ID was changed
//    - modified flag is true when any id was changed in the file
//    - subs gets shorter
//    - s is updated
func updateIDsShared(text string, lu TriceIDLookUp, tflu TriceFmtLookUp) string {
	subs := text[:] // create a copy of text and assign it to subs
	for {
		loc := matchNbTRICE.FindStringIndex(subs) // find the next TRICE location in file
		if nil == loc {
			return text // done
		}
		nbTRICE := subs[loc[0]:loc[1]] // full trice expression with
		nbID := matchNbID.FindString(nbTRICE)
		if "" == nbID {
			msg.Info(fmt.Sprintln("No 'Id(n)' found inside " + nbTRICE))
			subs = subs[loc[1]:]
			continue
		}
		var id TriceID

		_, err := fmt.Sscanf(nbID, "Id(%d", &id) // closing bracket in format string omitted intensionally
		if nil != err {                          // because spaces after id otherwise are not tolerated
			msg.Info(fmt.Sprintln("No 'Id(n)' found inside " + nbID))
			subs = subs[loc[1]:]
			continue
		}

		if 0 == id {
			zeroID := nbID
			zeroTRICE := nbTRICE
			// It is possible tf is already in tflu here
			id = lu.newID()
			newID := fmt.Sprintf("Id(%5d)", id)
			if Verbose {
				fmt.Println(zeroID, " -> ", newID)
			}
			nbTRICE := strings.Replace(nbTRICE, zeroID, newID, 1)
			text = strings.Replace(text, zeroTRICE, nbTRICE, 1)
		}

		// prepare subs for next loop
		subs = subs[loc[1]:] // The replacement makes s not shorter, so next search can start at loc[1]

		// At this place loc contains a trice with an ID and there are different cases:
		// tf
		// The map can be empty or not.
		// assuming tf is in the map with id0
		//		tflu[tf] = id will overwrite id0 in tflu
		//		lu[id] = tf will add id and tf will be found
		// It is possible this ID is completely new and need to be added to the map.
		// It is possible this ID is already in the map but with a different triceFmt. (ERROR)
		//

		var tf TriceFmt
		tf.Type = matchTypNameTRICE.FindString(nbTRICE)
		if "" == tf.Type {
			msg.Info(fmt.Sprintln("no 'TRICE*' found inside " + nbTRICE))
			continue
		}
		match := matchFmtString.FindAllStringSubmatch(nbTRICE, 1)
		tf.Strg = match[0][1]

		tflu[tf] = id
		lu[id] = tf

		nID, flag := p.ExtendIDList(id, typNameTRICE, fmtString)
		if flag {
			*pListModified = true
			if nID != id { // a new id was generated
				//oID := fmt.Sprintf("Id(%5d)", id)
				newID := fmt.Sprintf("Id(%5d)", nID)
				if Verbose {
					fmt.Println(nbID, " -> ", newID)
				}
				newTRICE := strings.Replace(nbTRICE, nbID, newID, 1)
				s = strings.Replace(s, nbTRICE, newTRICE, 1)
				modified = true
			}
		}
		return true, modified, subs, s // next done
	}
}

// updateParamCount stays in each file as long TRICE* statements without ID() are found.
// If a TRICE* is found it is getting an Id(0) inserted and it is also extended by _n
// according to the format specifier count inside the formatstring
//
// text is the full filecontents, which could be modified, therefore it is also returned
func updateParamCount(text string) string {
	subs := text[:] // create a copy of text and assign it to subs
	for {
		loc := matchFullTriceWithoutID.FindStringIndex(subs) // find the next TRICE location in file
		if nil == loc {
			return text // done
		}
		trice := subs[loc[0]:loc[1]]                                  // the whole TRICE*(*);
		triceO := matchTriceStartWithoutIDo.FindString(trice)         // TRICE*( part (the trice start)
		triceS := matchTriceStartWithoutID.FindString(trice)          // TRICE* part (the trice start)
		triceN := strings.Replace(trice, triceO, triceO+" Id(0),", 1) // insert Id(0)

		// count % format spezifier inside formatstring
		p := triceN
		var n int
		xs := "any"
		for "" != xs {
			lo := matchNextFormatSpezifier.FindStringIndex(p)
			xs = matchNextFormatSpezifier.FindString(p)
			if "" != xs { // found
				n++
				p = p[lo[1]:]
			} else {
				xs = ""
			}
		}
		if n > 0 { // patch
			newName := fmt.Sprintf(triceS+"_%d", n)              // TRICE*_n
			triceN = strings.Replace(triceN, triceS, newName, 1) // insert _n
		} else {
			// to do: handle special case 0==n
		}

		if Verbose {
			fmt.Println(trice)
			fmt.Println("->")
			fmt.Println(triceN)
		}
		text = strings.Replace(text, trice, triceN, 1) // modify s
		subs = subs[loc[1]:]                           // The replacement makes s not shorter, so next search can start at loc[1]
	}
}

// ZeroSourceTreeIds is overwriting with 0 all id's from source code tree srcRoot. It does not touch idlist.
func ZeroSourceTreeIds(srcRoot string, run bool) {
	err := filepath.Walk(srcRoot, visitZeroSourceTreeIds(run))
	if err != nil {
		panic(err)
	}
}

func visitZeroSourceTreeIds(run bool) filepath.WalkFunc {
	// WalkFunc is the type of the function called for each file or directory
	// visited by Walk. The path argument contains the argument to Walk as a
	// prefix; that is, if Walk is called with "dir", which is a directory
	// containing the file "a", the walk function will be called with argument
	// "dir/a". The info argument is the os.FileInfo for the named path.
	//
	// If there was a problem walking to the file or directory named by path, the
	// incoming error will describe the problem and the function can decide how
	// to handle that error (and Walk will not descend into that directory). In the
	// case of an error, the info argument will be nil. If an error is returned,
	// processing stops. The sole exception is when the function returns the special
	// value SkipDir. If the function returns SkipDir when invoked on a directory,
	// Walk skips the directory's contents entirely. If the function returns SkipDir
	// when invoked on a non-directory file, Walk skips the remaining files in the
	// containing directory.
	return func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() || !isSourceFile(fi) || err != nil {
			return err // forward any error and do nothing
		}

		fmt.Println(path)
		read, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		s := string(read)
		a, b := 0, len(s)
		subs := s[a:b]
		modified := false

		for {
			var found bool
			found, modified, subs, s = zeroNextID(modified, subs, s)
			if false == found {
				break
			}
		}

		if modified && true == run {
			err = ioutil.WriteFile(path, []byte(s), 0)
		}
		return err
	}
}

// first retval flag is if an ID was zeroed, others are updated input values. if an ID wsa zeroed
// - modified gets true
// - subs gets shorter
// - s is updated
func zeroNextID(modified bool, subs, s string) (bool, bool, string, string) {
	loc := matchNbTRICE.FindStringIndex(subs)
	if nil == loc {
		return false, modified, subs, s
	}
	nbTRICE := subs[loc[0]:loc[1]]
	nbID := matchNbID.FindString(nbTRICE)
	if "" == nbID {
		msg.Info(fmt.Sprintln("No 'Id(n)' found inside " + nbTRICE))
		return false, modified, subs, s
	}

	zeroID := "Id(0)"
	fmt.Println(nbID, " -> ", zeroID)

	zeroTRICE := strings.Replace(nbTRICE, nbID, zeroID, 1)
	s = strings.Replace(s, nbTRICE, zeroTRICE, 1)
	// 2^32 has 9 ciphers and shortest trice has 14 chars: TRICE0(Id(1),"");
	// The replacement of n with 0 makes s shorter, so the next search shoud start like 10 chars earlier.
	subs = subs[loc[1]-10:]
	return true, true, subs, s
}
