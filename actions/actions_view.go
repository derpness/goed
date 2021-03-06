package actions

import "github.com/tcolar/goed/core"

// Add a text selection to the view. from l1,c1 to l2,c2 (1 indexed)
func (a *ar) ViewAddSelection(viewId int64, l1, c1, l2, c2 int) {
	d(viewAddSelection{viewId: viewId, l1: l1, c1: c1, l2: l2, c2: c2})
}

// Enable/disable a view autoscrolling, while selecting (dragged selection + scrolling)
// By y,x increments. 0,0 means off
func (a *ar) ViewAutoScroll(viewId int64, y, x int) {
	d(viewAutoScroll{viewId: viewId, x: x, y: y, on: y != 0 || x != 0})
}

// send 'backspace' to the view
func (a *ar) ViewBackspace(viewId int64) {
	d(viewBackspace{viewId: viewId})
}

// return the current view location in the ui (1 indexed)
func (a *ar) ViewBounds(viewId int64) (ln, col, ln2, col2 int) {
	answer := make(chan int, 4)
	d(viewBounds{viewId: viewId, answer: answer})
	return <-answer, <-answer, <-answer, <-answer
}

// return the absolute path of the raw buffer backing the view
func (a *ar) ViewBufferLoc(viewId int64) string {
	answer := make(chan string, 1)
	d(viewBufferLoc{viewId: viewId, answer: answer})
	return <-answer
}

// remove all the view selections.
func (a *ar) ViewClearSelections(viewId int64) {
	d(viewClearSelections{viewId: viewId})
}

// stop the command currenty running in the view (for exec views.)
func (a *ar) ViewCmdStop(viewId int64) {
	d(viewCmdStop{viewId: viewId})
}

// return the nuber of columns (width) of the view.
func (a *ar) ViewCols(viewId int64) (cols int) {
	answer := make(chan int, 1)
	d(viewCols{viewId: viewId, answer: answer})
	return <-answer
}

// copy text from the view (current selection, if none, current line)
func (a *ar) ViewCopy(viewId int64) {
	d(viewCopy{viewId: viewId})
}

// cut text from the view (current selection, if none, current line)
func (a *ar) ViewCut(viewId int64) {
	d(viewCut{viewId: viewId})
}

// return the current cursor UI position in the view (1 indexed)
func (a *ar) ViewCursorCoords(viewId int64) (y, x int) {
	answer := make(chan int, 2)
	d(viewCursorCoords{viewId: viewId, answer: answer})
	return <-answer, <-answer
}

// return the file offset at the current position in the backing buffer
func (a *ar) ViewCursorFileOffset(viewId int64) int64 {
	answer := make(chan int64, 1)
	d(viewCursorFileOffset{viewId: viewId, answer: answer})
	return <-answer
}

// return the current cursor text position in the view (1 indexed)
func (a *ar) ViewCursorPos(viewId int64) (y, x int) {
	answer := make(chan int, 2)
	d(viewCursorPos{viewId: viewId, answer: answer})
	return <-answer, <-answer
}

// send a movement event to the view. (ie: down, up, left, right, etc...)
func (a *ar) ViewCursorMvmt(viewId int64, mvmt core.CursorMvmt) {
	d(viewCursorMvmt{viewId: viewId, mvmt: mvmt})
}

// delete text from the view (from row1,col1 to row2,col2). 1 indexed
func (a *ar) ViewDelete(viewId int64, row1, col1, row2, col2 int, undoable bool) {
	d(viewDeleteAction{viewId: viewId, row1: row1, col1: col1, row2: row2, col2: col2, undoable: undoable})
}

// delete text from the view (current selection, if none, current line)
func (a *ar) ViewDeleteCur(viewId int64) {
	d(viewDeleteCur{viewId: viewId})
}

// is the view dirty or not ?
func (a *ar) ViewDirty(viewId int64) bool {
	answer := make(chan bool, 1)
	d(viewDirty{answer: answer, viewId: viewId})
	return <-answer
}

// insert text into the view at the row,col location. 1 indexed
func (a *ar) ViewInsert(viewId int64, row, col int, text string, undoable bool) {
	d(viewInsertAction{viewId: viewId, row: row, col: col, text: text, undoable: undoable})
}

// insert text into the view at the current cursor location
func (a *ar) ViewInsertCur(viewId int64, text string) {
	d(viewInsertCur{viewId: viewId, text: text})
}

// insert a newLine at the current cursor location
func (a *ar) ViewInsertNewLine(viewId int64) {
	d(viewInsertNewLine{viewId: viewId})
}

// move the cursor by ln, col runes (relative), scroll as needed
// roll means "roll" to prev/next line on column overflow
func (a *ar) ViewMoveCursor(viewId int64, y, x int, roll bool) {
	d(viewMoveCursor{viewId: viewId, x: x, y: y, roll: roll})
}

// paste text into the view at the curent location
// if in a selection, paste over it.
func (a *ar) ViewPaste(viewId int64) {
	d(viewPaste{viewId: viewId})
}

// try to "open" the current selection into a view (ie: expect a file path)
func (a *ar) ViewOpenSelection(viewId int64, newView bool) {
	d(viewOpenSelection{viewId: viewId, newView: newView})
}

// redo
func (a *ar) ViewRedo(viewId int64) {
	d(viewRedo{viewId: viewId})
}

// reload the view from it's source file, discard all unsaved buffer changes
func (a *ar) ViewReload(viewId int64) {
	d(viewReload{viewId: viewId})
}

// render/repaint the view
func (a *ar) ViewRender(viewId int64) {
	d(viewRender{viewId: viewId})
}

// return the number of rows (lines) in the view
func (a *ar) ViewRows(viewId int64) (rows int) {
	answer := make(chan int, 1)
	d(viewRows{viewId: viewId, answer: answer})
	return <-answer
}

// save the view content to the backing file
func (a *ar) ViewSave(viewId int64) {
	d(viewSave{viewId: viewId})
}

// select all
func (a *ar) ViewSelectAll(viewId int64) {
	d(viewSelectAll{viewId: viewId})
}

// Return a list of view selctions (one per line), 1 indexed, ie:
// 2 1 2 6
// 3 2 4 7
func (a *ar) ViewSelections(viewId int64) []core.Selection {
	answer := make(chan []core.Selection, 1)
	d(viewSelections{answer: answer, viewId: viewId})
	return <-answer
}

// select "word" at given path
func (a *ar) ViewSelectWord(viewId int64, ln, col int) {
	d(viewSelectWord{viewId: viewId, ln: ln, col: col})
}

// move the cursor to the given text position(1 indexed), scroll as needed
func (a *ar) ViewSetCursorPos(viewId int64, y, x int) {
	d(viewSetCursorPos{viewId: viewId, y: y, x: x})
}

// mark the view "dirty" or not (ie: modified, unsaved)
func (a *ar) ViewSetDirty(viewId int64, on bool) {
	d(viewSetDirty{viewId: viewId, on: on})
}

// set scrolling offsets as percentage (of text)
func (a *ar) ViewSetScrollPct(viewId int64, ypct int) {
	d(viewSetScrollPct{viewId: viewId, ypct: ypct})
}

// set scrolling offsets (1 indexed)
func (a *ar) ViewSetScrollPos(viewId int64, ln, col int) {
	d(viewSetScrollPos{viewId: viewId, ln: ln, col: col})
}

// set the view title (typically file path)
func (a *ar) ViewSetTitle(viewId int64, title string) {
	d(viewSetTitle{viewId: viewId, title: title})
}

// set the view type, see core.vars
func (a *ar) ViewSetType(viewId int64, viewType int) {
	d(viewSetType{viewId: viewId, viewType: viewType})
}

// set the number of vt100 columns, this is useful so that tty programs that
// can use the full view wisth properly
func (a *ar) ViewSetVtCols(viewId int64, cols int) {
	d(viewSetVtCols{viewId: viewId, cols: cols})
}

// set the current working dir of the view, especially usefull for terminal views.
// this is used when "opening" relative locations, among other things.
func (a *ar) ViewSetWorkDir(viewId int64, workDir string) {
	d(viewSetWorkDir{viewId: viewId, workDir: workDir})
}

// return the absolute path of the file backing the view (if any)
func (a *ar) ViewSrcLoc(viewId int64) string {
	answer := make(chan string, 1)
	d(viewSrcLoc{viewId: viewId, answer: answer})
	return <-answer
}

// return the current scrolling position (1,1 is top, left)
func (a *ar) ViewScrollPos(viewId int64) (ln, col int) {
	answer := make(chan int, 2)
	d(viewScrollPos{viewId: viewId, answer: answer})
	return <-answer, <-answer
}

// this forces a sync of the in memory slice representing the part of the content
// that is currently visible in the view (performance optimization)
func (a *ar) ViewSyncSlice(viewId int64) {
	d(viewSyncSlice{viewId: viewId})
}

// returns a slice of the buffer text from ln1,col1 to ln2,col2 (inclusive). 1 indexed
// note: col2==-1 means to end of line; ln2==-1 means to last line
func (a *ar) ViewText(viewId int64, ln1, col1, ln2, col2 int) []string {
	answer := make(chan []string, 1)
	d(viewText{viewId: viewId, ln1: ln1, col1: col1, ln2: ln2, col2: col2, answer: answer})
	return <-answer
}

// return the text position(1 indexed) for the given y,x cursor coordinates (0 indexed)
// if the given coordinates are not on text, return the closest text position.
// typically would be passed coordinates gotten from EdViewAt.
func (a *ar) ViewTextPos(viewId int64, y, x int) (ln, col int) {
	answer := make(chan int, 2)
	d(viewTextPos{viewId: viewId, answer: answer, y: y, x: x})
	return <-answer, <-answer
}

// return the vew title
func (a *ar) ViewTitle(viewId int64) string {
	answer := make(chan string, 1)
	d(viewTitle{viewId: viewId, answer: answer})
	return <-answer
}

// return the vew type (core.ViewType)
func (a *ar) ViewType(viewId int64) int {
	answer := make(chan int, 1)
	d(viewType{viewId: viewId, answer: answer})
	return <-answer
}

// undo
func (a *ar) ViewUndo(viewId int64) {
	d(viewUndo{viewId: viewId})
}

// working directory
func (a *ar) ViewWorkDir(viewId int64) string {
	answer := make(chan string, 1)
	d(viewWorkDir{viewId: viewId, answer: answer})
	return <-answer
}

// send raw bytes to a terminal view
func (a *ar) TermSendBytes(viewId int64, data []byte) {
	d(termSendBytes{viewId: viewId, data: data})
}

// ########  Impl ......

type viewAddSelection struct {
	viewId         int64
	l1, c1, l2, c2 int
}

func (a viewAddSelection) Run() {
	v := core.Ed.ViewById(a.viewId)
	if a.l2 != -1 {
		a.l2--
	}
	if a.c2 != -1 {
		a.c2--
	}
	if v != nil {
		s := core.NewSelection(a.l1-1, a.c1-1, a.l2, a.c2)
		selections := v.Selections()
		*selections = append(*selections, *s)
	}
}

type viewAutoScroll struct {
	viewId int64
	y, x   int
	on     bool
}

func (a viewAutoScroll) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SetAutoScroll(a.y, a.x, a.on)
	}
}

type viewBackspace struct {
	viewId int64
}

func (a viewBackspace) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.Backspace()
	}
}

type viewBounds struct {
	answer chan int
	viewId int64
}

func (a viewBounds) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		a.answer <- 0
		a.answer <- 0
		a.answer <- 0
		a.answer <- 0
		return
	}
	l1, c1, l2, c2 := v.Bounds()
	a.answer <- l1 + 1
	a.answer <- c1 + 1
	a.answer <- l2 + 1
	a.answer <- c2 + 1
}

type viewBufferLoc struct {
	viewId int64
	answer chan string
}

func (a viewBufferLoc) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil || v.Id() == 0 {
		a.answer <- ""
		return
	}
	a.answer <- v.Backend().BufferLoc()
}

type viewClearSelections struct {
	viewId int64
}

func (a viewClearSelections) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.ClearSelections()
	}
}

type viewCmdStop struct {
	viewId int64
}

func (a viewCmdStop) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		return
	}
	b := v.Backend()
	if b != nil {
		b.Close()
	}
	return
}

type viewCols struct {
	answer chan int
	viewId int64
}

func (a viewCols) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		a.answer <- 0
		return
	}
	a.answer <- v.LastViewCol()
}

type viewCopy struct {
	viewId int64
}

func (a viewCopy) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.Copy()
	}
}

type viewCursorCoords struct {
	answer chan int
	viewId int64
}

func (a viewCursorCoords) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		a.answer <- 0
		a.answer <- 0
		return
	}
	a.answer <- v.CurLine() + 1
	a.answer <- v.CurCol() + 1
}

type viewCursorFileOffset struct {
	answer chan int64
	viewId int64
}

func (a viewCursorFileOffset) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil || v.Backend() == nil {
		a.answer <- int64(-1)
		return
	}
	ln := v.CurLine()
	col := v.LineRunesTo(v.Slice(), v.CurLine(), v.CurCol())
	a.answer <- v.Backend().OffsetAt(ln, col)
}

type viewCursorPos struct {
	answer chan int
	viewId int64
}

func (a viewCursorPos) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		a.answer <- 0
		a.answer <- 0
		return
	}
	a.answer <- v.CurLine() + 1
	a.answer <- v.LineRunesTo(v.Slice(), v.CurLine(), v.CurCol()) + 1
}

type viewCursorMvmt struct {
	viewId int64
	mvmt   core.CursorMvmt
}

func (a viewCursorMvmt) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.CursorMvmt(a.mvmt)
	}
}

type viewCut struct {
	viewId int64
}

func (a viewCut) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.Cut()
	}
}

type viewDeleteAction struct {
	viewId                 int64
	row1, col1, row2, col2 int
	undoable               bool
}

func (a viewDeleteAction) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		if a.row2 != -1 {
			a.row2--
		}
		if a.col2 != -1 {
			a.col2--
		}
		v.Delete(a.row1-1, a.col1-1, a.row2, a.col2, a.undoable)
	}
}

type viewDeleteCur struct {
	viewId int64
}

func (a viewDeleteCur) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.DeleteCur()
	}
}

type viewDirty struct {
	viewId int64
	answer chan bool
}

func (a viewDirty) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		a.answer <- v.Dirty()
	}
	a.answer <- false
}

type viewInsertAction struct {
	viewId   int64
	row, col int
	text     string
	undoable bool
}

func (a viewInsertAction) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.Insert(a.row-1, a.col-1, a.text, a.undoable)
	}
}

type viewInsertCur struct {
	viewId int64
	text   string
}

func (a viewInsertCur) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		return
	}
	v.InsertCur(a.text)
}

type viewInsertNewLine struct {
	viewId int64
}

func (a viewInsertNewLine) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.InsertNewLineCur()
	}
}

type viewMoveCursor struct {
	viewId int64
	y, x   int
	roll   bool
}

func (a viewMoveCursor) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		return
	}
	if a.roll {
		v.MoveCursorRoll(a.y, a.x)
	} else {
		v.MoveCursor(a.y, a.x)
	}
}

type viewOpenSelection struct {
	viewId  int64
	newView bool
}

func (a viewOpenSelection) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.OpenSelection(a.newView)
	}
}

type viewPaste struct {
	viewId int64
}

func (a viewPaste) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.Paste()
	}
}

type viewRedo struct {
	viewId int64
}

func (a viewRedo) Run() {
	if viewExists(a.viewId) {
		Redo(a.viewId)
	}
}

type viewReload struct{ viewId int64 }

func (a viewReload) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.Reload()
	}
}

type viewRender struct {
	viewId int64
}

func (a viewRender) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.Render()
	}
}

type viewRows struct {
	answer chan int
	viewId int64
}

func (a viewRows) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		a.answer <- 0
		return
	}
	a.answer <- v.LastViewLine()
}

type viewSave struct {
	viewId int64
}

func (a viewSave) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.Save()
	}
}

type viewScrollPos struct {
	answer chan int
	viewId int64
}

func (a viewScrollPos) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		a.answer <- 0
		a.answer <- 0
		return
	}
	y, x := v.ScrollPos()
	a.answer <- y + 1
	a.answer <- x + 1
}

type viewSelectAll struct {
	viewId int64
}

func (a viewSelectAll) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SelectAll()
	}
}

type viewSelections struct {
	answer chan []core.Selection
	viewId int64
}

func (a viewSelections) Run() {
	v := core.Ed.ViewById(a.viewId)
	result := []core.Selection{}
	if v == nil {
		a.answer <- result
		return
	}
	for _, s := range *v.Selections() {
		ct := 1
		lt := 1
		if s.ColTo == -1 {
			ct = 0
		}
		if s.LineTo == -1 {
			lt = 0
		}
		result = append(result, *core.NewSelection(s.LineFrom+1, s.ColFrom+1,
			s.LineTo+lt, s.ColTo+ct))
	}
	a.answer <- result
}

type viewSelectWord struct {
	viewId  int64
	ln, col int
}

func (a viewSelectWord) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SelectWord(a.ln-1, a.col-1)
	}
}

type viewSetCursorPos struct {
	viewId int64
	y, x   int
}

func (a viewSetCursorPos) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SetCursorPos(a.y-1, a.x-1)
	}
}

type viewSetDirty struct {
	viewId int64
	on     bool
}

func (a viewSetDirty) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SetDirty(a.on)
	}
}

type viewSetScrollPct struct {
	viewId int64
	ypct   int
}

func (a viewSetScrollPct) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SetScrollPct(a.ypct)
	}
}

type viewSetScrollPos struct {
	viewId  int64
	ln, col int
}

func (a viewSetScrollPos) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SetScrollPos(a.ln-1, a.col-1)
	}
}

type viewSetTitle struct {
	viewId int64
	title  string
}

func (a viewSetTitle) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SetTitle(a.title)
	}
}

type viewSetType struct {
	viewId   int64
	viewType int
}

func (a viewSetType) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SetViewType(core.ViewType(a.viewType))
	}
}

type viewSetWorkDir struct {
	viewId  int64
	workDir string
}

func (a viewSetWorkDir) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SetWorkDir(a.workDir)
	}
}

type viewSetVtCols struct {
	viewId int64
	cols   int
}

func (a viewSetVtCols) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SetVtCols(a.cols)
	}
}

type viewSrcLoc struct {
	viewId int64
	answer chan string
}

func (a viewSrcLoc) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil || v.Id() == 0 {
		a.answer <- ""
		return
	}
	a.answer <- v.Backend().SrcLoc()
}

type viewSyncSlice struct {
	viewId int64
}

func (a viewSyncSlice) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v != nil {
		v.SyncSlice()
	}
}

type viewText struct {
	viewId               int64
	ln1, col1, ln2, col2 int
	answer               chan []string
}

func (a viewText) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil || v.Backend() == nil || a.col2 == 0 || a.col1 == 0 || a.ln1 == 0 || a.ln2 == 0 {
		a.answer <- []string{}
		return
	}
	if a.col2 > 0 {
		a.col2--
	}
	if a.ln2 > 0 {
		a.ln2--
	}
	strs := []string{}
	text := v.Text(a.ln1-1, a.col1-1, a.ln2, a.col2)
	for _, s := range text {
		strs = append(strs, string(s))
	}
	a.answer <- strs
}

type viewTextPos struct {
	viewId int64
	y, x   int
	answer chan int
}

func (a viewTextPos) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil {
		a.answer <- 1
		a.answer <- 1
		return
	}
	sy, sx := v.ScrollPos()
	ln := a.y - 2 + sy
	if ln < 1 {
		ln = 1
	} else if ln > v.LineCount() {
		ln = v.LineCount()
	}
	to := a.x - 2 + sx - 1
	if to < 0 {
		to = 0
	}
	col := v.LineRunesTo(v.Slice(), ln-1, to) + 1
	a.answer <- ln
	a.answer <- col
}

type viewTitle struct {
	viewId int64
	answer chan string
}

func (a viewTitle) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil || v.Id() == 0 {
		a.answer <- ""
		return
	}
	a.answer <- v.Title()
}

type viewType struct {
	viewId int64
	answer chan int
}

func (a viewType) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil || v.Id() == 0 {
		a.answer <- int(core.ViewTypeStandard)
		return
	}
	a.answer <- int(v.Type())
}

type viewUndo struct {
	viewId int64
}

func (a viewUndo) Run() {
	if viewExists(a.viewId) {
		Undo(a.viewId)
	}
}

type viewWorkDir struct {
	viewId int64
	answer chan string
}

func (a viewWorkDir) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil || v.Id() == 0 {
		a.answer <- ""
		return
	}
	a.answer <- v.WorkDir()
}

type termSendBytes struct {
	viewId int64
	data   []byte
}

func (a termSendBytes) Run() {
	v := core.Ed.ViewById(a.viewId)
	if v == nil || v.Type() != core.ViewTypeShell {
		return
	}
	v.Backend().SendBytes(a.data)
}

func NewViewInsertAction(viewId int64, row, col int, text string, undoable bool) core.Action {
	return viewInsertAction{viewId: viewId, row: row + 1, col: col + 1,
		text: text, undoable: undoable}
}

func NewViewDeleteAction(viewId int64, row1, col1, row2, col2 int, undoable bool) core.Action {
	return viewDeleteAction{viewId: viewId, row1: row1 + 1, col1: col1 + 1,
		row2: row2 + 1, col2: col2 + 1, undoable: undoable}
}

func NewSetCursorAction(viewId int64, ln, col int) core.Action {
	return viewSetCursorPos{viewId: viewId, y: ln + 1, x: col + 1}
}

func NewSetSelectionsActions(viewId int64, selections *[]core.Selection) []core.Action {
	a := []core.Action{
		viewClearSelections{viewId: viewId},
	}
	for _, s := range *selections {
		a = append(a, viewAddSelection{viewId: viewId, l1: s.LineFrom + 1, c1: s.ColFrom + 1, l2: s.LineTo + 1, c2: s.ColTo + 1})
	}
	return a
}
