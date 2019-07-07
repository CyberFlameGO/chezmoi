// +build windows

package cmd

// exec, on windows, calls run since legit exec doesn't really exist.
func (c *Config) exec(argv []string) error {
	return c.run("", argv[0], argv[1:]...)
}
