package provider

import (
	"fmt"
	"os/exec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVMClone() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMCloneCreate,
		Read:   resourceVMCloneRead,
		Delete: resourceVMCloneDelete,

		Schema: map[string]*schema.Schema{
			"source_vm": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVMCloneCreate(d *schema.ResourceData, meta interface{}) error {
	sourceVM := d.Get("source_vm").(string)
	name := d.Get("name").(string)

	cmd := exec.Command("VBoxManage", "clonevm", sourceVM,
		"--name", name,
		"--register",
		"--mode", "machine",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("помилка клонування VM: %s\n%s", err, out)
	}

	d.SetId(name)
	return nil
}

func resourceVMCloneRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	cmd := exec.Command("VBoxManage", "showvminfo", name)
	err := cmd.Run()
	if err != nil {
		d.SetId("")
	}

	return nil
}

func resourceVMCloneDelete(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	// Зупинити VM
	exec.Command("VBoxManage", "controlvm", name, "poweroff").Run()

	// Видалити VM
	cmd := exec.Command("VBoxManage", "unregistervm", name, "--delete")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("помилка видалення VM: %s\n%s", err, out)
	}

	d.SetId("")
	return nil
}
