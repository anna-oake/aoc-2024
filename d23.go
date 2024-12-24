package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type d23net struct {
	pcs   map[string]*d23pc
	links map[string]bool
}

type d23pc struct {
	n     *d23net
	id    string
	links []*d23pc
}

func (n *d23net) getPC(id string) *d23pc {
	pc := n.pcs[id]
	if pc == nil {
		pc = &d23pc{
			n:  n,
			id: id,
		}
		n.pcs[id] = pc
	}
	return pc
}

func (p1 *d23pc) addLink(p2 *d23pc) {
	p1.links = append(p1.links, p2)
	p2.links = append(p2.links, p1)
	k1 := p1.id + "-" + p2.id
	k2 := p2.id + "-" + p1.id
	p1.n.links[k1] = true
	p1.n.links[k2] = true
}

func (p1 *d23pc) hasLink(p2 *d23pc) bool {
	key := p1.id + "-" + p2.id
	return p1.n.links[key]
}

func (pc *d23pc) linksAsStrings() []string {
	var slinks []string
	for _, pc2 := range pc.links {
		slinks = append(slinks, pc2.id)
	}
	return slinks
}

func (*methods) D23P1(input string) string {
	net := &d23net{
		pcs:   make(map[string]*d23pc),
		links: make(map[string]bool),
	}
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		parts := strings.Split(l, "-")
		pc1 := net.getPC(parts[0])
		pc2 := net.getPC(parts[1])
		pc1.addLink(pc2)
	}

	groups := make(map[string]bool)

	for id, pc1 := range net.pcs {
		if !strings.HasPrefix(id, "t") {
			continue
		}
		for _, pc2 := range pc1.links {
			for _, pc3 := range pc1.links {
				if pc2.id == pc3.id || pc1.id == pc2.id || pc1.id == pc3.id {
					continue
				}
				if pc2.hasLink(pc3) {
					g := []string{pc1.id, pc2.id, pc3.id}
					sort.StringSlice(g).Sort()
					groups[strings.Join(g, ",")] = true
				}
			}
		}
	}

	return strconv.Itoa(len(groups))
}

func (*methods) D23P2(input string) string {
	net := &d23net{
		pcs:   make(map[string]*d23pc),
		links: make(map[string]bool),
	}
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		parts := strings.Split(l, "-")
		pc1 := net.getPC(parts[0])
		pc2 := net.getPC(parts[1])
		pc1.addLink(pc2)
	}

	var largestGroup string

	var i int
	for _, pc := range net.pcs {
		i++
		fmt.Println(i, len(net.pcs))
		g := pc.findGroup()
		if len(g) > len(largestGroup) {
			largestGroup = g
		}
	}

	return largestGroup
}

func (pc1 *d23pc) findGroup() string {
	inters := [][]string{}

	for _, pc2 := range pc1.links {
		pc2sl := pc2.linksAsStrings()
		inters = append(inters, pc2sl)
	}

	for {
		c := make(map[string]bool)
		tmp := [][]string{}
		for i, i1 := range inters {
			for j, i2 := range inters {
				if j <= i {
					continue
				}
				inter := intersection(i1, i2)
				sort.StringSlice(inter).Sort()
				key := strings.Join(inter, ",")
				if len(tmp) == 0 || len(inter) > len(tmp[0]) {
					tmp = [][]string{inter}
					c[key] = true
				} else if len(inter) == len(tmp[0]) {
					if !c[key] {
						tmp = append(tmp, inter)
						c[key] = true
					}
				}
			}
		}
		if len(tmp[0]) == 2 {
			ans := make(map[string]bool)
			for _, t := range tmp {
				if pc1.n.getPC(t[0]).hasLink(pc1.n.getPC(t[1])) {
					ans[t[0]] = true
					ans[t[1]] = true
				}
			}
			var members []string
			for m := range ans {
				members = append(members, m)
			}
			sort.StringSlice(members).Sort()
			return strings.Join(members, ",")
		}
		inters = tmp
		if len(inters) < 2 {
			return ""
		}
	}
}

func intersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	return
}
