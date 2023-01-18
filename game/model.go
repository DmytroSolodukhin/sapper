package game

type ShowCell struct {
	x int
	y int
	open bool
}

type Cell struct {
	ShowCell
	hflag bool
	marker int
}
