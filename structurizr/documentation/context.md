## Contexto

La empresa PITITA desea implementar un software de control de ASISTENCIA (SEGURITO)

![](embed:Contexto)

### Software de control de asistencia 

- Existe una terminal donde se instala el modulo EVO (accesso al sistema)
- En este modulo de acceso (WEB), un usuario puede registrar su acceso/salida a/de la oficina (entradas/salidas) introduciendo un codigo de 7 caracteres (y tambien indicar que se trata de una salida o entrada)
- Para la verificaccion de acceso, EVO, debe usar un sistema ERP legacy (MESA*), que lastimosamente no se puede remplazar. Este sistema es una aplicacion EXE que permite indicar si el codigo ingresado (via STDIN) es valido imprimiendo en consola (STDOUT), e.j.
* El app MESA puede ser bien un ejecutable windows, mac o linux.

Con codigo valido

```
c:> mesa.exe 123456
ID: JROCA101
User: Javier ROCA

Con un codigo invalido
c:> mesa.exe 123451
Unknown
```

- El sistema MESA utiliza un archivo XML para esta verificacion (usuario-codigo), este archivo por lo general no cambia.  
- Luego de que un usuario es validado, SEGURITO debe registrar este evento de entrada/salida con la fecha del mismo en sus modulo LOG.
- El  modulo LOG debe registrar esta informacion en otro modulo legacy (PUMARI), que es un ERP (applicacion EXE) que cuenta con una linea de comando donde se realiza el registro, e.j.:

```
c:>pumari.exe -u JROCA101 -e ENTRADA -d 10/10/2019 12:23:00

```

- El modulo PUMARI almacena este log en un archivo de texto. Por razones tecnicas, el grabado de un evento no puede ser concurrente, solo se peude grabar un archivo a la vez, y el grabado implica un tiempo de 15s por registro.

- El modulo CAMACHO es un servicio (RESTful) que permite consultar los eventos de un determinado usuario (e.g. lista de entradas y salidas en el tiempo). Este sistema puede utilizar cualquier tecnologia y persitencia.

                        
    

    
    