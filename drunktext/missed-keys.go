package drunktext

var MissedKeys = map[rune][]rune{
	'q': {'w', 'a', 's'},
	'w': {'q', 's', 'e'},
	'e': {'w', 'd', 'r'},
	'r': {'e', 'f', 't'},
	't': {'r', 'g', 'y'},
	'y': {'t', 'h', 'u'},
	'u': {'y', 'j', 'i'},
	'i': {'u', 'k', 'o'},
	'o': {'i', 'l', 'p'},
	'p': {'o', 'l', ';'},
	'a': {'q', 's', 'z'},
	's': {'a', 'd', 'w'},
	'd': {'s', 'e', 'f'},
	'f': {'d', 'r', 'g'},
	'g': {'f', 't', 'h'},
	'h': {'g', 'y', 'j'},
	'j': {'h', 'u', 'k'},
	'k': {'j', 'i', 'l'},
	'l': {'k', 'o', ';'},
	'z': {'a', 's', 'x'},
	'x': {'z', 's', 'x'},
	'c': {'x', 'd', 'v'},
	'v': {'c', 'f', 'b'},
	'b': {'v', 'g', 'n'},
	'n': {'b', 'h', 'm'},
	'm': {'n', 'j', ','},
}