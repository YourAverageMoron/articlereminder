package reminders

import (
	"os/exec"
)

type List struct {
	name string
}

func NewList(name string)*List{
   return &List{
        name: name,
   } 
}

func (r *List) Add(article string) error {
	_, err := exec.Command("reminders", "add", r.name, article).Output()
	if err != nil {
		return err
	}
	return nil
}
