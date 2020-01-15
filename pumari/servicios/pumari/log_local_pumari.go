package mesa

import (
	"context"
	log2 "log"
	"os/exec"
	"pumari/log"
	"pumari/modelos"
	"time"
)

type servicioLogPumari struct {
	contextTimeout    time.Duration
	path              string
	nombreApplicacion string
}

func NuevoServicioDeLogPumari(timeout time.Duration, path string) log.ServicioLog {
	return &servicioLogPumari{
		contextTimeout:    timeout,
		path:              path,
		nombreApplicacion: "pumari",
	}
}

func (s servicioLogPumari) RegistrarEvento(ctx context.Context, mensaje modelos.MensanjeEvento) error {
	strPath := s.path + s.nombreApplicacion
	typeString := ""
	switch mensaje.Tipo {
	case modelos.EntradaTipoEvento:
		typeString = "ENTRADA"
	case modelos.SalidaTipoEvento:
		typeString = "SALIDA"
	}
	//-u JROCA101 -e ENTRADA -d 10/10/2019 12:23:00
	log2.Print("Comando enviado a pumari exec")
	cmd := exec.Command(strPath, "-u", mensaje.UserName, "-e", typeString, mensaje.Fecha.Format("dd/MM/yyyy"), mensaje.Fecha.Format("hh:mm:ss"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log2.Print(err)
		return err
	}
	response := string(out)
	log2.Print(response)
	return nil
}
