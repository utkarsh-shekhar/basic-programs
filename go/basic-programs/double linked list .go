
...
The Go Programming Language
Documents Packages The Project Help Blog Play
Source file src/container/list/list.go

     1	// Copyright 2009 The Go Authors. All rights reserved.
     2	// Use of this source code is governed by a BSD-style
     3	// license that can be found in the LICENSE file.
     4	
     5	// Package list implements a doubly linked list.
     6	//
     7	// To iterate over a list (where l is a *List):
     8	//	for e := l.Front(); e != nil; e = e.Next() {
     9	//		// do something with e.Value
    10	//	}
    11	//
    12	package list
    13	
    14	// Element is an element of a linked list.
    15	type Element struct {
    16		// Next and previous pointers in the doubly-linked list of elements.
    17		// To simplify the implementation, internally a list l is implemented
    18		// as a ring, such that &l.root is both the next element of the last
    19		// list element (l.Back()) and the previous element of the first list
    20		// element (l.Front()).
    21		next, prev *Element
    22	
    23		// The list to which this element belongs.
    24		list *List
    25	
    26		// The value stored with this element.
    27		Value interface{}
    28	}
    29	
    30	// Next returns the next list element or nil.
    31	func (e *Element) Next() *Element {
    32		if p := e.next; e.list != nil && p != &e.list.root {
    33			return p
    34		}
    35		return nil
    36	}
    37	
    38	// Prev returns the previous list element or nil.
    39	func (e *Element) Prev() *Element {
    40		if p := e.prev; e.list != nil && p != &e.list.root {
    41			return p
    42		}
    43		return nil
    44	}
    45	
    46	// List represents a doubly linked list.
    47	// The zero value for List is an empty list ready to use.
    48	type List struct {
    49		root Element // sentinel list element, only &root, root.prev, and root.next are used
    50		len  int     // current list length excluding (this) sentinel element
    51	}
    52	
    53	// Init initializes or clears list l.
    54	func (l *List) Init() *List {
    55		l.root.next = &l.root
    56		l.root.prev = &l.root
    57		l.len = 0
    58		return l
    59	}
    60	
    61	// New returns an initialized list.
    62	func New() *List { return new(List).Init() }
    63	
    64	// Len returns the number of elements of list l.
    65	// The complexity is O(1).
    66	func (l *List) Len() int { return l.len }
    67	
    68	// Front returns the first element of list l or nil.
    69	func (l *List) Front() *Element {
    70		if l.len == 0 {
    71			return nil
    72		}
    73		return l.root.next
    74	}
    75	
    76	// Back returns the last element of list l or nil.
    77	func (l *List) Back() *Element {
    78		if l.len == 0 {
    79			return nil
    80		}
    81		return l.root.prev
    82	}
    83	
    84	// lazyInit lazily initializes a zero List value.
    85	func (l *List) lazyInit() {
    86		if l.root.next == nil {
    87			l.Init()
    88		}
    89	}
    90	
    91	// insert inserts e after at, increments l.len, and returns e.
    92	func (l *List) insert(e, at *Element) *Element {
    93		n := at.next
    94		at.next = e
    95		e.prev = at
    96		e.next = n
    97		n.prev = e
    98		e.list = l
    99		l.len++
   100		return e
   101	}
   102	
   103	// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
   104	func (l *List) insertValue(v interface{}, at *Element) *Element {
   105		return l.insert(&Element{Value: v}, at)
   106	}
   107	
   108	// remove removes e from its list, decrements l.len, and returns e.
   109	func (l *List) remove(e *Element) *Element {
   110		e.prev.next = e.next
   111		e.next.prev = e.prev
   112		e.next = nil // avoid memory leaks
   113		e.prev = nil // avoid memory leaks
   114		e.list = nil
   115		l.len--
   116		return e
   117	}
   118	
   119	// Remove removes e from l if e is an element of list l.
   120	// It returns the element value e.Value.
   121	func (l *List) Remove(e *Element) interface{} {
   122		if e.list == l {
   123			// if e.list == l, l must have been initialized when e was inserted
   124			// in l or l == nil (e is a zero Element) and l.remove will crash
   125			l.remove(e)
   126		}
   127		return e.Value
   128	}
   129	
   130	// PushFront inserts a new element e with value v at the front of list l and returns e.
   131	func (l *List) PushFront(v interface{}) *Element {
   132		l.lazyInit()
   133		return l.insertValue(v, &l.root)
   134	}
   135	
   136	// PushBack inserts a new element e with value v at the back of list l and returns e.
   137	func (l *List) PushBack(v interface{}) *Element {
   138		l.lazyInit()
   139		return l.insertValue(v, l.root.prev)
   140	}
   141	
   142	// InsertBefore inserts a new element e with value v immediately before mark and returns e.
   143	// If mark is not an element of l, the list is not modified.
   144	func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
   145		if mark.list != l {
   146			return nil
   147		}
   148		// see comment in List.Remove about initialization of l
   149		return l.insertValue(v, mark.prev)
   150	}
   151	
   152	// InsertAfter inserts a new element e with value v immediately after mark and returns e.
   153	// If mark is not an element of l, the list is not modified.
   154	func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
   155		if mark.list != l {
   156			return nil
   157		}
   158		// see comment in List.Remove about initialization of l
   159		return l.insertValue(v, mark)
   160	}
   161	
   162	// MoveToFront moves element e to the front of list l.
   163	// If e is not an element of l, the list is not modified.
   164	func (l *List) MoveToFront(e *Element) {
   165		if e.list != l || l.root.next == e {
   166			return
   167		}
   168		// see comment in List.Remove about initialization of l
   169		l.insert(l.remove(e), &l.root)
   170	}
   171	
   172	// MoveToBack moves element e to the back of list l.
   173	// If e is not an element of l, the list is not modified.
   174	func (l *List) MoveToBack(e *Element) {
   175		if e.list != l || l.root.prev == e {
   176			return
   177		}
   178		// see comment in List.Remove about initialization of l
   179		l.insert(l.remove(e), l.root.prev)
   180	}
   181	
   182	// MoveBefore moves element e to its new position before mark.
   183	// If e or mark is not an element of l, or e == mark, the list is not modified.
   184	func (l *List) MoveBefore(e, mark *Element) {
   185		if e.list != l || e == mark || mark.list != l {
   186			return
   187		}
   188		l.insert(l.remove(e), mark.prev)
   189	}
   190	
   191	// MoveAfter moves element e to its new position after mark.
   192	// If e or mark is not an element of l, or e == mark, the list is not modified.
   193	func (l *List) MoveAfter(e, mark *Element) {
   194		if e.list != l || e == mark || mark.list != l {
   195			return
   196		}
   197		l.insert(l.remove(e), mark)
   198	}
   199	
   200	// PushBackList inserts a copy of an other list at the back of list l.
   201	// The lists l and other may be the same.
   202	func (l *List) PushBackList(other *List) {
   203		l.lazyInit()
   204		for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
   205			l.insertValue(e.Value, l.root.prev)
   206		}
   207	}
   208	
   209	// PushFrontList inserts a copy of an other list at the front of list l.
   210	// The lists l and other may be the same.
   211	func (l *List) PushFrontList(other *List) {
   212		l.lazyInit()
   213		for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
   214			l.insertValue(e.Value, &l.root)
   215		}
   216	}
   217	

View as plain text
Build version go1.5.2.
Except as noted, the content of this page is licensed under the Creative Commons Attribution 3.0 License, and code is licensed under a BSD license.
Terms of Service | Privacy Policy
