package reload

import (
	"net/http"
	"os/exec"
	"log"
	"io"
)

func run(w http.ResponseWriter, cmd *exec.Cmd, errmsg string) bool {
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("ERROR:", errmsg)
		log.Println(string(output))
		http.Error(w, `{"error": "`+errmsg+`"}`, http.StatusInternalServerError)
		return false
	}
	return true
}

func Fitsbensmall(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("sh", "-c", "supervisorctl stop apache")
	cmd.Dir = "/"
	if !run(w, cmd, "unable to stop apache service") { return }

	cmd = exec.Command("sh", "-c", "pkill -f apache2")
	cmd.Dir = "/"
	run(w, cmd, "unable to kill apache service")

	cmd = exec.Command("sh", "-c", "rm -rf /usr/itsbensmall.com")
	cmd.Dir = "/"
	if !run(w, cmd, "unable to remove files") { return }

	cmd = exec.Command("sh", "-c", "git clone https://github.com/bsmallyt/itsbensmall.com.git")
	cmd.Dir = "/usr"
	if !run(w, cmd, "unable to clone site") { return }

	cmd = exec.Command("sh", "-c", "npx ng build itsbensmall")
	cmd.Dir = "/usr/itsbensmall.com"
	if !run(w, cmd, "unable to build site") { return }

	cmd = exec.Command("sh", "-c", "cp -r /usr/itsbensmall.com/dist/itsbensmall /var/www/html")
	cmd.Dir = "/"
	if !run(w, cmd, "unable to copy build") { return }

	cmd = exec.Command("sh", "-c", "supervisorctl start apache")
	cmd.Dir = "/"
	if !run(w, cmd, "unable to restart apache service") { return }

	io.WriteString(w, `{"status": "success"}`)
}
