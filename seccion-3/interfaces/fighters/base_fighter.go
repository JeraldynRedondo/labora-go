package fighters

// BaseFighter it is a structure that represents the fighter its default variables.
type BaseFighter struct {
	Life        int // 0..100
	InitialLife int
}

// IsAlive it is a function that validates if the fighter is alive.
func (bf BaseFighter) IsAlive() bool {
	return bf.Life > 0
}

// GetLife it is a function that returns the life of the fighter.
func (bf BaseFighter) GetLife() int {
	return bf.Life
}
