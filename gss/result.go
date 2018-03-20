package gss

// Result Parser response
type Result struct {
	Feed         *Feed
	RSSType      RSSType
	isSuccessful bool
}

// IsSuccess Result is success or error
func (r *Result) IsSuccessful() bool {
	if r.isSuccess {
		return true
	}
	return false
}

// IsRSS1 feed is RSS1?
func (r *Result) IsRSS1() bool {
	if r.RSSType == RSS1 {
		return true
	}
	return false
}

// IsRSS2 feed is RSS2?
func (r *Result) IsRSS2() bool {
	if r.RSSType == RSS2 {
		return true
	}
	return false
}

// IsAtom feed is atom?
func (r *Result) IsAtom() bool {
	if r.RSSType == Atom {
		return true
	}
	return false
}
