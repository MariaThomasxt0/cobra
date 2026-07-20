func (c *Command) MarkFlagRequired(name string) error {
	flag := c.Flags().Lookup(name)
	if flag == nil {
		flag = c.PersistentFlags().Lookup(name)
	}
	if flag == nil {
		return fmt.Errorf("flag %q does not exist", name)
	}
	flag.Annotations = append(flag.Annotations, map[string][]string{BashCompOneRequiredFlag: {}}...)
	return nil
}