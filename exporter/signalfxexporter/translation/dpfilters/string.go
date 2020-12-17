// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dpfilters

import (
	"regexp"

	"github.com/gobwas/glob"
)

// stringFilter matches against simple strings
type stringFilter interface {
	Matches(string) bool
}

// basicStringFilter will match if any one of the given strings is a match.
type basicStringFilter struct {
	staticSet        map[string]bool
	regexps          []regexMatcher
	globs            []globMatcher
	anyStaticNegated bool
}

var _ stringFilter = (*basicStringFilter)(nil)

// Matches returns true if any one of the strings in the filter matches the
// input s
func (f *basicStringFilter) Matches(s string) bool {
	negated, matched := f.staticSet[s]
	if matched {
		return !negated
	}
	if f.anyStaticNegated {
		return true
	}

	for _, reMatch := range f.regexps {
		if reMatch.re.MatchString(s) != reMatch.negated {
			return true
		}
	}

	for _, globMatch := range f.globs {
		if globMatch.glob.Match(s) != globMatch.negated {
			return true
		}
	}

	return false
}

// newBasicStringFilter returns a filter that can match against the provided items.
func newBasicStringFilter(items []string) (*basicStringFilter, error) {
	staticSet := make(map[string]bool)
	var regexps []regexMatcher
	var globs []globMatcher

	anyStaticNegated := false
	for _, i := range items {
		m, negated := stripNegation(i)
		switch {
		case isRegex(m):
			var re *regexp.Regexp
			var err error

			reText := stripSlashes(m)
			re, err = regexp.Compile(reText)

			if err != nil {
				return nil, err
			}

			regexps = append(regexps, regexMatcher{re: re, negated: negated})
		case isGlobbed(m):
			g, err := glob.Compile(m)
			if err != nil {
				return nil, err
			}

			globs = append(globs, globMatcher{glob: g, negated: negated})
		default:
			staticSet[m] = negated
			if negated {
				anyStaticNegated = true
			}
		}
	}

	return &basicStringFilter{
		staticSet:        staticSet,
		regexps:          regexps,
		globs:            globs,
		anyStaticNegated: anyStaticNegated,
	}, nil
}

// overridableStringFilter matches input strings that are positively matched by
// one of the input filters AND are not excluded by any negated filters (they
// work kind of like how .gitignore patterns work), OR are exactly matched by a
// literal filter input (e.g. not a globbed or regex pattern).  Order of the
// items does not matter.
type overridableStringFilter struct {
	*basicStringFilter
}

// newOverridableStringFilter makes a new overridableStringFilter with the given
// items.
func newOverridableStringFilter(items []string) (*overridableStringFilter, error) {
	basic, err := newBasicStringFilter(items)
	if err != nil {
		return nil, err
	}

	return &overridableStringFilter{
		basicStringFilter: basic,
	}, nil
}

// Matches if s is positively matched by the filter items OR
// if it is positively matched by a non-glob/regex pattern exactly
// and is negated as well.  See the unit tests for examples.
func (f *overridableStringFilter) Matches(s string) bool {
	negated, matched := f.staticSet[s]
	// If a metric is negated and it matched it won't match anything else by
	// definition.
	if matched && negated {
		return false
	}

	for _, reMatch := range f.regexps {
		reMatched, negated := reMatch.Matches(s)
		if reMatched && negated {
			return false
		}
		matched = matched || reMatched
	}

	for _, globMatcher := range f.globs {
		globMatched, negated := globMatcher.Matches(s)
		if globMatched && negated {
			return false
		}
		matched = matched || globMatched
	}
	return matched
}
