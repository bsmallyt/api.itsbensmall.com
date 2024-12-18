package reload

import (
	"net/http"
	"os/exec"
)

func Fitsbensmall(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sh", "-c", "supervisorctl stop apache")
	cmd.Path = "/"
	err := cmd.Run()
	if err != nil {
		http.Error(w, `{"error": "unable to stop apache service"}`, http.StatusInternalServerError)
	}
	cmd = exec.Command("sh", "-c", "pkill apache2")
	err = cmd.Run()
	if err != nil {
		http.Error(w, `{"error": "unable to kill apache service"}`, http.StatusInternalServerError)
	}
	cmd = exec.Command("sh", "-c", "rm -r usr/itsbensmall.com")
	err = cmd.Run()
	if err != nil {
		http.Error(w, `{"error": "unable to remove files"}`, http.StatusInternalServerError)
	}
	
	cmd = exec.Command("sh", "-c", "git clone https://github.com/bsmallyt/itsbensmall.com.git")
	cmd.Path = "/usr"
	err = cmd.Run()
	if err != nil {
		http.Error(w, `{"error": "unable to clone site"}`, http.StatusInternalServerError)
	}
	cmd = exec.Command("sh", "-c", "ng build itsbensmall")
	cmd.Path = "/usr/itsbensmall.com"
	err = cmd.Run()
	if err != nil {
		http.Error(w, `{"error": "unable to build site"}`, http.StatusInternalServerError)
	}

	cmd = exec.Command("sh", "-c", "cp -r /usr/itsbensmall.com/dist/itsbensmall /var/www/html")
	cmd.Path = "/"
	err = cmd.Run()
	if err != nil {
		http.Error(w, `{"error": "unable to copy build"}`, http.StatusInternalServerError)
	}
	cmd = exec.Command("sh", "-c", "supervisorctl start apache")
	err = cmd.Run()
	if err != nil {
		http.Error(w, `{"error": "unable to restart apache service"}`, http.StatusInternalServerError)
	}
}