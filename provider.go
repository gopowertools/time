package ttime

import (
	"sync"
	"time"
)

// Provider provides the time for systems. Use instead of the standard library
// time functions to make it possible to run asynchronous tests which need to
// modify time.
type Provider struct {
	sync.RWMutex

	// now is the time to use. If nil, then the current system time is used.
	now *time.Time
}

// NewProvider returns a Provider.
func NewProvider() *Provider {
	return &Provider{}
}

// Now returns a time previously set with FixNow(), or if FixNow() wasn't called
// earlier, the current system time.
func (p *Provider) Now() time.Time {
	p.RLock()
	defer p.RUnlock()

	if p.now == nil {
		return time.Now()
	}
	return *p.now
}

// FixNow causes the given time t to be stored. All future calls to Now(),
// Until() and Since() will then use t instead of the current system time.
func (p *Provider) FixNow(t time.Time) {
	p.Lock()
	defer p.Unlock()

	p.now = &t
}

// Since returns the duration between past time pt and now.
// If pt is in the future, the returned duration is negative.
func (p *Provider) Since(pt time.Time) time.Duration {
	return p.Now().Sub(pt)
}

// Until returns the duration between now and future time ft.
// If ft is in the past, the returned duration is negative.
func (p *Provider) Until(ft time.Time) time.Duration {
	return -p.Since(ft)
}
