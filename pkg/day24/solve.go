package day24

import (
	"adventofcode2024/pkg/utils"
	"adventofcode2024/pkg/utils/assert"
	"adventofcode2024/pkg/utils/logger"
	"adventofcode2024/pkg/utils/maps"

	// slice "adventofcode2024/pkg/utils/slices"
	"slices"
	"strconv"
	"strings"
)

type Operator int

const (
	And Operator = iota
	Or
	Xor
)

func NewOperator(value string) Operator {
	switch value {
	case "XOR":
		return Xor
	case "OR":
		return Or
	case "AND":
		return And
	default:
		panic(assert.Fail("invalid operator", value))
	}
}

func (o Operator) String() string {
	switch o {
	case And:
		return "AND"
	case Or:
		return "OR"
	case Xor:
		return "XOR"
	default:
		panic(assert.Fail("invalid operator", o))
	}
}

type Rule struct {
	left     string
	operator Operator
	right    string
	original string
}

func NewRule(left string, operator Operator, right string, output string) Rule {
	return Rule{left, operator, right, output}
}

// func lookup(left string, operator Operator, right string) string {
// 	if strings.Compare(left, right) > 0 {
// 		left, right = right, left
// 	}
// 	return left + " " + operator.String() + " " + right
// }

var log = logger.New("day24")

func PartOne(path string) int {
	log.Reset("starting")
	state, rules := parse(path)
	done := false
	for !done {
		done = true
		for _, rule := range rules {
			done = rule.apply(state, make(map[string]string)) && done
		}
	}
	return readValue(state, "z")
}

type Adder struct {
	bit      int    // the z bit index we're outputting to
	carryOut string // (z{bit} AND z{bit}) will give us the carry bit forward
	carryIn  string // the carry bit from bit-1
}

func PartTwo(path string) string {
	/*
		7.  [day24] z00 path=x00 XOR y00, Δ=888.5 µs
		8.  [day24] z01 path=(x00 AND y00) XOR (x01 XOR y01), Δ=910.9 µs
		9.  [day24] z02 path=(((x00 AND y00) AND (x01 XOR y01)) OR (x01 AND y01)) XOR (x02 XOR y02), Δ=1.0 ms
		10. [day24] z03 path=((x02 AND y02) OR ((((x00 AND y00) AND (x01 XOR y01)) OR (x01 AND y01)) AND (x02 XOR y02))) XOR (x03 XOR y03), Δ=1.0 ms
		11. [day24] z04 path=((x03 AND y03) OR (((x02 AND y02) OR ((((x00 AND y00) AND (x01 XOR y01)) OR (x01 AND y01)) AND (x02 XOR y02))) AND (x03 XOR y03))) XOR (x04 XOR y04), Δ=1.0 ms

	*/
	log.Reset("part two")

	// _, rulesSlice := parse(path)
	// rules := make(map[string]*Rule)
	// for _, rule := range rulesSlice {
	// 	rules[lookup(rule.left, rule.operator, rule.right)] = &rule
	// }
	//
	// // known := make(map[string]*Rule)
	//
	// for i := range 44 {
	// 	x := fmt.Sprintf("x%02d", i)
	// 	y := fmt.Sprintf("y%02d", i)
	// 	z := fmt.Sprintf("z%02d", i)
	// 	v, exist := rules[lookup(x, Xor, y)]
	// 	if !exist {
	// 		assert.Fail("expected value not found", "rule", lookup(x, Xor, y))
	// 	}
	// 	v.expected = z
	// }
	//
	// rules[lookup("x00", Xor, "y00")].expected = "z00"
	// rules[lookup("x00", And, "y00")].expected = "z00"
	//
	// for _, rule := range rules {
	// 	if rule.expected != rule.original && rule.expected != "" {
	// 		// log.Log("found difference", logger.With("rule", rule))
	// 	}
	// }

	log.Log("done new part two")
	return ""
}

func debugPath(rules []Rule, output string) string {
	for _, rule := range rules {
		if rule.original == output {
			left := rule.left
			right := rule.right
			if !strings.HasPrefix(left, "x") && !strings.HasPrefix(left, "y") {
				left = "(" + debugPath(rules, left) + ")"
			}
			if !strings.HasPrefix(right, "y") && !strings.HasPrefix(right, "x") {
				right = "(" + debugPath(rules, right) + ")"
			}
			return left + " " + rule.operator.String() + " " + right
		}
	}

	panic(assert.Fail("should not reach here", "output", output))
}

func readValue(state map[string]int, variable string) int {
	out := ""
	keys := make([]string, 0)
	for key := range state {
		if strings.HasPrefix(key, variable) {
			keys = append(keys, key)
		}
	}
	slices.Sort(keys)
	slices.Reverse(keys)
	for _, key := range keys {
		out += strconv.Itoa(state[key])
	}

	return int(utils.Must(strconv.ParseInt(out, 2, 64)))
}

func parse(path string) (map[string]int, []Rule) {
	starting := make(map[string]int)
	rules := make([]Rule, 0)
	initializing := true
	for _, line := range utils.MustReadInput(path) {
		if line == "" {
			initializing = false
			continue
		}
		if initializing {
			parts := strings.Split(line, ": ")
			starting[parts[0]] = utils.MustAtoi(parts[1])
			// log.Log("initial value", logger.With("wire", parts[0]), logger.With("value", parts[1]))
		} else {
			parts := strings.Split(line, " ")
			rules = append(rules, NewRule(parts[0], NewOperator(parts[1]), parts[2], parts[4]))
			// log.Log("adding rule", logger.With("rule", rules[len(rules)-1]))
		}
	}
	return starting, rules
}

func Original(path string) string {
	log.Reset("original jank solution")
	state, rules := parse(path)

	x := readValue(state, "x")
	y := readValue(state, "y")
	log.Log("x value", logger.With("value", x))
	log.Log("y value", logger.With("value", y))

	expected := strconv.FormatInt(int64(x+y), 2)
	log.Log("z expected", logger.With("value", expected), logger.With("decimal", x+y))

	/*
		5. [day24] z expected value=1011110001110011000101110011000001010101001000, Δ=332.0 µs
		6. [day24] z actual   value=1011110010010011000101110011000001010101001000, Δ=1.0 ms
	*/
	swaps := make(map[string]string)
	swaps["vss"] = "z14" // 14
	swaps["z14"] = "vss"

	swaps["kdh"] = "hjf" // 22
	swaps["hjf"] = "kdh"

	swaps["kpp"] = "z31" // 31
	swaps["z31"] = "kpp"

	// swaps["bbc"] = "z35" // 35 - wrong
	// swaps["z35"] = "bbc"

	swaps["sgj"] = "z35" // 35 - right
	swaps["z35"] = "sgj"

	done := false
	for !done {
		done = true
		for _, rule := range rules {
			done = rule.apply(state, swaps) && done
		}
	}
	actual := readValue(state, "z")
	log.Log("z actual  ", logger.With("value", strconv.FormatInt(int64(actual), 2)), logger.With("decimal", actual))
	log.Log("equals?", logger.With("equals", actual == (x+y)))

	// assert.Equal(x+y, actual, "expected z value to be x+y")

	// combinations := slice.Combinations(rules, 8)
	// log.Log("combination count", logger.With("count", len(combinations)))

	// for idx := range 44 {
	// 	z := fmt.Sprintf("z%02d", idx)
	// 	path := debugPath(rules, z)
	// 	stats := make(map[string]int)
	//
	// 	for range 44 {
	// 		// looking := fmt.Sprintf("x%02d", oi)
	// 		// count := len(utils.Indexes(path, looking))
	// 		// if count > 0 {
	// 		// 	stats[looking] += count
	// 		// }
	//
	// 		// looking = fmt.Sprintf("y%02d", oi)
	// 		// count = len(utils.Indexes(path, looking))
	// 		// if count > 0 {
	// 		// 	stats[looking] += count
	// 		// }
	//
	// 		stats["ands"] = len(utils.Indexes(path, " AND "))
	// 		stats["ors"] = len(utils.Indexes(path, " OR "))
	// 		stats["xors"] = len(utils.Indexes(path, " XOR "))
	// 	}
	//
	// 	if idx >= 1 && (stats["xors"] != idx+1 || stats["ors"] != idx-1 || stats["ands"] != (idx*2)-1) {
	// 		log.Log(z)
	// 		// log.Log(z, logger.With("path", path))
	// 		// log.Log(z, logger.With("stats", stats), logger.IndentOnce)
	// 	}
	//
	// }
	// log.Log("z02", logger.With("path", debugPath(rules, "z02")))
	// log.Log("z29", logger.With("path", debugPath(rules, "z29")))
	// log.Log("--")
	// log.Log("z35", logger.With("path", debugPath(rules, "z35")))

	log.Log("jank done", logger.IncludeTotal)

	keys := maps.Keys(swaps)
	slices.Sort(keys)
	return strings.Join(keys, ",")
}

func (r Rule) apply(state map[string]int, swaps map[string]string) bool {
	output := r.original
	if swap, exists := swaps[output]; exists {
		output = swap
	}

	if _, exists := state[output]; exists {
		return true
	}

	var left int
	var right int
	var exists bool

	if left, exists = state[r.left]; !exists {
		return false
	}
	if right, exists = state[r.right]; !exists {
		return true
	}

	switch r.operator {
	case And:
		state[output] = left & right
	case Or:
		state[output] = left | right
	case Xor:
		state[output] = left ^ right
	}

	return true
}
